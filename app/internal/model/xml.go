package model

import "gorm.io/datatypes"

type Xml struct {
	ID               int    `json:"id"`
	UID              int    `json:"uid"`
	FirstName        string `json:"firstname,omitempty"`
	LastName         string `json:"lastname,omitempty"`
	Title            string `json:"title,omitempty"`
	SndType          string `json:"sdntype,omitempty"`
	ProgramList      datatypes.JSON
	IdList           datatypes.JSON
	AddressList      datatypes.JSON
	NationalityList  datatypes.JSON
	DataOfBirthList  datatypes.JSON
	PlaceOfBirthList datatypes.JSON
}
