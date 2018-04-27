 delivery_route


In this optmizing the routes between the source and destination using list of 
coordinates as input 

Set of  23 Waypoints from Source to Destination is  used as  input to the   external library api google direction (https://github.com/googlemaps/google-maps-services-go) 


The Optimized route method removes the ambigious coordinates and creates the 
 data which is saved in optimized_points.csv 

TODO: Implemention of Recurisive logic  to remove the  ambigious coordinates to avoid loops in the route
Combining the result of  all the sets  form a single route 
 


Usage:

File name or  absolute path of the file can be passed flag argument(-file_name  ) to the application

If file name is passed as argument it need to be there in same folder

If no argument is passed the File which is placed in same folder with the name 
points.csv will be Used as a input data 