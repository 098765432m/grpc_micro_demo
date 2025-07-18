package service

type Hotel struct {
	Id   string
	Name string
}

type HotelService interface {
}

type HotelServiceImpl struct{}

func GetHotelById(id string) (*Hotel, error) {
	return &Hotel{
		Id:   "121",
		Name: "asdasd",
	}, nil
}
