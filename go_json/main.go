package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON video")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {

	MaverickCourses := []course{
		{"ReactJS Bootcamp", 299, "maverickpallav.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "maverickpallav.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "maverickpallav.in", "hit123", nil},
	}

	//package this data as JSON data

	finalJson, err := json.MarshalIndent(MaverickCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "maverickpallav.in",
		"tags": ["web-dev","js"]
	}
	`)

	var MaverickCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &MaverickCourse)
		fmt.Printf("%#v\n", MaverickCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v  and value is %v and Type is: %T\n", k, v, v)
	}

}
