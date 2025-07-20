package server

import "github.com/098765432m/grpc-micro-demo/hotel-service/scripts/pb"

type HotelServer struct {
	pb.UnimplementedHotelServiceServer
}

func (hs *HotelServer) NewHotelServer() *HotelServer {
	return &HotelServer{}
}
