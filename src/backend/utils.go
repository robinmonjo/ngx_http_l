package main

import (
	"os"
	"os/user"
	"strconv"
)

func chown(username string, file string) error {
	user, err := user.Lookup(username)
	if err != nil {
		return err
	}
	uid, err := strconv.Atoi(user.Uid)
	if err != nil {
		return err
	}
	gid, err := strconv.Atoi(user.Gid)
	if err != nil {
		return err
	}
	return os.Chown(file, uid, gid)
}
