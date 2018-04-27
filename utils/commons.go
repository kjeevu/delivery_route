package utils

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

// Coordinates represents the lattitude , langitude and timestamp
type Coordinates struct {
	Latitude  float64
	Longitude float64
	TimeStamp int64
}

// ReadCSV reads the coordinates from the input csv file
func ReadCSV(filePath string) ([]Coordinates, error) {
	coOrdinates := make([]Coordinates, 0)
	csvFile, err := os.Open(filePath)
	if err != nil {
		err = errors.New("UnableToReadInputFile")
		return coOrdinates, err
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		var ltemp Coordinates
		ltemp.Latitude, err = strconv.ParseFloat(line[0], 64)
		if err != nil {
			err = errors.New("InvalidDataForLatitude")
			return coOrdinates, err
		}
		ltemp.Longitude, err = strconv.ParseFloat(line[1], 64)
		if err != nil {
			err = errors.New("InvalidDataForLangitude")
			return coOrdinates, err
		}
		ltemp.TimeStamp, err = strconv.ParseInt(line[2], 10, 64)
		if err != nil {

			err = errors.New("InvalidDataForLangitude")
			return coOrdinates, err
		}
		coOrdinates = append(coOrdinates, ltemp)
	}
	if len(coOrdinates) <= 0 {
		err = errors.New("InvalidData")
		return coOrdinates, err

	}
	return coOrdinates, nil
}

// Get Distance uses the Haversine formula to find the distance given coordinates in KM
func GetDistance(lat1, lat2, long1, long2 float64) float64 {
	s1, c1 := math.Sincos(GetRadians(lat1))
	s2, c2 := math.Sincos(GetRadians(lat2))
	clong := math.Cos(GetRadians(long1 - long2))
	return RADIUS * math.Acos(s1*s2+c1*c2*clong)
}

//To find Radian from degree
func GetRadians(degree float64) float64 {
	return degree * math.Pi / 180
}
func WriteCSV(coOrdinates []Coordinates) (err error) {
	file, err := os.Create("updated_points.csv")
	if err != nil {
		err = errors.New("UnableToCreateOutputFile")
		return
	}
	defer file.Close()
	//data := make([]string, len(coOrdinates))
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for i := 0; i < len(coOrdinates); i++ {
		line := []string{fmt.Sprintf("%.10f", coOrdinates[i].Latitude), fmt.Sprintf("%.10f", coOrdinates[i].Longitude), fmt.Sprintf("%d", coOrdinates[i].TimeStamp)}
		err = writer.Write(line)
		if err != nil {
			err = errors.New("UnableToWriteFile")
			return
		}
	}
	return
}
