package handler

import (
	"context"

	"github.com/nandanugg/BeliMangTestCasesPB2W4/caddy/entity"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Handler) InitMerchantNearestLocations(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	h.srv.InitMerchantNearestLocations(2)
	return &emptypb.Empty{}, nil
}

func (h *Handler) InitPegeneratedTSPMerchants(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	// TODO: change this to be based on the request input
	h.srv.InitPegeneratedTSPMerchants(2)
	return &emptypb.Empty{}, nil
}

func (h *Handler) InitZonesWithPregeneratedMerchants(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	// TODO: change this to be based on the request input
	h.srv.InitZonesWithPregeneratedMerchants(entity.MerchantZoneOpts{
		Area:                     1,
		Gap:                      10,
		NumberOfZones:            2,
		NumberOfMerchantsPerZone: 10,
	})
	return &emptypb.Empty{}, nil
}

func (h *Handler) ResetAll(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	h.srv.ResetAll()
	return &emptypb.Empty{}, nil
}
