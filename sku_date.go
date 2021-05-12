package main

type Result struct {
	Code     int      `json:"code"`
	ErrorMsg string   `json:"errorMsg"`
	Model    *SchInfo `json:"model"`
	Success  bool     `json:"success"`
}

type SchInfo struct {
	AmScheduleInfoDtoList []*Sku `json:"amScheduleInfoDtoList"`
	PmScheduleInfoDtoList []*Sku `json:"pmScheduleInfoDtoList"`
	Number                int    `json:"number"`
}

type Sku struct {
	Amount           string `json:"amount"`
	AmountTotal      string `json:"amountTotal"`
	AmountUsed       string `json:"amountUsed"`
	CategFlag        string `json:"categFlag"`
	EditTime         string `json:"editTime"`
	EmpCode          string `json:"empCode"`
	EmpId            string `json:"empId"`
	HaveDateInTheDay bool   `json:"haveDateInTheDay"`
	Levelname        string `json:"levelname"`
	OpenHours        string `json:"openHours"`
	PayType          string `json:"payType"`
	PriceDetail      []byte `json:"priceDetail"`
	ReplaceTime      string `json:"replaceTime"`
	RoomLocation     string `json:"roomLocation"`
	SkdDate          string `json:"skdDate"`
	SkdDateMD        string `json:"skdDateMD"`
	SkdSno           string `json:"skdSno"`
	Status           string `json:"status"`
	StopTime         string `json:"stopTime"`
	TotalCost        string `json:"totalCost"`
	Urpuls           string `json:"urpuls"`
	WhatDay          string `json:"whatDay"`
}
