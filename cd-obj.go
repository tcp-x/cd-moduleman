package main

import "fmt"

type cd_obj string

/*
Info get details for specified module
  - module name
  - version
*/
func (controller cd_obj) GetCdObj(data string) string {
	fmt.Println("CdObj::GetCdObj():", data)
	resp := data
	return resp
}

var CdObj cd_obj
