package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../entity"
)

var (
	ownerService     OwnerService = NewOwnerService()
	carService       CarService   = NewCarService()
	carDatachannel                = make(chan *http.Response)
	ownerDatachannel              = make(chan *http.Response)
)

type CarDetailsService interface {
	GetDetails() entity.CarDetails
}

type service struct{}

func NewCarDetailsService() CarDetailsService {
	return &service{}
}

func (*service) GetDetails() entity.CarDetails {
	// goroutine call endpoint 1
	go carService.FetchData()

	// goroutine call endpoint 2
	go ownerService.FetchData()

	// create carChannel to get the data from endpoint 1:
	car, _ := getCarData()
	// ownerChannel: get endpoint 2 data
	owner, _ := getOwnerData()

	return entity.CarDetails{
		ID:             car.CarData.ID,
		Brand:          car.CarData.Brand,
		Model:          car.CarData.Model,
		Year:           car.CarData.Year,
		Vin:            car.CarData.Vin,
		OwnerFirstName: owner.OwnerData.FirstName,
		OwnerLastName:  owner.OwnerData.LastName,
		OwnerEmail:     owner.OwnerData.Email,
		OwnerJobTitle:  owner.OwnerData.JobTitle,
	}
}

func getCarData() (entity.Car, error) {
	r1 := <-carDatachannel
	var car entity.Car
	err := json.NewDecoder(r1.Body).Decode(&car)
	if err != nil {
		fmt.Print(err.Error())
		return car, err
	}
	return car, nil
}

func getOwnerData() (entity.Owner, error) {
	r2 := <-ownerDatachannel
	var owner entity.Owner
	err := json.NewDecoder(r2.Body).Decode(&owner)
	if err != nil {
		fmt.Print(err.Error())
		return owner, err
	}
	return owner, nil
}
