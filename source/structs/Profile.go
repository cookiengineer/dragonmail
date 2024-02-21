package structs

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

type Profile struct {
	Folder   string     `json:"folder"`
	Accounts []*Account `json:"accounts"`
}

func NewProfile(folder string) Profile {

	var profile Profile

	profile.Accounts = make([]*Account, 0)
	profile.SetFolder(folder)

	if profile.Folder != "" {

		err := os.MkdirAll(profile.Folder, 0776)

		if err != nil {
			fmt.Println(err.Error())
		}

	}

	profile.Read()

	return profile

}

func (profile *Profile) Read() {

	if profile.Folder != "" {

		entries, err1 := os.ReadDir(profile.Folder)

		if err1 == nil {

			for e := 0; e < len(entries); e++ {

				name := entries[e].Name()

				fmt.Println(name)

			}

		}

	}

}

func (profile *Profile) SetFolder(value string) {

	if strings.HasPrefix(value, "~/") {

		tmp1, err := user.Current()

		if err == nil {
			profile.Folder = "/home/" + tmp1.Username + "/" + value[2:]
		} else {

			tmp2 := os.Getenv("USER")

			if tmp2 != "" {
				profile.Folder = "/home/" + tmp2 + "/" + value[2:]
			}

		}

	} else if strings.HasPrefix(value, "/") {

		profile.Folder = value

	}

}
