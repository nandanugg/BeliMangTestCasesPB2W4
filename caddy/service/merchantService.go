package service

import (
	"errors"
	"fmt"
	"log"
	"math/rand/v2"
	"sort"

	"github.com/nandanugg/BeliMangTestCasesPB2W4/caddy/entity"
)

type MerchantService struct {
	MerchantZoneRecord    []*entity.MerchantZoneRecord
	MerchantNearestRecord []*entity.MerchantNearestRecord
	PregeneratedMerchants map[entity.PregeneratedId]*entity.Merchant
	AssignedMerchants     map[entity.PregeneratedId]*entity.Merchant
}

func NewMerchantService() *MerchantService {
	return &MerchantService{}
}

func (service *MerchantService) GetAllMerchantNearestLocations() ([]*entity.MerchantNearestRecord, error) {
	zoneRecords := service.MerchantZoneRecord
	for _, zone := range zoneRecords {
		for preRegeneratedId := range zone.MerchantPregeneratedId {
			_, isExists := service.AssignedMerchants[entity.PregeneratedId(fmt.Sprint(preRegeneratedId))]
			if !isExists {
				return nil, errors.New("some merchant is still not assigned")
			}
		}
	}
	return service.MerchantNearestRecord, nil
}

func (service *MerchantService) GetAllMerchantRoutes() ([]*entity.MerchantZoneRecord, error) {
	zoneRecords := service.MerchantZoneRecord

	for _, zone := range zoneRecords {
		for preRegeneratedId := range zone.MerchantPregeneratedId {
			_, isExists := service.AssignedMerchants[entity.PregeneratedId(fmt.Sprint(preRegeneratedId))]
			if !isExists {
				return nil, errors.New("some merchant is still not assigned")
			}
		}
	}

	return service.MerchantZoneRecord, nil
}

// Get all pregenerated merchants
func (service *MerchantService) GetAllPregeneratedMerchants() map[entity.PregeneratedId]*entity.Merchant {
	return service.PregeneratedMerchants
}

func (service *MerchantService) InitMerchantNearestLocations(generateCount int) {
	for a := 0; a < len(service.MerchantZoneRecord); a++ {
		currentZone := service.MerchantZoneRecord[a]
		userLocation := entity.GenerateRandomLocation(entity.LocationPoint{
			Lat:  currentZone.StartZoneRange.Lat,
			Long: currentZone.StartZoneRange.Long,
		}, entity.LocationPoint{
			Lat:  currentZone.EndZoneRange.Lat,
			Long: currentZone.EndZoneRange.Long,
		})

		for i := 0; i < generateCount; i++ {
			merchantPregeneratedIds := currentZone.MerchantPregeneratedId
			merchants := make([]*entity.Merchant, 0)
			distances := make(map[string]float64)

			for _, m := range merchantPregeneratedIds {
				merchant, isExists := service.PregeneratedMerchants[entity.PregeneratedId(m)]
				if isExists {
					distance := entity.CalculateDistance(userLocation, merchant.Location)
					distances[merchant.PregeneratedId] = distance
					merchants = append(merchants, merchant)
				}
			}

			sort.SliceStable(merchants, func(i, j int) bool {
				return distances[merchants[i].PregeneratedId] < distances[merchants[j].PregeneratedId]
			})

			nearestMerchant := make(map[entity.Order]*entity.Merchant)
			for j, merchant := range merchants {
				nearestMerchant[entity.Order(j)] = merchant
			}
			service.MerchantNearestRecord = append(service.MerchantNearestRecord, &entity.MerchantNearestRecord{
				StartingPoint: userLocation,
				MerchantOrder: nearestMerchant,
			})
		}
	}
	log.Println("Init merchant nearest locations")
}

// Initialize pregenerated merchants
func (service *MerchantService) InitPegeneratedTSPMerchants(generateCount int) {
	for a := 0; a < len(service.MerchantZoneRecord); a++ {
		// for each zone, generate random user location
		merchantZone := service.MerchantZoneRecord[a]
		for i := 0; i < generateCount; i++ {
			// generate random user location based on zone bounds
			userLocation := entity.GenerateRandomLocation(entity.LocationPoint{
				Lat: merchantZone.StartZoneRange.Lat, Long: merchantZone.StartZoneRange.Long,
			}, entity.LocationPoint{
				Lat: merchantZone.EndZoneRange.Lat, Long: merchantZone.EndZoneRange.Long,
			})

			// select random 2-8 merchant pregeneratedId from the zone
			selectedMerchant := make([]*entity.Merchant, 0)
			merchantCount := len(merchantZone.MerchantPregeneratedId)
			itemCount := rand.IntN(7) + 2
			rand.Shuffle(merchantCount, func(i, j int) {
				merchantZone.MerchantPregeneratedId[i], merchantZone.MerchantPregeneratedId[j] = merchantZone.MerchantPregeneratedId[j], merchantZone.MerchantPregeneratedId[i]
			})

			// from pregeneratedId, get the merchant
			for j := 0; j < itemCount; j++ {
				preRegeneratedId := entity.PregeneratedId(merchantZone.MerchantPregeneratedId[j])
				merchant, isExists := service.PregeneratedMerchants[preRegeneratedId]
				if isExists {
					selectedMerchant = append(selectedMerchant, merchant)
				}
			}

			// generate TSP route for the selected merchants
			startingIndex := rand.IntN(len(selectedMerchant))
			routes, totalDistance := entity.GenerateTSPMerchantRouteFromStartingPoint(userLocation, startingIndex, selectedMerchant)
			service.MerchantZoneRecord[a].GeneratedRoutes = append(service.MerchantZoneRecord[a].GeneratedRoutes, &entity.MerchantTSPRoute{
				GeneratedRoutes:       routes,
				StartingPoint:         userLocation,
				TotalDistance:         totalDistance,
				TotalDurationInMinute: entity.CalculateTimeInMinute(totalDistance, 40),
				StartingIndex:         startingIndex,
			})
		}
	}
	log.Println("Init pregenerated TSP merchants")
}

// Initialize zones with pregenerated merchants
func (service *MerchantService) InitZonesWithPregeneratedMerchants(params entity.MerchantZoneOpts) {
	service.PregeneratedMerchants = make(map[entity.PregeneratedId]*entity.Merchant)
	service.AssignedMerchants = make(map[entity.PregeneratedId]*entity.Merchant)
	// init starting point
	var startingPoint entity.LocationPoint
	// create zones contains collection of merchants
	for i := 0; i < params.NumberOfZones; i++ {
		// generate flat square location bounds
		x1, y1, x2, y2 := entity.GenerateFlatSquareLocationBounds(startingPoint, params.Area)

		// create merchants for the zone
		var pregeneratedMerchantIds []string
		for j := 0; j < params.NumberOfMerchantsPerZone; j++ {
			merchants := make(map[string]*entity.Merchant)

			merchant := entity.GenerateRandomMerchant(
				entity.LocationPoint{Lat: x1, Long: y1},
				entity.LocationPoint{Lat: x2, Long: y2},
			)
			merchants[merchant.PregeneratedId] = merchant
			service.PregeneratedMerchants[entity.PregeneratedId(merchant.PregeneratedId)] = merchant

			pregeneratedMerchantIds = append(pregeneratedMerchantIds, merchant.PregeneratedId)
		}

		// append the zone with the merchants into the service
		merchantPregeneratedIds := make([]entity.PregeneratedId, len(pregeneratedMerchantIds))
		for i, id := range pregeneratedMerchantIds {
			merchantPregeneratedIds[i] = entity.PregeneratedId(id)
		}

		service.MerchantZoneRecord = append(service.MerchantZoneRecord, &entity.MerchantZoneRecord{
			StartZoneRange:         entity.LocationPoint{Lat: x1, Long: y1},
			EndZoneRange:           entity.LocationPoint{Lat: x2, Long: y2},
			MerchantPregeneratedId: merchantPregeneratedIds,
		})

		// update the starting point for the next loop
		startingPoint = entity.LocationPoint{Lat: x2, Long: y2 + params.Gap}
	}
	log.Println("Init zones with pregenerated merchants")
}

// Assign pregenerated merchant to assigned merchant
func (service *MerchantService) AssignPregeneratedMerchant(pregeneratedId, merchantId string) error {
	merchant, isExists := service.PregeneratedMerchants[entity.PregeneratedId(pregeneratedId)]
	if !isExists {
		return errors.New("merchant not found")
	}
	merchant.MerchantId = merchantId
	service.AssignedMerchants[entity.PregeneratedId(pregeneratedId)] = merchant
	return nil
}

func (service *MerchantService) ResetAll() {
	service.MerchantZoneRecord = nil
	service.MerchantNearestRecord = nil
	service.PregeneratedMerchants = nil
	service.AssignedMerchants = nil
	log.Println("Reset all")
}
