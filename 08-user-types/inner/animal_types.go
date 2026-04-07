package inner

import (
	"encoding/json"
	"fmt"
)

// type human struct {
// 	Name string
// 	age  int
// }

// func GetData() human {
// 	return human{
// 		Name: "InnerHuman",
// 		age:  45,
// 	}
// }

type Animal struct {
	Name string
	age  int
}

func (a Animal) Breed() string {
	return "Done"
}

type Creature struct {
	Something string
}

type Dog struct {
	Animal
	Creature
	TailLength int `json:"tail"`
	Height     int `json:"size" validate:"positive"`
}

func NewDog() *Dog {
	d := &Dog{}
	d.Animal.age = 4
	d.Name = "Rex"
	return d
}

func MyAnimalFunc() {
	d := Dog{}
	d.Animal.age = 4
	d.Name = "Rex"
	d.Something = "Something"
	d.Height = 50
	d.age = 3
	fmt.Printf("Dog: %+v\n", d)
	d.Breed()

	dogJson, err := json.Marshal(d)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Dog JSON:", string(dogJson))

}

type FriendsOfUser struct {
	BestFriendId int64
	OtherFriends []int64
}

type User struct { // Структура с именованными полями
	Id      int64
	Name    string
	Age     int
	friends FriendsOfUser
	Height  int
}

func (u User) String() string {
	return fmt.Sprintf("User{Id: %d, Name: %s, Age: %d, Height: %d}", u.Id, u.Name, u.Age, u.Height)
}

func (u *User) AddFriend(friendId int64, isBest bool) {
	if isBest {
		u.friends.BestFriendId = friendId
	} else {
		u.friends.OtherFriends = append(u.friends.OtherFriends, friendId)
	}
}

func MyUserFunc() {

	u1 := User{
		Id:     123456789,
		Name:   "Alice",
		Age:    12,
		Height: 170,
		friends: FriendsOfUser{
			BestFriendId: 987654321,
			OtherFriends: []int64{123123123},
		},
	}

	fmt.Println(u1.friends)

	u1.AddFriend(123, false)
	u1.AddFriend(321, true)

	fmt.Println(u1.friends)
}
