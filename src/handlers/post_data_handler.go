package handlers

import (
	"github.com/ugorji/go/codec"
	"go-api/src/models"
	"go-api/src/repository"
	"encoding/json"
	"log"
)

func PostCustomerCompanies(result [][]string, r *repository.Repository) error {
	return nil
}

func PostOrderItemsData(result [][]string, r *repository.Repository) error {
	return nil
}

func PostDeliveriesData(result [][]string, r *repository.Repository) error {
	return nil
}

func PostCustomerData(result [][]string, r *repository.Repository) error {
	var row []string
	var customers []models.Customers
	var customer models.Customers
	for i := 1 ; i<len(result) ; i++ {
		row = DecodeFileToString(result[i])

		b, err := json.Marshal(row)
		if err != nil {
			log.Fatal("\nERROR while marshaling the data ", err)
		}
		codec.NewDecoderBytes(b, new(codec.JsonHandle)).Decode(&customer)
		customers = append(customers, customer)
	}

	err := r.CreateCustomers(customers)
	return err
}

func PostOrderData(result [][]string, r *repository.Repository) error {
	var row []string
	var orders []models.Orders
	var order models.Orders
	for i := 1 ; i<len(result) ; i++ {
		row = DecodeFileToString(result[i])

		b, err := json.Marshal(row)
		if err != nil {
			log.Fatal("\nERROR while marshaling the data ", err)
		}
		codec.NewDecoderBytes(b, new(codec.JsonHandle)).Decode(&order)
		orders = append(orders, order)
	}

	err := r.CreateOrders(orders)
	return err
}
