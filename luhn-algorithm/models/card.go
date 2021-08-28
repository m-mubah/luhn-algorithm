package models

type Card struct {
	ImgPath           string
	CardNumber        string
	MII               string
	Issuer            string
	PersonalAccNumber string
	Checksum          int
}
