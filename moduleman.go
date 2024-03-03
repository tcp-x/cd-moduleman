package main

import "fmt"

/*
Info get details for specified module
  - module name
  - version
*/
func Info(data string) string {
	fmt.Println("User::Auth()/input data:", data)
	resp := data
	return resp
}

/*
Get is used to get module data
  - list of modules
  - details of specific module
  - arbritrary query of module database
*/
func Get(data string) string {
	fmt.Println("User::Auth()/input data:", data)
	resp := "{name:User, version:0.0.7 publisher: \"EMP Services Ltd\"}"
	return resp
}

/*
Update is used by developers or modulemanager to update modules data
  - module name
  - version
*/
func Update(data string) string {
	fmt.Println("User::Auth()/input data:", data)
	resp := "{name:User, version:0.0.7 publisher: \"EMP Services Ltd\"}"
	return resp
}

/*
Update is used by developers or modulemanager to update modules data
  - get the latest tag:
  - git ls-remote --tags https://github.com/tcp-x/cd-core.git
  - version
*/
func Upgrade(data string) string {
	fmt.Println("User::Auth()/input data:", data)
	resp := "{name:User, version:0.0.7 publisher: \"EMP Services Ltd\"}"
	return resp
}

/*
Install is used by module consumers to install module to corpdesk system

  - clone repository with specifi tag

  - https://stackoverflow.com/questions/20280726/how-to-clone-a-specific-git-tag

  - create temp directory ~/tmp/cd-user

  - git clone --depth 1 --branch v0.0.1 https://github.com/tcp-x/cd-user.git ~/tmp/cd-user

  - mv ./tmp/user.go ~/cd-cli/plugins/cd-user/

  - test plugin

  - do cd-cli req "{test request}"

  - on successfull test, update local cd installed modules data, eg date of installation, version etc

  - module name

  - version
*/
func Install(data string) string {
	fmt.Println("User::Auth()/input data:", data)
	resp := "{name:User, version:0.0.7 publisher: \"EMP Services Ltd\"}"
	return resp
}

/*
Uninstall is used by module consumer to uninstall module
  - module name
  - version
*/
func Uninstall(data string) string {
	fmt.Println("User::Auth()/input data:", data)
	resp := "{name:User, version:0.0.7 publisher: \"EMP Services Ltd\"}"
	return resp
}

/*
Uninstall is used by module consumer to uninstall module
  - module name
  - version
*/
func Block(data string) string {
	fmt.Println("User::Auth()/input data:", data)
	resp := "{name:User, version:0.0.7 publisher: \"EMP Services Ltd\"}"
	return resp
}

/*
Expunge is used by module consumer to uninstall module
  - module name
  - version
*/
func ModEnable(data string) string {
	fmt.Println("User::Auth()/input data:", data)
	resp := "{name:User, version:0.0.7 publisher: \"EMP Services Ltd\"}"
	return resp
}

/*
Expunge is used by module consumer to uninstall module
  - module name
  - version
*/
func ModDisable(data string) string {
	fmt.Println("User::Auth()/input data:", data)
	resp := "{name:User, version:0.0.7 publisher: \"EMP Services Ltd\"}"
	return resp
}
