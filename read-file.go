package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("./impression")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)

	content := strings.Split(string(data), "\n")

	for _, placement := range content {
		placementImp := strings.Split(placement, "|")
		id, _ := strconv.ParseInt(placementImp[0], 10, 64)
		imp, _ := strconv.ParseInt(placementImp[1], 10, 64)

		fmt.Printf("ID: %d Impression: %d \n", id, imp)
	}
}
