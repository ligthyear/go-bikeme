package bikeshareservice

import (
	"bytes"
	"encoding/xml"
	"go-bikeme/station"
	"net/http"
)

const CAPITAL_BIKE_SHARE_URL string = "http://www.capitalbikeshare.com/data/stations/bikeStations.xml"

type capitalBikeShareService struct {
	baseService
}

func NewCapitalBikeShareService() (*capitalBikeShareService) {
	service := capitalBikeShareService{}
	service.serviceImpl = &service
	return &service
}

func (service *capitalBikeShareService) queryService() (response *http.Response, err error) {
	return http.Get(CAPITAL_BIKE_SHARE_URL)
}

func (service *capitalBikeShareService) parse(capitalbikshareXML []byte) (stations []station.Station, err error) {
	xmlStations := &XMLStations{}

	xml.NewDecoder(bytes.NewReader(capitalbikshareXML)).Decode(xmlStations)

	for _, xmlStation := range xmlStations.XMLStations {
		stations = append(stations, service.createStation(xmlStation))
	}

	return
}

func (service *capitalBikeShareService) createStation(xmlStation XMLStation) station.Station {
	stationObject := station.Station{}
	stationObject.StationId = xmlStation.Id
	stationObject.StationName = xmlStation.Name

	stationObject.Status = station.Status{xmlStation.NbBikes, xmlStation.NbEmptyDocks}

	return stationObject
}

type XMLStations struct {
	XMLName     xml.Name     `xml:"stations"`
	XMLStations []XMLStation `xml:"station"`
}

type XMLStation struct {
	XMLName      xml.Name `xml:"station"`
	Id           string   `xml:"id"`
	Name         string   `xml:"name"`
	Lat          float32  `xml:"lat"`
	Long         float32  `xml:"long"`
	NbBikes      int64    `xml:"nbBikes"`
	NbEmptyDocks int64    `xml:"nbEmptyDocks"`
}
