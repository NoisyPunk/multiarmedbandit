package internalgrpc

import (
	"context"
	"github.com/NoisyPunk/multiarmedbandit/internal/server/grpc/pb"
)

func (g *GRPCServer) AddBanner(ctx context.Context,
	req *pb.AddBannerRequest,
) (*pb.AddBannerResponse, error) {
	bannerId, err := g.application.AddBanner(ctx, req.Description)
	if err != nil {
		return nil, err
	}
	response := &pb.AddBannerResponse{
		BannerId: bannerId.String(),
		Message:  "Banner created successfully!",
	}

	return response, nil
}

func (g *GRPCServer) AddGroup(ctx context.Context, req *pb.AddGroupRequest) (*pb.AddGroupResponse, error) {
	groupId, err := g.application.AddGroup(ctx, req.Description)
	if err != nil {
		return nil, err
	}
	response := &pb.AddGroupResponse{
		GroupId: groupId.String(),
		Message: "Group created successfully",
	}
	return response, nil
}

func (g *GRPCServer) AddSlot(ctx context.Context, req *pb.AddSlotRequest) (*pb.AddSlotResponse, error) {
	slotId, err := g.application.AddSlot(ctx, req.Description)
	if err != nil {
		return nil, err
	}
	response := &pb.AddSlotResponse{
		SlotId:  slotId.String(),
		Message: "Slot created successfully",
	}
	return response, nil
}

func (g *GRPCServer) AddRotation(ctx context.Context, req *pb.AddRotationRequest) (*pb.AddRotationResponse, error) {
	rotationId, err := g.application.AddRotation(ctx, req.BannerId, req.SlotId, req.GroupId)
	if err != nil {
		return nil, err
	}
	response := &pb.AddRotationResponse{
		RotationId: rotationId.String(),
		Message:    "Rotation created successfully",
	}
	return response, nil
}

func (g *GRPCServer) RegisterClick(ctx context.Context,
	req *pb.RegisterClickRequest) (*pb.RegisterClickResponse, error) {
	err := g.application.RegisterClick(ctx, req.RotationId)
	if err != nil {
		return nil, err
	}
	response := &pb.RegisterClickResponse{
		Message: "Click registered",
	}
	return response, nil
}

func (g *GRPCServer) ShowBanner(ctx context.Context, req *pb.ShowBannerRequest) (*pb.ShowBannerResponse, error) {
	rotation, err := g.application.ChooseRotationForSlot(ctx, req.SlotId, req.GroupId)
	if err != nil {
		return nil, err
	}
	response := &pb.ShowBannerResponse{
		BannerId:   rotation.BannerID.String(),
		RotationId: rotation.ID.String(),
	}
	return response, nil
}
