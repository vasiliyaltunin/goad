// Copyright 2019 Vasiliy Altunin <skyr@altunin.info>. All rights reserved.
// Use of this source code is governed by a MIT license found in the LICENSE file.

/*
Package goad provides a
a tools to authennicate users in Active Directory (AD) and
read user info from AD.

Imortant!

Before use package create or update ".env" file in root of you progect folder!

In this file add folowing strings:

1. AD_SERVER_NAME - AD server name, keep in mind, you must use FQDN for
server name, or you can't connect to server with StartTLS.

Example:
	AD_SERVER_NAME=SRV1.ACME.CONTOSO.COM
2. AD_SERVER_PORT - AD server port is 389.

Example:
	AD_SERVER_PORT=389
3. AD_BASEDN - base path to place, where all you  users stored in AD.

Example:
	AD_BASEDN=OU=USERS,DC=ACME,DC=CONTOSO,DC=COM
4. AD_BIND_LOGIN - user login, that used to bind to LDAPS.

Example:
	AD_BIND_LOGIN=ldap-bind-acme
5. AD_BIND_PASS - user password, that used to bind to LDAPS.

Example:
	AD_BIND_PASS=xxxxxxxxxxxxxxxxxxxxxxxx

Package use init() to connect and bind to you AD server.

After connection established you can use functions like CheckAuth or GetAllUserAttrs
to get info from AD or auth you users!

*/
package goad