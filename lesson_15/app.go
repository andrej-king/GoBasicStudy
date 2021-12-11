package main

import (
	"encoding/json"
	"fmt"
)

// work with json
// libs for work with json:
// https://github.com/tidwall/gjson
func main() {
	//jsonSerialize()
	jsonDeserialize()
}

type User struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	IsBlocked bool   `json:"isBlocked"`
	Books     []Book `json:"books"`
}

type Book struct {
	Name string `json:"name"`
	Year int    `json:"year"`
}

// jsonSerialize serialize data
func jsonSerialize() {
	var books []Book
	book := Book{
		"BN1",
		1990,
	}
	book2 := Book{
		"BN2",
		1991,
	}
	books = append(books, book, book2)

	someStruct := User{
		Name:      "John",
		Age:       40,
		IsBlocked: true,
		Books:     books,
	}
	mapStruct, _ := json.Marshal(someStruct)
	fmt.Println(string(mapStruct)) // {"name":"John","age":40,"isBlocked":true,"books":[{"name":"BN1","year":1990},{"name":"BN2","year":1991}]}

	someMap := map[string]string{"field1": "value1", "field2": "value2"}
	mapVar, _ := json.Marshal(someMap)
	fmt.Println(string(mapVar)) // {"field1":"value1","field2":"value2"}

	someMapWithAnyValue := map[string]interface{}{"field1": true, "field2": 1234, "arr": []string{"a", "b", "c"}}
	mapVar2, _ := json.Marshal(someMapWithAnyValue)
	fmt.Println(string(mapVar2)) // {"arr":["a","b","c"],"field1":true,"field2":1234}

	someSlice := []string{"A", "B", "C"}
	sliceVar, _ := json.Marshal(someSlice)
	fmt.Println(string(sliceVar)) // ["A","B","C"]

	boolVar, _ := json.Marshal(true)
	fmt.Println(string(boolVar)) // true
}

// jsonDeserialize deserialize data
func jsonDeserialize() {
	jsonString := []byte(`{"name":"John","age":40,"isBlocked":true,"books":[{"name":"BN1","year":1990},{"name":"BN2","year":1991}]}`)
	var data map[string]interface{}

	if err := json.Unmarshal(jsonString, &data); err != nil {
		panic(err)
	}
	fmt.Println(data)                                                              // map[age:40 books:[map[name:BN1 year:1990] map[name:BN2 year:1991]] isBlocked:true name:John]
	fmt.Println(data["name"])                                                      // John
	fmt.Println(data["books"].([]interface{})[0].(map[string]interface{})["name"]) // BN1

	var data2 User
	if err := json.Unmarshal(jsonString, &data2); err != nil {
		panic(err)
	}
	fmt.Println(data2.Books[0].Name) // BN1
}
