//********************************************************************************************************************//
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
//
// This file is part of gosendcloud.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor (aka gowizzard)
//
//********************************************************************************************************************//

package gosendcloud

import "encoding/json"

// CreateAParcelBody is to structure the json data
type CreateAParcelBody struct {
	Parcel CreateAParcelBodyParcel `json:"parcel"`
}

type CreateAParcelBodyParcel struct {
	Name                       string                    `json:"name"`
	CompanyName                string                    `json:"company_name"`
	Address                    string                    `json:"address"`
	HouseNumber                string                    `json:"house_number"`
	City                       string                    `json:"city"`
	PostalCode                 string                    `json:"postal_code"`
	Telephone                  string                    `json:"telephone"`
	RequestLabel               bool                      `json:"request_label"`
	Email                      string                    `json:"email"`
	Data                       []interface{}             `json:"data,omitempty"`
	Country                    string                    `json:"country"`
	Shipment                   CreateAParcelBodyShipment `json:"shipment"`
	Weight                     string                    `json:"weight"`
	OrderNumber                string                    `json:"order_number"`
	InsuredValue               int                       `json:"insured_value"`
	TotalOrderValueCurrency    string                    `json:"total_order_value_currency"`
	TotalOrderValue            string                    `json:"total_order_value"`
	Quantity                   int                       `json:"quantity"`
	ShippingMethodCheckoutName string                    `json:"shipping_method_checkout_name"`
}

type CreateAParcelBodyShipment struct {
	Id int `json:"id"`
}

// CreateAParcelReturn is to decode the json return
type CreateAParcelReturn struct {
	Parcel struct {
		Id             int    `json:"id"`
		Address        string `json:"address"`
		Address2       string `json:"address_2"`
		AddressDivided struct {
			Street      string `json:"street"`
			HouseNumber string `json:"house_number"`
		} `json:"address_divided"`
		City        string `json:"city"`
		CompanyName string `json:"company_name"`
		Country     struct {
			Iso2 string `json:"iso_2"`
			Iso3 string `json:"iso_3"`
			Name string `json:"name"`
		} `json:"country"`
		Data struct {
		} `json:"data"`
		DateCreated   string `json:"date_created"`
		DateAnnounced string `json:"date_announced"`
		DateUpdated   string `json:"date_updated"`
		Email         string `json:"email"`
		Name          string `json:"name"`
		PostalCode    string `json:"postal_code"`
		Reference     string `json:"reference"`
		Shipment      struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"shipment"`
		Status struct {
			Id      int    `json:"id"`
			Message string `json:"message"`
		} `json:"status"`
		ToServicePoint interface{} `json:"to_service_point"`
		Telephone      string      `json:"telephone"`
		TrackingNumber string      `json:"tracking_number"`
		Weight         string      `json:"weight"`
		Label          struct {
			NormalPrinter []string `json:"normal_printer"`
			LabelPrinter  string   `json:"label_printer"`
		} `json:"label"`
		CustomsDeclaration struct {
		} `json:"customs_declaration"`
		OrderNumber         string        `json:"order_number"`
		InsuredValue        int           `json:"insured_value"`
		TotalInsuredValue   int           `json:"total_insured_value"`
		ToState             interface{}   `json:"to_state"`
		CustomsInvoiceNr    string        `json:"customs_invoice_nr"`
		CustomsShipmentType interface{}   `json:"customs_shipment_type"`
		ParcelItems         []interface{} `json:"parcel_items"`
		Documents           []struct {
			Type string `json:"type"`
			Size string `json:"size"`
			Link string `json:"link"`
		} `json:"documents"`
		Type                       string      `json:"type"`
		ShipmentUuid               interface{} `json:"shipment_uuid"`
		ShippingMethod             int         `json:"shipping_method"`
		ExternalOrderId            string      `json:"external_order_id"`
		ExternalShipmentId         string      `json:"external_shipment_id"`
		ExternalReference          interface{} `json:"external_reference"`
		IsReturn                   bool        `json:"is_return"`
		Note                       string      `json:"note"`
		ToPostNumber               string      `json:"to_post_number"`
		TotalOrderValue            string      `json:"total_order_value"`
		TotalOrderValueCurrency    string      `json:"total_order_value_currency"`
		ColliTrackingNumber        string      `json:"colli_tracking_number"`
		ColliUuid                  string      `json:"colli_uuid"`
		ColloNr                    int         `json:"collo_nr"`
		ColloCount                 int         `json:"collo_count"`
		AwbTrackingNumber          interface{} `json:"awb_tracking_number"`
		BoxNumber                  interface{} `json:"box_number"`
		Length                     interface{} `json:"length"`
		Width                      interface{} `json:"width"`
		Height                     interface{} `json:"height"`
		ShippingMethodCheckoutName string      `json:"shipping_method_checkout_name"`
		Carrier                    struct {
			Code string `json:"code"`
		} `json:"carrier"`
		TrackingUrl string `json:"tracking_url"`
	} `json:"parcel"`
	Error struct {
		Code    int    `json:"code"`
		Request string `json:"request"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

// CreateAParcel is to create a new parcel
func CreateAParcel(body CreateAParcelBody, r Request) (CreateAParcelReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return CreateAParcelReturn{}, err
	}

	// Set config for new request
	c := Config{
		Path:   "/parcels",
		Method: "POST",
		Body:   convert,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return CreateAParcelReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode CreateAParcelReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return CreateAParcelReturn{}, err
	}

	// Return data
	return decode, err

}
