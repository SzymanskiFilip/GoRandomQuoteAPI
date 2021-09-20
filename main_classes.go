package main

type Quotes struct{
	Quotes []Quote `json:"quotes"`
}

type Quote struct{
	Person string `json:"person"`
	Quote string `json:"quote"`
}