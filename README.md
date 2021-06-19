# Random Location Generator
Generate random locations around a certain location radius

## Install

````
go get github.com/anjulapaulus/random-location-generator
````

## Example
The radius is in metres.
````
import github.com/anjulapaulus/random-location-generator

func main(){
    loc := RandomLocation(
        Location{
            Latitude:6.540076897187259,
            Longitude:80.1094552048789,
        }, 100
    )
    fmt.Println(loc)
}
````