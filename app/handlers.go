package app

import (
	"encoding/json"
	"encoding/xml"
	"github.com/dleond29/banking/service"
	"net/http"
)

type Customer struct {
	Name    string `json:"name,omitempty" xml:"name"`
	City    string `json:"city,omitempty" xml:"city"`
	Zipcode string `json:"zipcode,omitempty" xml:"zipcode"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	//customers := []Customer{
	//	{"Diego", "Giron", "68003"},
	//	{"Angie", "Madrid", "88001"},
	//}

	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}
