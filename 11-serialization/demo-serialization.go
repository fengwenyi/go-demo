package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//testPerson()
	testUserVo()
}

func testPerson() {
	person := Person{
		Id:   "01",
		Name: "张三",
		Age:  18,
	}
	fmt.Println(person)
	b, _ := json.Marshal(person)
	fmt.Println(b)
	fmt.Println(string(b))
}

func testUserVo() {
	u := User{
		Id:       "01",
		Username: "zs",
		Password: "123456",
	}
	b, _ := u.toJson()
	fmt.Println(string(b))
	var u2 User
	err := u2.toObject(b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u2)
	}
}

type User struct {
	Id       string
	Username string
	Password string
}

type Person struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserVo struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func (u User) toJson() ([]byte, error) {
	/*vo := new(struct{
		Id string `json:"id"`
		Username string `json:"username"`
	})*/
	vo := new(UserVo)
	vo.Id = u.Id
	vo.Username = u.Username
	return json.Marshal(vo)
}

func (u *User) toObject(in []byte) error {
	vo := new(UserVo)
	err := json.Unmarshal(in, vo)
	if err == nil {
		u.Id = vo.Id
		u.Username = vo.Username
		return nil
	}
	return err
}
