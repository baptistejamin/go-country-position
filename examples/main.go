// Copyright 2016 Crisp IM, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
  CountryPosition "github.com/mywaystar/go-country-position"
  "fmt"
)

func main() {
  countryPosition := CountryPosition.New()

  coordinates, err := countryPosition.GetCoordinatesFromCode("GB")

  if err != nil {
    fmt.Printf("Error: %s", err)
  } else {
    fmt.Printf("UK Coordinates (raw): %+v\n", coordinates)
  }

  location, err := countryPosition.GetCountryFromCode("FR")

  if err != nil {
    fmt.Printf("Error: %s", err)
  } else {
    fmt.Printf("FR Location: %s\n", location)
  }
}