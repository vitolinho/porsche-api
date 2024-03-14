package database

import (
	"porsche-api/internal/domain/model"
	"porsche-api/internal/domain/services"
)

func SeedCar() {
	var count int64
	DB.Model(&model.Car{}).Count(&count)

	if count == 0 {
		cars := []model.Car{
			{ Name: "Porsche 911", Price: services.GeneratePrice() },
			{ Name: "Porsche Cayman", Price: services.GeneratePrice() },
			{ Name: "Porsche Boxster", Price: services.GeneratePrice() },
			{ Name: "Porsche Panamera", Price: services.GeneratePrice() },
			{ Name: "Porsche Macan", Price: services.GeneratePrice() },
			{ Name: "Porsche Cayenne", Price: services.GeneratePrice() },
			{ Name: "Porsche 718", Price: services.GeneratePrice() },
			{ Name: "Porsche Taycan", Price: services.GeneratePrice() },
		}
		for _, car := range cars {
			DB.Create(&car)
		}
	}
}
