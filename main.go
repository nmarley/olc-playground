package main

import (
	"fmt"
	"io/ioutil"

	olc "github.com/google/open-location-code/go"
	yaml "gopkg.in/yaml.v2"
)

// Place ...
type Place struct {
	Name    string `yaml:"place"`
	Address string `yaml:"addr"`
	Code    string `yaml:"code"`
}

func readYAMLConfig(filename string) ([]Place, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var places []Place
	err = yaml.Unmarshal(bytes, &places)
	if err != nil {
		return nil, err
	}

	return places, nil
}

func main() {
	places, err := readYAMLConfig("places.yaml")
	if err != nil {
		panic(err)
	}

	for _, place := range places {
		codeArea, err := olc.Decode(place.Code)
		if err != nil {
			panic(err)
		}
		lat, long := codeArea.Center()
		fmt.Printf("place: %v, lat: %0.4f, long: %0.4f\n", place.Name, lat, long)
	}

	// homePlusCode := "584H8GP7+8J"
	// codeArea, err := olc.Decode(homePlusCode)
	// if err != nil {
	// 	panic(err)
	// }
	// lat, long := codeArea.Center()
	// fmt.Printf("lat: %0.4f, long: %0.4f\n", lat, long)
}
