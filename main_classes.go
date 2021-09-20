package main

type Quotes struct{
	Quotes []Quote `json:"quotes"`
}

type Quote struct{
	Person string
	Quote string
}