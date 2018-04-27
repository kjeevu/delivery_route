package tracking

import (
	"delivery_route/utils"
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
	flightDistance := utils.GetDistance(coOrdinates[0].Latitude, coOrdinates[len(coOrdinates)-1].Latitude, coOrdinates[0].Longitude, coOrdinates[len(coOrdinates)-1].Longitude)
	log.Println(flightDistance, "KM")
	var errnoeusindex []int
	errnoeusindex = OptimizeRoute(coOrdinates, flightDistance)
	if len(errnoeusindex) == 0 {
		err = utils.WriteCSV(coOrdinates)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("No Erroneus coordinates found")
		return
	}
	optimizedData := make([]utils.Coordinates, 0)
	var j, k = 0, 0
	for i := 0; i < len(coOrdinates); i++ {
		if i == errnoeusindex[j] && j != len(errnoeusindex)-1 {
			j++
		} else {
			optimizedData = append(optimizedData, coOrdinates[i])
			k++
		}
	}
	//OptimizeRoute(coOrdinates)
	err = utils.WriteCSV(coOrdinates)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Input Data is updated and saved in optimized_points.csv")
	return
}

//func OptimizeRoute(coOrdinates []utils.Coordinates) {
//	points := make([]string, len(coOrdinates))
//	for i := 0; i < len(coOrdinates); i++ {
//		points[i] = fmt.Sprintf("%.10f,%.10f", coOrdinates[i].Latitude, coOrdinates[i].Longitude)
//	}

//	routeMap := make(map[string][]maps.Route)

//	// here we are splitting the total waypoints multiple chunks
//	// each chunk will have 23 waypoints including source and destination
//	i := 0
//	for i < len(points) {
//		source := points[i]
//		var destination string
//		var waypoints []string

//		if i+22 < len(points) {
//			destination = points[i+22]
//			waypoints = points[i+1 : i+21]
//		} else {
//			destination = points[len(points)-1]
//			waypoints = points[i+1:]
//		}
//		i = i + 22

//		route := getRoute(source, destination, waypoints)
//		routeMap[fmt.Sprintf("%s,%s", source, destination)] = route
//		pretty.Println(routeMap)
//	}

//}

// OPtimizeRoute uses the distance of concurrent three  coordinates  and checks it with
//The flight distance  and find the   ambigous coordinates
// TODO : Insted of using concurrent three coordinates  devlopment of recursive logic
//To find the loops in the  route and remove the loops
func OptimizeRoute(coOrdinates []utils.Coordinates, flightDistance float64) (errnoeusindex []int) {
	i := 0
	errnoeusindex = make([]int, 0)
	// Recursive logic  needs to be implemented
	for {
		if i >= len(coOrdinates)-2 {
			break
		}
		distance1 := utils.GetDistance(coOrdinates[i].Latitude, coOrdinates[i+1].Latitude, coOrdinates[i].Longitude, coOrdinates[i+1].Longitude)
		distance2 := utils.GetDistance(coOrdinates[i+1].Latitude, coOrdinates[i+2].Latitude, coOrdinates[i+1].Longitude, coOrdinates[i+2].Longitude)
		if distance1+distance2 >= (.5 * flightDistance) {
			errnoeusindex = append(errnoeusindex, i+1)
			i = i + 1
		}
		i++

	}
	return
}

// it accepts the source, destination and waypoints, provides the result with considering or
// rejecting the way point  with status
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
