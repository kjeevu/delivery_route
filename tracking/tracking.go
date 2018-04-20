package tracking

import (
	"delivery_route/utils"
	"fmt"
	"log"

	"github.com/kr/pretty"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

// Process Info  takes the file name as argument from the command line and file will parsed and
// Input is collected as array of Coordinates along with time stamp

func ProcessInfo(fileName string) {
	coOrdinates, err := utils.ReadCSV(fileName)
	if err != nil {
		log.Fatal(err)
	}
	totalDistance := utils.GetDistance(coOrdinates[0].Latitude, coOrdinates[len(coOrdinates)-1].Latitude, coOrdinates[0].Longitude, coOrdinates[len(coOrdinates)-1].Longitude)
	log.Println(totalDistance, "KM")
	log.Println("total Time", coOrdinates[len(coOrdinates)-1].TimeStamp-coOrdinates[0].TimeStamp)
	//Some clarifaction is required when data was  observed and differential data was collected
	// In some cases  the sum consecutive distance  is greater than  difference between start and end point
	//But  the diffrential time  is same as  the time difference between  start and end point
	// In some cases the consecutive  coordinates are same  but timestamp is diffrent
	// Whether traffic  has been considered in sample data

	/*	var time int64
		Longitude	var distance float64

		Longitude	for i := 0; i < len(coOrdinates)-1; i++ {
		Longitude		distance = distance + utils.GetDistance(coOrdinates[i].Latitude, coOrdinates[i+1].Latitude, coOrdinates[i].Longitude, coOrdinates[i+1].Longitude)
		Longitude		log.Println("Distance between ", i, "-->", i+1, utils.GetDistance(coOrdinates[i].Latitude, coOrdinates[i+1].Latitude, coOrdinates[i].Longitude, coOrdinates[i+1].Longitude))
		Longitude		log.Println("time difference between", i, "-->", i+1, coOrdinates[i+1].TimeStamp-coOrdinates[i].TimeStamp)
		Longitude		time = time + coOrdinates[i+1].TimeStamp - coOrdinates[i].TimeStamp
		Longitude	}
		Longitude	log.Println("Differential time", time)
		Longitude	log.Println("Differential distance", distance)
	*/
	OptimizeRoute(coOrdinates)
}

//TODO:   logic to remove the ambigious waypoints needs to be revisted
// OptimizeRoute get the all coordinates(Lattitude and langitude)
// we are using external google direction api(https://github.com/googlemaps/google-maps-services-go)
// googlemaps api will accept only 23 waypoints including source and destination
// hence we are splitting the the routes according to limits
// for every request we get the optimized route for 23 points
// map all the routes with source and destination as the key
// TODO: merge all the splitted routes and make single optimized route.

func OptimizeRoute(coOrdinates []utils.Coordinates) {
	points := make([]string, len(coOrdinates))
	for i := 0; i < len(coOrdinates); i++ {
		points[i] = fmt.Sprintf("%.10f,%.10f", coOrdinates[i].Latitude, coOrdinates[i].Longitude)
	}

	routeMap := make(map[string][]maps.Route)

	// here we are splitting the total waypoints multiple chunks
	// each chunk will have 23 waypoints including source and destination
	i := 0
	for i < len(points) {
		source := points[i]
		var destination string
		var waypoints []string

		if i+22 < len(points) {
			destination = points[i+22]
			waypoints = points[i+1 : i+21]
		} else {
			destination = points[len(points)-1]
			waypoints = points[i+1:]
		}
		i = i + 22

		route := getRoute(source, destination, waypoints)
		routeMap[fmt.Sprintf("%s,%s", source, destination)] = route
		pretty.Println(routeMap)
	}

}

// it accepts the source, destination and waypoints, returs the optimised route
func getRoute(source, destination string, waypoints []string) []maps.Route {

	c, err := maps.NewClient(maps.WithAPIKey(utils.API_KEY))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	r := &maps.DirectionsRequest{
		Origin:      source,
		Destination: destination,
		Waypoints:   waypoints,
		Optimize:    true,
	}

	route, points, err := c.Directions(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	pretty.Println(points)
	return route
}
