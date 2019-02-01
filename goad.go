// Copyright 2019 Vasiliy Altunin <skyr@altunin.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goad

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	auth "gopkg.in/korylprince/go-ad-auth.v2"
)

var config *auth.Config
var conn *auth.Conn

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	adSName := os.Getenv("AD_SERVER_NAME")

	adSPort, err := strconv.Atoi(os.Getenv("AD_SERVER_PORT"))
	if err != nil {
		log.Println("AD server port error! Check .env file: ", err)
		//handle err
		return
	}

	adSDN := os.Getenv("AD_BASEDN")

	config = &auth.Config{

		Server:   adSName,
		Port:     adSPort,
		BaseDN:   adSDN,
		Security: auth.SecurityStartTLS,
	}

	adBindName := os.Getenv("AD_BIND_LOGIN")
	adBindPass := os.Getenv("AD_BIND_PASS")

	connect(adBindName, adBindPass)
}

//GetConfig - returns config
func GetConfig() *auth.Config {
	return config
}

//GetConn - returns connection
func GetConn() *auth.Conn {
	return conn
}

//CheckAuth - check user auth in AD
func CheckAuth(username, password string) (bool, error) {
	status, err := auth.Authenticate(config, username, password)

	if err != nil {
		//handle err
		log.Println("Auth error:", err)
		return false, err
	}

	if !status {
		log.Println("Not authed, bad password or username!")
		//handle failed authentication
		return false, err
	}

	return true, nil

}

//Connect - connects and bind to AD server
func connect(bindUser, bindPass string) {
	var err error
	conn, err = config.Connect()
	if err != nil {
		log.Println("Error connecting AD server:", err)
		return
	}

	var status bool
	status, err = conn.Bind(bindUser, bindPass)
	if err != nil {
		log.Println("Error binding to AD server:", err)
	}

	if !status {
		panic("Not binded to AD server, bad bind password or username!")
	}

}

//GetAllUserAttrs - get all user attribute
func GetAllUserAttrs(userName string) UserAttibutes {

	val, err := conn.Search("(&(objectClass=person)(samaccountname="+userName+"))", nil, 1)
	if err != nil {
		log.Println("User search error:", err)
	}

	if len(val) == 0 {
		log.Println("User not found on AD server: " + userName)

		return UserAttibutes{}
	}

	var userAttr UserAttibutes
	userAttr = make(map[string][]string)

	for _, v := range val[0].Attributes {
		if (v.Name != "userCertificate") && (v.Name != "userParameters") {
			userAttr[v.Name] = v.Values
			// fmt.Printf("=%#v\n", v)
		}
	}

	return userAttr

}
