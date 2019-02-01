[![GoDoc](https://godoc.org/github.com/vasiliyaltunin/goad?status.svg)](https://godoc.org/github.com/vasiliyaltunin/goad)

# About

`goad` is simple wrapper around `go-ad-auth` the great [go-ad-auth](https://github.com/korylprince/go-ad-auth) library to help with authentication and search in  Active Directory.

# Installing

`go get github.com/vasiliyaltunin/goad`

**Dependencies:**

* [github.com/joho/godotenv](https://github.com/joho/godotenv)
* [gopkg.in/korylprince/go-ad-auth.v2](https://godoc.org/gopkg.in/korylprince/go-ad-auth.v2)


If you have any issues or questions [create an issue](https://github.com/vasiliyaltunin/goad/issues/).

# Setup 

Package use init() to connect and bind to you AD server, so before use package create or update ".env" file in root of you progect folder!

In this file add folowing strings:

| Name                    | Description |
| ---------------| ------------- |
| AD_SERVER_NAME | AD server name, keep in mind, you must use FQDN for server name, or you can't connect to server with StartTLS. |
| AD_SERVER_PORT | AD server port - defaults to 389 |
| AD_BASEDN      | base path to place, where all you users stored in AD. |
| AD_BIND_LOGIN  | user login, that used to bind to LDAPS. |
| AD_BIND_PASS   | user password, that used to bind to LDAPS. |


Example:

```
AD_SERVER_NAME=SRV1.ACME.CONTOSO.COM
AD_SERVER_PORT=389
AD_BASEDN=OU=USERS,DC=ACME,DC=CONTOSO,DC=COM
AD_BIND_LOGIN=ldap-bind-acme
AD_BIND_PASS=xxxxxxxxxxxxxxxxxxxxxxxx
```

# Usage

`godoc github.com/vasiliyaltunin/goad`

You can easy check auth with single function

Example:

```go
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
```
You can read user attributes from AD

Example:

```go
userName := "acme"
val := goad.GetAllUserAttrs(userName)

fmt.Println(val.GetAttr("sn"))
fmt.Println(val.GetGroups())
fmt.Println(val.GetMail())
```

You can check is user in some group

Example:

```go
userName := "acme"
val := goad.GetAllUserAttrs(userName)
access := val.IsInGroup("acme-admins")

if access == true {
    fmt.Println("ACCESS GRANTED!")
} else {
    fmt.Println("ACCESS DINIED!")
}

```

See more examples on [GoDoc](https://godoc.org/github.com/vasiliyaltunin/goad).

