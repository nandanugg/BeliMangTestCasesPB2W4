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

	var record []*pb.MerchantNearestRecord
	for _, r := range res {
		merchants := make(map[int64]*pb.Merchant)
		for k, v := range r.MerchantOrder {
			merchants[int64(k)] = &pb.Merchant{
				MerchantId:     v.MerchantId,
				PregeneratedId: v.PregeneratedId,
				Location: &pb.LocationPoint{
					Lat:  float32(v.Location.Lat),
					Long: float32(v.Location.Long),
				},
			}
		}
		record = append(record, &pb.MerchantNearestRecord{
			StartingPoint: &pb.LocationPoint{
				Lat:  float32(r.StartingPoint.Lat),
				Long: float32(r.StartingPoint.Long),
			},
			Merchants: merchants,
		})
	}
	return &pb.AllMerchantNearestRecord{
		Records: record,
	}, nil
}

func (h *Handler) GetAllMerchantRoutes(ctx context.Context, _ *emptypb.Empty) (*pb.AllGeneratedRoutes, error) {
	srvRes, err := h.srv.GetAllMerchantRoutes()
	if err != nil {
		return nil, err
	}

	var allGeneratedRoutes []*pb.GeneratedRoutes
	for _, zone := range srvRes {
		generatedRoute := make(map[int64]*pb.Merchant)
		for _, a := range zone.GeneratedRoutes {
			for order, r := range a.GeneratedRoutes {
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
					Lat:  float32(a.StartingPoint.Lat),
					Long: float32(a.StartingPoint.Long),
				},
				TotalDuration: int64(a.TotalDurationInMinute),
				TotalDistance: float32(a.TotalDistance),
				StartingIndex: int64(a.StartingIndex),
			})
		}

	}

	return &pb.AllGeneratedRoutes{
		Routes: allGeneratedRoutes,
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
