package main

import (
	"fmt"
	//"github.com/mitchellh/mapstructure"
)

// range, map
func main() {
	//foreach() // foreach example
	//exampleMaps()

	isKeyExists("a")
	isKeyExists("d")

	// decoding map in structure
	/*	pointsMap := map[string]int{
		"x": 1,
		"y": 2,
	}*/

	// convert map in structure
	//mapstructure.Decode(pointsMap)
}

type Point struct {
	X, Y int
}

// Point method
func (p Point) method() {
	fmt.Println(p.X, p.Y)
}

// practice with map
func exampleMaps() {
	// map initialize
	pointsMap := map[string]Point{
		"b": {
			Y: 11,
			X: 8,
		},
	} // string - key, Point - value
	otherMap := map[string]int{}          // string - key, int - value
	otherPointsMap := make(map[int]Point) // int - key, Point - value
	pointsMap["a"] = Point{
		Y: 2,
		X: 1,
	}
	otherPointsMap[1] = Point{1, 2}

	fmt.Println(pointsMap)         // map[a:{1 2} b:{8 11}]
	fmt.Println(pointsMap["a"])    // {1 2}
	fmt.Println(otherMap)          // map[]
	fmt.Println(otherPointsMap[1]) // {1 2}

	// not initialized
	var oneMorePointsMap map[string]Point // string - key, Point - value
	if oneMorePointsMap == nil {
		fmt.Println("oneMorePointsMap is nil")
		oneMorePointsMap = map[string]Point{
			"a": {
				5,
				7,
			},
		}
	} else {
		oneMorePointsMap["a"] = Point{5, 7}
	}
	fmt.Println(oneMorePointsMap)  // map[a:{5 7}]
	oneMorePointsMap["a"].method() // 5 7
}

// check if exist key in map
func isKeyExists(key string) {
	pointsMap := map[string]Point{
		"a": {
			1,
			2,
		},
	}

	value, ok := pointsMap[key]
	if ok {
		fmt.Printf("key = %s exist in map\n", key) // key = a exist in map
		fmt.Println(value)                         // {1 2}
	} else {
		fmt.Printf("key = %s does not exist in map\n", key) // key = d does not exist in map
		fmt.Println(value)                                  // {0 0}
	}
}

// foreach example
func foreach() {
	arr := []string{"a", "b", "c"}
	for i, l := range arr {
		fmt.Println(i, l) // 0 a, 1 b, 2 c
	}

	// other way (if you need only value)
	for _, l := range arr {
		fmt.Println(l) // a b c
	}

	// if you need only index
	for i := range arr {
		fmt.Println(i) // 0 1 2
	}
}
