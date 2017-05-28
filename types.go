package main

import "encoding/json"

type ServerError struct {
	Error   error
	Message string
	Code    int
}

type Response struct {
	Status int		`json:"status"`
	Descr  string		`json:"descr"`
}

type DevConfig struct {
	TurnedOn	bool	`json:"turnedOn"`
	StreamOn	bool	`json:"streamOn"`
	CollectFreq	int32	`json:"collectFreq"`
	SendFreq	int32	`json:"sendFreq"`
}

type Metadata struct {
	Type	string	`json:"type"`
	Name	string	`json:"name"`
	MAC	string	`json:"mac"`
	IP	string	`json:"ip"`
}

type Request struct {
	Action	string		`json:"action"`
	Time	int64		`json:"time"`
	Meta	Metadata	`json:"meta"`
	Data	json.RawMessage	`json:"data"`
}

type FridgeData struct {
	TempCam1 map[int64]float32	`json:"tempCam1"`
	TempCam2 map[int64]float32	`json:"tempCam2"`
}

type WasherData struct {
	Mode	string
	Drying	string
	Temp	map[int64]float32
}

type DevData struct {
	Site	string			`json:"site"`
	Meta	Metadata		`json:"meta"`
	Data 	map[string][]string	`json:"data"`
}
