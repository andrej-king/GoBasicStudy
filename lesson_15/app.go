package main

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"strings"
)

// work with json
// libs for work with json:
// https://github.com/tidwall/gjson
// https://github.com/tidwall/sjson
func main() {
	//jsonSerialize()
	//jsonDeserialize()
	//gjsonLibPractice()
	sjsonLibPractice()
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

// https://github.com/tidwall/gjson
// go get -u github.com/tidwall/gjson
func gjsonLibPractice() {
	jsonString := `{"name":"John","age":40,"isBlocked":true, "name2": {"first": "John", "last": "Doe"}, "books":[{"name":"BN1","year":1990},{"name":"BN2","year":1991}]}`
	value := gjson.Get(jsonString, "name2.first")
	fmt.Println("gjson value: ", value.String()) // John

	jsonString2 := `{
					  "name": {"first":  "John", "last":  "Doe"},
					  "age": 37,
					  "children": ["Sara", "Alex", "Jack"],
					  "fav.movie": "Dear Hunter",
					  "friends": [
						{"first": "Dale", "last":  "Murphy", "age": 44, "nets":  ["ig", "fb", "tw"]},
						{"first": "Roger", "last":  "Craig", "age": 68, "nets":  ["fb", "tw"]},
						{"first": "Jane", "last":  "Murphy", "age": 47, "nets":  ["ig", "tw"]}
					  ]
					}`
	valueArray := gjson.Get(jsonString2, "children")
	fmt.Println("gjson arrayValue:", valueArray) // ["Sara", "Alex", "Jack"]

	valueWithDot := gjson.Get(jsonString2, "fav\\.movie")
	fmt.Println("gjson valueWithDot:", valueWithDot) // Dear Hunter

	valueFromArray := gjson.Get(jsonString2, "children.1")
	fmt.Println("gjson valueFromArray:", valueFromArray) // Alex

	arrayLength := gjson.Get(jsonString2, "children.#")
	fmt.Println("gjson arrayLength:", arrayLength) // 3

	valueFromArray2 := gjson.Get(jsonString2, "child*.2")
	fmt.Println("gjson valueFromArray2:", valueFromArray2) // Jack

	valueFromArray3 := gjson.Get(jsonString2, "friends.1.age")
	fmt.Println("gjson valueFromArray3:", valueFromArray3) // 68

	valueFromArray4 := gjson.Get(jsonString2, `friends.#(last=="Murphy").first`)
	fmt.Println("gjson valueFromArray4:", valueFromArray4) // Dale

	valueFromArray5 := gjson.Get(jsonString2, `friends.#(last=="Murphy")#.first`)
	fmt.Println("gjson valueFromArray5:", valueFromArray5) // ["Dale","Jane"]

	valueFromArray6 := gjson.Get(jsonString2, `friends.#(age>=40)#.first`)
	fmt.Println("gjson valueFromArray6:", valueFromArray6) // ["Dale","Roger","Jane"]

	// custom modifier
	gjson.AddModifier("case", func(json, arg string) string {
		if arg == "upper" {
			return strings.ToUpper(json)
		} else if arg == "lower" {
			return strings.ToLower(json)
		}
		return json
	})

	valueFromArrayWithUpperCase := gjson.Get(jsonString2, `friends.#(age>=40)#.first|@case:upper`)
	fmt.Println("gjson valueFromArrayWithUpperCase:", valueFromArrayWithUpperCase) // ["DALE","ROGER","JANE"]

	fmt.Println("gjson parse:", gjson.Parse(jsonString2).Get("name")) // {"first":  "John", "last":  "Doe"}

	//invalidJsonString := `{"name": abc}`
	if !gjson.Valid(jsonString2) {
		panic("JSON IS NOT VALID")
	}

	valueFromArray7 := gjson.Get(jsonString2, `friends.#(age>=40)#.first|@case:upper`)
	for _, firstName := range valueFromArray7.Array() {
		fmt.Println("firstName:", firstName) // DALE ROGER JANE
	}

	// parsing json to map
	result, ok := gjson.Parse(jsonString2).Value().(map[string]interface{})
	if !ok {
		panic("Error parsing to map")
	}
	fmt.Println(result) // json converted to map
}

// https://github.com/tidwall/sjson
// go get -u github.com/tidwall/sjson
func sjsonLibPractice() {
	jsonString := `{
					  "name": {"first":  "John", "last":  "Doe"},
					  "age": 37,
					  "children": ["Sara", "Alex", "Jack"],
					  "fav.movie": "Dear Hunter",
					  "friends": [
						{"first": "Dale", "last":  "Murphy", "age": 44, "nets":  ["ig", "fb", "tw"]},
						{"first": "Roger", "last":  "Craig", "age": 68, "nets":  ["fb", "tw"]},
						{"first": "Jane", "last":  "Murphy", "age": 47, "nets":  ["ig", "tw"]}
					  ]
					}`

	updateJsonValue, _ := sjson.Set(jsonString, "name.first", "Jack")
	fmt.Println("sjson updateJsonValue:", updateJsonValue) //  "name": {"first":  "Jack", "last":  "Doe"}...

	addJsonValue, _ := sjson.Set(jsonString, "children.4", "Jack")
	fmt.Println("sjson addJsonValue:", addJsonValue) //  ..."children": ["Sara","Alex","Jack",null,"Jack"]...

	correctAddJsonValue, _ := sjson.Set(jsonString, "children.-1", "John")
	fmt.Println("sjson addJsonValue:", correctAddJsonValue) //  ..."children": ["Sara", "Alex", "Jack","John"]...

	addJsonElement, _ := sjson.Set(jsonString, "newObj", map[string]interface{}{"hello": "world"})
	fmt.Println("sjson addJsonElement:", addJsonElement) //  ..."newObj":{"hello":"world"}}...

	emptyString := ""
	addJsonElement2, _ := sjson.Set(emptyString, "newObj", map[string]interface{}{"hello": "world"})
	fmt.Println("sjson addJsonElement2:", addJsonElement2) //  "newObj":{"hello":"world"}}

	removeElement, _ := sjson.Delete(jsonString, "friends")
	fmt.Println("sjson removeElement:", removeElement) // json with deleted element
}
