package main

import (
	"fmt"
)

func main() {

	type Msg struct {
		date     string
		isHot    bool
		isOnsale bool
		item     []string
	}

	type Goods struct {
		pname string
		price int
		msg   Msg
	}

	msg := Msg{
		date:     "2020-09-29",
		isHot:    false,
		isOnsale: false,
	}

	phone := Goods{
		pname: "xxx",
		price: 2000,
		msg:   msg,
	}

	fmt.Print(phone)
}
