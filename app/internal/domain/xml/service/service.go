package service

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"project/internal/domain/xml/storage"
	models "project/internal/model"
	"strconv"

	"gorm.io/datatypes"
)

type XmlService interface {
	All() []*models.Xml
	Import() error
}

type xmlService struct {
	storage storage.XmlStorage
}

// ----------------------------------------------------------------------------------------xml structs start
type SdnList struct {
	XMLName  xml.Name   `xml:"sdnList"`
	SdnEntry []SdnEntry `xml:"sdnEntry"`
}

type SdnEntry struct {
	UID              string           `xml:"uid"`
	FirstName        string           `xml:"firstName"`
	LastName         string           `xml:"lastName"`
	Title            string           `xml:"title"`
	SdnType          string           `xml:"sdnType"`
	ProgramList      ProgramList      `xml:"programList"`
	IdList           IdList           `xml:"idList"`
	AddressList      AddressList      `xml:"addressList"`
	NationalityList  NationalityList  `xml:"nationalityList"`
	DateOfBirthList  DateOfBirthList  `xml:"dateOfBirthList"`
	PlaceOfBirthList PlaceOfBirthList `xml:"placeOfBirthList"`
}

type ProgramList struct {
	Program string `xml:"program"`
}

type IdList struct {
	Id []Id `xml:"id"`
}

type Id struct {
	UID      string `xml:"uid"`
	IdType   string `xml:"idType"`
	IdNumber string `xml:"idNumber"`
}

type AddressList struct {
	Address []Address `xml:"address"`
}

type Address struct {
	UID     string `xml:"uid"`
	City    string `xml:"city"`
	Country string `xml:"country"`
}

type NationalityList struct {
	Nationality []Nationality `xml:"nationality"`
}

type Nationality struct {
	UID       string `xml:"uid"`
	Country   string `xml:"country"`
	MainEntry string `xml:"mainEntry"`
}

type DateOfBirthList struct {
	DateOfBirthItem []DateOfBirthItem `xml:"dateOfBirthItem"`
}

type DateOfBirthItem struct {
	UID         string `xml:"uid"`
	DateOfBirth string `xml:"dateOfBirth"`
	MainEntry   string `xml:"mainEntry"`
}

type PlaceOfBirthList struct {
	PlaceOfBirthItem []PlaceOfBirthItem `xml:"placeOfBirthItem"`
}

type PlaceOfBirthItem struct {
	UID          string `xml:"uid"`
	PlaceOfBirth string `xml:"placeOfBirth"`
	MainEntry    string `xml:"mainEntry"`
}

// ----------------------------------------------------------------------------------------xml structs end

func NewXmlService(storage storage.XmlStorage) XmlService {
	return &xmlService{
		storage: storage,
	}
}

func (s *xmlService) All() []*models.Xml {
	return s.storage.All()
}

func (s *xmlService) Import() error {
	url := "https://www.treasury.gov/ofac/downloads/sdn.xml"

	// Fetch the XML data from the URL
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read the response body into a byte array
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Unmarshal the XML data into a SdnList struct
	var sdnList SdnList
	err = xml.Unmarshal(body, &sdnList)
	if err != nil {
		return err
	}
	// Print the parsed data
	for _, sdnEntry := range sdnList.SdnEntry {
		idlist := []Id{}
		for _, id := range sdnEntry.IdList.Id {
			uid := id.IdType
			idnumber := id.IdNumber
			p := Id{UID: uid, IdNumber: idnumber}
			idlist = append(idlist, p)
		}
		// ----------------------------------------------------------------
		address := []Address{}
		for _, id := range sdnEntry.AddressList.Address {
			uid := id.UID
			city := id.City
			country := id.Country
			p := Address{UID: uid, City: city, Country: country}
			address = append(address, p)
		}
		// ----------------------------------------------------------------
		nationality := []Nationality{}
		for _, id := range sdnEntry.NationalityList.Nationality {
			uid := id.UID
			country := id.Country
			mainEntry := id.MainEntry
			p := Nationality{UID: uid, Country: country, MainEntry: mainEntry}
			nationality = append(nationality, p)
		}
		// ----------------------------------------------------------------
		dateOfBirthItem := []DateOfBirthItem{}
		for _, id := range sdnEntry.DateOfBirthList.DateOfBirthItem {
			uid := id.UID
			dateOfBirth := id.DateOfBirth
			mainEntry := id.MainEntry
			p := DateOfBirthItem{UID: uid, DateOfBirth: dateOfBirth, MainEntry: mainEntry}
			dateOfBirthItem = append(dateOfBirthItem, p)
		}
		// ----------------------------------------------------------------
		placeOfBirthItem := []PlaceOfBirthItem{}
		for _, id := range sdnEntry.PlaceOfBirthList.PlaceOfBirthItem {
			uid := id.UID
			dateOfBirth := id.PlaceOfBirth
			mainEntry := id.MainEntry
			p := PlaceOfBirthItem{UID: uid, PlaceOfBirth: dateOfBirth, MainEntry: mainEntry}
			placeOfBirthItem = append(placeOfBirthItem, p)
		}
		// -----------------------------------------------------------------------------------------------------------
		marksStr := sdnEntry.UID
		marks, err := strconv.Atoi(marksStr)

		if err != nil {
			return err
		}
		xml := &models.Xml{
			UID:              marks,
			FirstName:        sdnEntry.FirstName,
			LastName:         sdnEntry.LastName,
			Title:            sdnEntry.Title,
			SndType:          sdnEntry.SdnType,
			ProgramList:      datatypes.JSON(`{"Program": "` + fmt.Sprintf("%v", sdnEntry.ProgramList.Program) + `"}`),
			IdList:           datatypes.JSON(`{"Id": "` + fmt.Sprintf("%v", idlist) + `"}`),
			AddressList:      datatypes.JSON(`{"address": "` + fmt.Sprintf("%v", address) + `"}`),
			NationalityList:  datatypes.JSON(`{"nationality": "` + fmt.Sprintf("%v", nationality) + `"}`),
			DataOfBirthList:  datatypes.JSON(`{"dateOfBirthItem": "` + fmt.Sprintf("%v", dateOfBirthItem) + `"}`),
			PlaceOfBirthList: datatypes.JSON(`{"placeOfBirthItem": "` + fmt.Sprintf("%v", placeOfBirthItem) + `"}`),
		}

		err = s.storage.CreateOrUpdate(xml, marks)
		if err != nil {
			return err
		}

	}
	return nil
}
