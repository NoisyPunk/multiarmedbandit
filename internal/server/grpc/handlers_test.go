package internalgrpc

import (
	"context"
	"testing"

	rotator "github.com/NoisyPunk/multiarmedbandit/internal/app"
	"github.com/NoisyPunk/multiarmedbandit/internal/queue"
	"github.com/NoisyPunk/multiarmedbandit/internal/server/grpc/pb"
	"github.com/NoisyPunk/multiarmedbandit/internal/storage"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGRPCServer(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	DBstorage := storage.New()

	dsn := "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=rotator sslmode=disable"

	err := DBstorage.Connect(context.Background(), dsn)
	require.NoError(t, err)

	producer := queue.Producer{}

	require.NoError(t, err)
	app := rotator.App{
		Storage:  *DBstorage,
		Producer: &producer,
	}
	grpcServer := NewGRPCServer(context.Background(), app, "8186")

	go func() {
		if err := grpcServer.Start(); err != nil {
			grpcServer.Stop()
		}
	}()

	conn, err := grpc.Dial("127.0.0.1:8186", grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)

	client := pb.NewRotatorClient(conn)

	reqBanner := &pb.AddBannerRequest{Description: "Test banner"}
	respBanner, err := client.AddBanner(context.Background(), reqBanner)
	require.NoError(t, err)

	require.NotEqual(t, respBanner.BannerId, "")
	require.Equal(t, respBanner.Message, "Banner created successfully!")

	reqGroup := &pb.AddGroupRequest{Description: "Test group"}
	respGroup, err := client.AddGroup(context.Background(), reqGroup)
	require.NoError(t, err)

	require.NotEqual(t, respGroup.GroupId, "")
	require.Equal(t, respGroup.Message, "Group created successfully")

	reqSlot := &pb.AddSlotRequest{Description: "Test slot"}
	respSlot, err := client.AddSlot(context.Background(), reqSlot)
	require.NoError(t, err)

	require.NotEqual(t, respSlot.SlotId, "")
	require.Equal(t, respSlot.Message, "Slot created successfully")

	reqRotation := &pb.AddRotationRequest{
		BannerId: respBanner.BannerId,
		GroupId:  respGroup.GroupId,
		SlotId:   respSlot.SlotId,
	}
	respRotation, err := client.AddRotation(context.Background(), reqRotation)
	require.NoError(t, err)

	require.NotEqual(t, respRotation.RotationId, "")
	require.Equal(t, respRotation.Message, "Rotation created successfully")
}
