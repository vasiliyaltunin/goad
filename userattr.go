// Copyright 2019 Vasiliy Altunin <skyr@altunin.info>. All rights reserved.
// Use of this source code is governed by a MIT license found in the LICENSE file.

package goad

import "regexp"

//UserAttibutes - user attributes map
type UserAttibutes map[string][]string

//GetAttr - get user attr
func (attrs *UserAttibutes) GetAttr(name string) []string {
	return (*attrs)[name]
}

//GetGroups - get user groups
func (attrs *UserAttibutes) GetGroups() []string {
	return (*attrs)["memberOf"]
}

//GetDepartment - get user department
func (attrs *UserAttibutes) GetDepartment() string {
	return (*attrs)["department"][0]
}

//GetGivenName - get user given Name
func (attrs *UserAttibutes) GetGivenName() string {
	return (*attrs)["givenName"][0]
}

//GetDisplayName - get user Display Name
func (attrs *UserAttibutes) GetDisplayName() string {
	return (*attrs)["displayName"][0]
}

//GetStreetAddress - get user StreetAddress
func (attrs *UserAttibutes) GetStreetAddress() string {
	return (*attrs)["streetAddress"][0]
}

//GetCompany - get user company
func (attrs *UserAttibutes) GetCompany() string {
	return (*attrs)["company"][0]
}

//GetInitials - get user Initials
func (attrs *UserAttibutes) GetInitials() string {
	return (*attrs)["initials"][0]
}

//GetName - get user Name
func (attrs *UserAttibutes) GetName() string {
	return (*attrs)["name"][0]
}

//GetTitle - get user Title
func (attrs *UserAttibutes) GetTitle() string {
	return (*attrs)["title"][0]
}

//GetTelephoneNumber - get user TelephoneNumber
func (attrs *UserAttibutes) GetTelephoneNumber() string {
	return (*attrs)["telephoneNumber"][0]
}

//GetMail - get user Mail
func (attrs *UserAttibutes) GetMail() string {
	return (*attrs)["mail"][0]
}

//GetUserPrincipalName - get UserPrincipalName
func (attrs *UserAttibutes) GetUserPrincipalName() string {
	return (*attrs)["userPrincipalName"][0]
}

//IsInGroup - checks is user in given group
func (attrs *UserAttibutes) IsInGroup(groupName string) bool {

	for _, v := range (*attrs)["memberOf"] {
		re := regexp.MustCompile("CN=(.*?),")
		if re.FindStringSubmatch(v)[1] == groupName {
			return true
		}
	}
	return false
}
