package service

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Sup3r-Us3r/go-grpc/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SmartwatchGrpcService struct {
	pb.UnimplementedSmartwatchServiceServer
}

func NewSmartwatchGrpcService() *SmartwatchGrpcService {
	return &SmartwatchGrpcService{}
}

func (s *SmartwatchGrpcService) BeatsPerMinute(
	in *pb.BeatsPerMinuteRequest,
	stream pb.SmartwatchService_BeatsPerMinuteServer,
) error {
	fmt.Println("Smartwatch UUID: ", in.GetUuid())

	var beatCount uint32 = 0

	for {
		select {
		case <-stream.Context().Done():
			return status.Errorf(codes.Canceled, "stream has ended")
		default:
			time.Sleep(2 * time.Second)

			value := 30 + rand.Int31n(80)

			err := stream.SendMsg(&pb.BeatsPerMinuteResponse{
				Value:  uint32(value),
				Minute: uint32(time.Now().Second()),
			})

			beatCount++

			if err != nil {
				return status.Error(codes.Canceled, "stream has ended")
			}

			if beatCount == 10 {
				beatCount = 0

				return nil
			}
		}
	}
}
