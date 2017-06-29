package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type users struct {
	Username string
	Password string
}

var userDatabase string

func main() {
	input, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(-1)
	}

	userInfo := bytes.Split(input, []byte("\n"))

	reply(checkCredentials(string(userInfo[0]), string(userInfo[1])))
}

func checkCredentials(username string, password string) bool {
	bytes := []byte(userDatabase)
	var u []users
	json.Unmarshal(bytes, &u)

	for i := range u {
		if strings.ToLower(username) == strings.ToLower(u[i].Username) {
			if password == u[i].Password {
				return true
			}
		}
	}

	return false
}

func reply(success bool) {
	if success {
		fmt.Print("OK\n")
		os.Exit(0)
	} else {
		fmt.Print("Wrong username or password\n")
		os.Exit(-1)
	}
}

func init() {
	userDatabase = `[{"Username": "user1", "Password": "Password1"},{"Username": "user2", "Password": "abc123"}]`
}
