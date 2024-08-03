package main

import (
	"chapter-c29/model"
	"fmt"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
)

func main()  {
    var user1 = &model.User{
        Id: "u001",
        Name: "Sylvana Winnrunner",
        Password: "f0r th3 h0rd3",
        Gender: model.UserGender_FEMALE,
    }

    // var userList = &model.UserList{
    //     List: []*model.User{
    //         user1,
    //     },
    // }

    // var garage1 = &model.Garage{
    //     Id: "g001",
    //     Name: "Kalimdor",
    //     Coordinate: &model.GarageCoordinate{
    //         Latitude: 23.22131212,
    //         Longitude: 53.22101233,
    //     },
    // }

    // var garageList = &model.GarageList{
    //     List: []*model.Garage{
    //         garage1,
    //     },
    // }

    // var garageListByUser = &model.GarageListByUser{
    //     List: map[string]*model.GarageList{
    //         user1.Id:garageList,
    //     },
    // }

    fmt.Println(user1)
    fmt.Println(user1.String())
    // fmt.Println(garageListByUser)

    jsonb, err1 := protojson.Marshal(user1)
    if err1 != nil {
        fmt.Println(err1.Error())
        os.Exit(0)
     }
    fmt.Printf("# ==== As JSON String\n       %s \n", string(jsonb))

    protoObject := new(model.User)
    err2 := protojson.Unmarshal(jsonb, protoObject)
    if err2 != nil {
        fmt.Println(err2.Error())
        os.Exit(0)
    }
    
  fmt.Printf("# ==== As String\n       %s \n", protoObject.String())
}