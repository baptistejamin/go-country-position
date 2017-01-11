Go Country Position
==========

This package is a tiny library to convert a country code to a position

Telebot offers a convenient wrapper to Bots API, so you shouldn't even
bother about networking at all. You may install it with

## Install it

  go get github.com/mywaystar/go-country-position

(after setting up your `GOPATH` properly).

## Example

```go
import (
  CountryPosition "github.com/mywaystar/go-country-position"
  "fmt"
)

func main() {
  countryPosition := CountryPosition.New()

  //Lat[0], Lng[1]
  coordinates, _ := countryPosition.GetCoordinatesFromCode("GB")
  fmt.Printf("UK Coordinates (raw): %+v\n", coordinates)
  
  location, _ := countryPosition.GetCountryFromCode("FR")
  fmt.Printf("FR Location: %s\n", location)
}
```
