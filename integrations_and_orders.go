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

import (
	"encoding/json"
	"fmt"
)

// InsertingShipmentsBody is to structure the data
type InsertingShipmentsBody struct {
	Address                    string                                 `json:"address"`
	Address2                   string                                 `json:"address_2"`
	City                       string                                 `json:"city"`
	CompanyName                string                                 `json:"company_name"`
	Country                    string                                 `json:"country"`
	CreatedAt                  string                                 `json:"created_at"`
	Currency                   string                                 `json:"currency"`
	CustomsInvoiceNr           string                                 `json:"customs_invoice_nr"`
	CustomsShipmentType        interface{}                            `json:"customs_shipment_type"`
	Email                      string                                 `json:"email"`
	ExternalOrderId            string                                 `json:"external_order_id"`
	ExternalShipmentId         *string                                `json:"external_shipment_id"`
	HouseNumber                string                                 `json:"house_number"`
	Name                       string                                 `json:"name"`
	OrderNumber                string                                 `json:"order_number"`
	OrderStatus                *InsertingShipmentsBodyOrderStatus     `json:"order_status"`
	ParcelItems                []InsertingShipmentsBodyParcelItems    `json:"parcel_items"`
	PaymentStatus              *InsertingShipmentsBodyPaymentStatus   `json:"payment_status"`
	PostalCode                 string                                 `json:"postal_code"`
	SenderAddress              int                                    `json:"sender_address,omitempty"`
	ShippingMethod             int                                    `json:"shipping_method,omitempty"`
	ShippingMethodCheckoutName string                                 `json:"shipping_method_checkout_name"`
	Telephone                  string                                 `json:"telephone"`
	ToPostNumber               string                                 `json:"to_post_number"`
	ToServicePoint             interface{}                            `json:"to_service_point"`
	ToState                    interface{}                            `json:"to_state"`
	TotalOrderValue            string                                 `json:"total_order_value,omitempty"`
	UpdatedAt                  string                                 `json:"updated_at"`
	Weight                     string                                 `json:"weight"`
	Width                      string                                 `json:"width,omitempty"`
	Height                     string                                 `json:"height,omitempty"`
	Length                     string                                 `json:"length,omitempty"`
	Shipments                  []interface{}                          `json:"shipments,omitempty"`
	CheckoutPayload            *InsertingShipmentsBodyCheckoutPayload `json:"checkout_payload,omitempty"`
}

type InsertingShipmentsBodyOrderStatus struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

type InsertingShipmentsBodyParcelItems struct {
	Description   string                 `json:"description"`
	HsCode        string                 `json:"hs_code"`
	OriginCountry string                 `json:"origin_country"`
	ProductId     string                 `json:"product_id"`
	Properties    map[string]interface{} `json:"properties"`
	Quantity      int                    `json:"quantity"`
	Sku           string                 `json:"sku"`
	Value         string                 `json:"value"`
	Weight        string                 `json:"weight"`
}

type InsertingShipmentsBodyPaymentStatus struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

type InsertingShipmentsBodyCheckoutPayload struct {
	SenderAddressId    int                                      `json:"sender_address_id"`
	ShippingProduct    InsertingShipmentsBodyShippingProduct    `json:"shipping_product"`
	DeliveryMethodType string                                   `json:"delivery_method_type"`
	DeliveryMethodData InsertingShipmentsBodyDeliveryMethodData `json:"delivery_method_data"`
}

type InsertingShipmentsBodyShippingProduct struct {
	Code                    string                                        `json:"code"`
	Name                    string                                        `json:"name"`
	SelectedFunctionalities InsertingShipmentsBodySelectedFunctionalities `json:"selected_functionalities"`
}

type InsertingShipmentsBodySelectedFunctionalities struct {
	AgeCheck           int    `json:"age_check"`
	LastMile           string `json:"last_mile"`
	FirstMile          string `json:"first_mile"`
	Multicollo         bool   `json:"multicollo"`
	FormFactor         string `json:"form_factor"`
	ServiceArea        string `json:"service_area"`
	WeekendDelivery    string `json:"weekend_delivery"`
	DeliveryDeadline   string `json:"delivery_deadline"`
	DirectContractOnly bool   `json:"direct_contract_only"`
}

type InsertingShipmentsBodyDeliveryMethodData struct {
	DeliveryDate          string `json:"delivery_date"`
	FormattedDeliveryDate string `json:"formatted_delivery_date"`
	ParcelHandoverDate    string `json:"parcel_handover_date"`
}

// InsertingShipmentsReturn is to decode the json return
type InsertingShipmentsReturn struct {
	ExternalOrderId    string `json:"external_order_id"`
	ExternalShipmentId string `json:"external_shipment_id"`
	ShipmentUuid       string `json:"shipment_uuid,omitempty"`
	Status             string `json:"status"`
	Error              struct {
		Address []string `json:"address"`
	} `json:"error,omitempty"`
}

// InsertingShipmentsWebhookReturn is to decode the json data
type InsertingShipmentsWebhookReturn struct {
	Action string `json:"action"`
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
		DateCreated string `json:"date_created"`
		Email       string `json:"email"`
		Name        string `json:"name"`
		PostalCode  string `json:"postal_code"`
		Reference   string `json:"reference"`
		Shipment    struct {
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
		OrderNumber         string      `json:"order_number"`
		InsuredValue        int         `json:"insured_value"`
		TotalInsuredValue   int         `json:"total_insured_value"`
		ToState             interface{} `json:"to_state"`
		CustomsInvoiceNr    string      `json:"customs_invoice_nr"`
		CustomsShipmentType interface{} `json:"customs_shipment_type"`
		ParcelItems         []struct {
			Description   string      `json:"description"`
			HsCode        string      `json:"hs_code"`
			ProductId     string      `json:"product_id"`
			Quantity      int         `json:"quantity"`
			Weight        string      `json:"weight"`
			Sku           string      `json:"sku"`
			ReturnReason  interface{} `json:"return_reason"`
			ReturnMessage interface{} `json:"return_message"`
			Value         string      `json:"value"`
			Properties    struct {
				Size  string `json:"Size"`
				Color string `json:"Color"`
				Cut   string `json:"Cut,omitempty"`
			} `json:"properties"`
		} `json:"parcel_items"`
		Documents []struct {
			Type string `json:"type"`
			Size string `json:"size"`
			Link string `json:"link"`
		} `json:"documents"`
		Type                       string      `json:"type"`
		ShipmentUuid               string      `json:"shipment_uuid"`
		ShippingMethod             int         `json:"shipping_method"`
		ShippingMethodCheckoutName interface{} `json:"shipping_method_checkout_name"`
		ExternalOrderId            string      `json:"external_order_id"`
		ExternalShipmentId         string      `json:"external_shipment_id"`
		ExternalReference          interface{} `json:"external_reference"`
		IsReturn                   bool        `json:"is_return"`
		Note                       interface{} `json:"note"`
		ToPostNumber               string      `json:"to_post_number"`
		TotalOrderValue            string      `json:"total_order_value"`
		TotalOrderValueCurrency    string      `json:"total_order_value_currency"`
		Carrier                    struct {
			Code                     string `json:"code"`
			Name                     string `json:"name"`
			ServicepointsCarrierCode string `json:"servicepoints_carrier_code"`
		} `json:"carrier"`
		TrackingUrl             string      `json:"tracking_url"`
		DateUpdated             string      `json:"date_updated"`
		DateAnnounced           string      `json:"date_announced"`
		ColliTrackingNumber     string      `json:"colli_tracking_number"`
		ColliUuid               string      `json:"colli_uuid"`
		ColloNr                 int         `json:"collo_nr"`
		ColloCount              int         `json:"collo_count"`
		AwbTrackingNumber       interface{} `json:"awb_tracking_number"`
		BoxNumber               interface{} `json:"box_number"`
		Length                  string      `json:"length"`
		Width                   string      `json:"width"`
		Height                  string      `json:"height"`
		TrackTraceNotifications bool        `json:"track_trace_notifications"`
		FromCountry             struct {
			Iso2 string `json:"iso_2"`
			Iso3 string `json:"iso_3"`
			Name string `json:"name"`
		} `json:"from_country"`
		FromPostalCode string `json:"from_postal_code"`
		Contract       struct {
			Type string `json:"type"`
		} `json:"contract"`
		DateShipped        interface{}   `json:"date_shipped"`
		HouseNumber        string        `json:"house_number"`
		SuppressedStatuses []interface{} `json:"suppressed_statuses"`
		ShouldSendFeedback bool          `json:"should_send_feedback"`
		Source             string        `json:"source"`
	} `json:"parcel"`
	Timestamp int64 `json:"timestamp"`
}

// InsertingShipments is to insert a shipment into an api integration
func InsertingShipments(id int, body []InsertingShipmentsBody, r Request) ([]InsertingShipmentsReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	// Set config for new request
	c := Config{
		Path:   fmt.Sprintf("/integrations/%d/shipments", id),
		Method: "POST",
		Body:   convert,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return nil, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode []InsertingShipmentsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return nil, err
	}

	// Return data
	return decode, err

}
