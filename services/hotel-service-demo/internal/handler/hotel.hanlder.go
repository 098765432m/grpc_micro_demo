package handler

import (
	"context"

	"github.com/098765432m/hotel-service/internal/service"
	"github.com/098765432m/hotel-service/pb"
)

type HotelHanlderImpl struct {
	pb.UnimplementedHotelServiceServer
	hotelService *service.HotelServiceImpl
}

func NewHotelHandler(hotelService *service.HotelServiceImpl) *HotelHanlderImpl {
	return &HotelHanlderImpl{
		hotelService: hotelService,
	}
}

func (hh *HotelHanlderImpl) GetHotelById(ctx context.Context, req *pb.GetHotelByIdRequest) (*pb.GetHotelByIdResponse, error) {
	return &pb.GetHotelByIdResponse{
		Id:   "123",
		Name: "Say dam",
	}, nil
}
