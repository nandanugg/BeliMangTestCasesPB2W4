package handler

import (
	"context"

	"github.com/nandanugg/BeliMangTestCasesPB2W4/caddy/entity/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Handler) GetAllMerchantNearestLocations(ctx context.Context, _ *emptypb.Empty) (*pb.AllMerchantNearestRecord, error) {
	res, err := h.srv.GetAllMerchantNearestLocations()
	if err != nil {
		return nil, err
	}

	var zones []*pb.MerchantNearestRecordZone
	for _, zone := range res {
		var recordsRes []*pb.MerchantNearestRecord
		merchantsOrder := make(map[int64]*pb.Merchant)
		for _, record := range zone.Records {
			for order, merchant := range record.MerchantOrder {
				merchantsOrder[int64(order)] = &pb.Merchant{
					MerchantId:     merchant.MerchantId,
					PregeneratedId: merchant.PregeneratedId,
					Location: &pb.LocationPoint{
						Lat:  float32(merchant.Location.Lat),
						Long: float32(merchant.Location.Long),
					},
				}
			}

			recordsRes = append(recordsRes, &pb.MerchantNearestRecord{
				StartingPoint: &pb.LocationPoint{
					Lat:  float32(record.StartingPoint.Lat),
					Long: float32(record.StartingPoint.Long),
				},
				Merchants: merchantsOrder,
			})
		}
		zones = append(zones, &pb.MerchantNearestRecordZone{
			Records: recordsRes,
		})
	}
	return &pb.AllMerchantNearestRecord{
		Zones: zones,
	}, nil
}

func (h *Handler) GetAllMerchantRoutes(ctx context.Context, _ *emptypb.Empty) (*pb.AllGeneratedRoutes, error) {
	srvRes, err := h.srv.GetAllMerchantRoutes()
	if err != nil {
		return nil, err
	}

	zones := []*pb.RouteZone{}
	for _, zone := range srvRes {
		var allGeneratedRoutes []*pb.GeneratedRoutes
		generatedRoute := make(map[int64]*pb.Merchant)
		for _, zoneGeneratedRoutes := range zone.GeneratedRoutes {
			for order, r := range zoneGeneratedRoutes.GeneratedRoutes {
				generatedRoute[int64(order)] = &pb.Merchant{
					MerchantId:     r.MerchantId,
					PregeneratedId: r.PregeneratedId,
					Location: &pb.LocationPoint{
						Lat:  float32(r.Location.Lat),
						Long: float32(r.Location.Long),
					},
				}
			}
			allGeneratedRoutes = append(allGeneratedRoutes, &pb.GeneratedRoutes{
				GeneratedRoutes: generatedRoute,
				StartingPoint: &pb.LocationPoint{
					Lat:  float32(zoneGeneratedRoutes.StartingPoint.Lat),
					Long: float32(zoneGeneratedRoutes.StartingPoint.Long),
				},
				TotalDuration: int64(zoneGeneratedRoutes.TotalDurationInMinute),
				TotalDistance: float32(zoneGeneratedRoutes.TotalDistance),
				StartingIndex: int64(zoneGeneratedRoutes.StartingIndex),
			})
		}
		zones = append(zones, &pb.RouteZone{
			Routes: allGeneratedRoutes,
		})
	}

	return &pb.AllGeneratedRoutes{
		Zone: zones,
	}, nil
}

func (h *Handler) GetAllPregeneratedMerchants(ctx context.Context, _ *emptypb.Empty) (*pb.PregeneratedMerchant, error) {
	srvRes := h.srv.GetAllPregeneratedMerchants()

	z := []*pb.Merchant{}
	for _, v := range srvRes {
		z = append(z, &pb.Merchant{
			MerchantId:     v.MerchantId,
			PregeneratedId: v.PregeneratedId,
			Location: &pb.LocationPoint{
				Lat:  float32(v.Location.Lat),
				Long: float32(v.Location.Long),
			},
		})
	}

	return &pb.PregeneratedMerchant{
		Merchant: z,
	}, nil
}

func (h *Handler) AssignPregeneratedMerchant(ctx context.Context, req *pb.AssignMerchant) (*emptypb.Empty, error) {
	err := h.srv.AssignPregeneratedMerchant(req.PregeneratedId, req.MerchantId)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
