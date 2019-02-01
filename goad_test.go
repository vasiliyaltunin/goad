// Copyright 2019 Vasiliy Altunin <skyr@altunin.info>. All rights reserved.
// Use of this source code is governed by a MIT license found in the LICENSE file.
package goad_test

import (
	"vasiliyaltunin/goad"
	"fmt"
	"log"
)

func ExampleCheckAuth() {

	username := "acme"
	password := "x1x2x3x4x5"

	result, err := goad.CheckAuth(username, password)
	if err != nil {
		log.Println("Error connecting :", err)
		return
	}

	if result {
		log.Println(username + " auth complete!")
	} else {
		log.Println(username + " access denied!")
	}

}

func ExampleGetAllUserAttrs() {

	userName := "acme"
	val := goad.GetAllUserAttrs(userName)

	fmt.Println(val.GetAttr("sn"))
	fmt.Println(val.GetGroups())
	fmt.Println(val.GetMail())

}

func ExampleUserAttibutes_GetAttr() {

	userName := "acme"
	val := goad.GetAllUserAttrs(userName)

	fmt.Println(val.GetAttr("sn"))
}

func ExampleUserAttibutes_IsInGroup() {
	userName := "acme"
	val := goad.GetAllUserAttrs(userName)
	access := val.IsInGroup("acme-admins")

	if access == true {
		fmt.Println("ACCESS GRANTED!")
	} else {
		fmt.Println("ACCESS DINIED!")
	}

}
