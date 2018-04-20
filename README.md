 delivery_route


In this optmizing the routes between the source and destination using list of 
coordinates as input 

Used external library which google direction api(https://github.com/googlemaps/google-maps-services-go)

right now splitting the routes according to limits, for every request we get the optimized route for 23 points

TODO: merge all the splitted routes and make single optimized route and
Logic to remove  erroneous/ambiguous coordinates needs to be revisited
Implemetion of Custom Erors

Usage:

File name or  absolute path of the file can be passed flag argument(-file_name  ) to the application

If file name is passed as argument it need to be there in same folder

If no argument is passed the File which is placed in same folder with the name 
points.csv will be Used as a input data 