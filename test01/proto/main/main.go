package main

import (
	"fmt"

	Student "Rpc.Study.go/test01/proto"
	"google.golang.org/protobuf/proto"
)

type Hello struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Courses []string `json:"courses"`
}

func main() {
	req := Student.Student{
		Name:    "bobby",
		Age:     18,
		Courses: []string{"go", "gin", "微服务"},
	}
	// jsonStruct := Hello{
	// 	Name:    "bobby",
	// 	Age:     18,
	// 	Courses: []string{"go", "gin", "微服务"},
	// }
	// Json, _ := json.Marshal(jsonStruct)
	// fmt.Println(len(Json))
	rsp, _ := proto.Marshal(&req)
	NewReq := Student.Student{}

	_ = proto.Unmarshal(rsp, &NewReq)
	fmt.Println(NewReq.Name, NewReq.Age, NewReq.Courses)
	// fmt.Println(len(rsp))

}
