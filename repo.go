package main

import "fmt"

/*
Publish is used by developers to register module
  - get url
    git config --get remote.origin.url
  - get the latest tag:
  - git ls-remote --tags https://github.com/georemo/cd-core.git
  - tag and push plugin project
  - go mod tidy
    git add <go-filenam>
    git commit -m "set version v0.0.1"
    git tag v0.0.1
    git push origin v0.0.1
  - get the latest tag to confirm update at github:
  - git ls-remote --tags https://github.com/georemo/cd-core.git
  - module name
  - Standards:
  - auto test
  - req test designed by corpdesk team
  - version
*/
func Publish(data string) string {
	fmt.Println("User::Auth()/input data:", data)
	resp := "{name:User, version:0.0.7 publisher: \"EMP Services Ltd\"}"
	return resp
}

/*
Info get details for specified module
  - module name
  - version
*/
func RepoInfo(data string) string {
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
func RepoGet(data string) string {
	fmt.Println("User::Auth()/input data:", data)
	resp := "{name:User, version:0.0.7 publisher: \"EMP Services Ltd\"}"
	return resp
}

/*
Update is used by developers or modulemanager to update modules data
  - module name
  - version
*/
func RepoUpdate(data string) string {
	fmt.Println("User::Auth()/input data:", data)
	resp := "{name:User, version:0.0.7 publisher: \"EMP Services Ltd\"}"
	return resp
}

/*
Expunge is used by module consumer to uninstall module
  - module name
  - version
*/
func RepoEnable(data string) string {
	fmt.Println("User::Auth()/input data:", data)
	resp := "{name:User, version:0.0.7 publisher: \"EMP Services Ltd\"}"
	return resp
}

/*
Expunge is used by module consumer to uninstall module
  - module name
  - version
*/
func RepoDisable(data string) string {
	fmt.Println("User::Auth()/input data:", data)
	resp := "{name:User, version:0.0.7 publisher: \"EMP Services Ltd\"}"
	return resp
}

/*
Expunge is used by module consumer to uninstall module
  - module name
  - version
*/
func Delete(data string) string {
	fmt.Println("User::Auth()/input data:", data)
	resp := "{name:User, version:0.0.7 publisher: \"EMP Services Ltd\"}"
	return resp
}
