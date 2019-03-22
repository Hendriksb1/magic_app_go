package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Your API details are below:
//Key:  50608513889ff85ea67f73e9773a0f10
//Secret: 30ddf0cb00004f14f99e88dc8c1de5cf
//Documentation can be found here: http://magicseaweed.com/developer/forecast-api
//Here's an example URL showing the forecast for Laga: http://magicseaweed.com/api/50608513889ff85ea67f73e9773a0f10/forecast/?spot_id=4401
//Please let us know if there are any problems.

type Result struct {
	Timestamp      int       `json:"timestamp"`
	LocalTimestamp int       `json:"localTimestamp"`
	IssueTimestamp int       `json:"issueTimestamp"`
	FadedRating    int       `json:"fadedRating"`
	SolidRating    int       `json:"solidRating"`
	Swell          Swell     `json:"swell"`
	Wind           Wind      `json:"wind"`
	Condition      Condition `json:"condition"`
	Charts         Charts    `json:"charts"`
}

type Swell struct {
	AbsMinBreakingHeight float64    `json:"absMinBreakingHeight"`
	AbsMaxBreakingHeight float64    `json:"absMaxBreakingHeight"`
	Probability          int        `json:"probability"`
	Unit                 string     `json:"unit"`
	MinBreakingHeight    float64    `json:"minBreakingHeight"`
	MaxBreakingHeight    float64    `json:"maxBreakingHeight"`
	Components           Components `json:"components"`
}

type Wind struct {
	Speed            int     `json:"speed"`
	Direction        float64 `json:"direction"`
	CompassDirection string  `json:"compassDirection"`
	Chill            int     `json:"chill"`
	Gusts            int     `json:"gusts"`
	Unit             string  `json:"unit"`
}

type Condition struct {
	Pressure     int    `json:"pressure"`
	Temperature  int    `json:"temperature"`
	Weather      string `json:"weather"`
	UnitPressure string `json:"unitPressure"`
	Unit         string `json:"unit"`
}

type Charts struct {
	Swell    string `json:"swell"`
	Period   string `json:"period"`
	Wind     string `json:"wind"`
	Pressure string `json:"pressure"`
	SST      string `json:"sst"`
}

type Components struct {
	Combined Component `json:"combined"`
	Primary  Component `json:"primary"`
}

type Component struct {
	Height           float64 `json:"height"`
	Period           int     `json:"period"`
	Direction        float64 `json:"direction"`
	CompassDirection string  `json:"compassDirection"`
}

func get_content() {
	url := "http://magicseaweed.com/api/50608513889ff85ea67f73e9773a0f10/forecast/?spot_id=4401"

	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	var data []Result

	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("Results: %v\n", data)

	///**********************
	// this creates an json file with magicseaweed data
	///**********************
	a := data

	out, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}

	fmt.Println("json was written")

	err = ioutil.WriteFile("../../../dev/websites/magic_seaweed/data/laga.json", []byte(string(out)), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	///**********************
}

func main() {
	get_content()
}
