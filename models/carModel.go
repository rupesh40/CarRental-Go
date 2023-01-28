package models

import (
	//"time"
)
type Car struct{
	ID		  			 string 	`json:"id"`
	CarModel	  	     string 	`json:"carModel" validate:"required"`	
	DateOfManufacture	 string 	`json:"date_of_manufacture" validate:"required"`
	LastServiceDate	 	 string 	`json:"last_service_date"`
	LastUsedDate		 string 	`json:"last_used_date"`
	Status 				 string		`json:"status"`
	// `json:"phone" validate:"required"`
}