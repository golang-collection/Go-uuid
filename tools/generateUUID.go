package tools

import (
	"github.com/satori/go.uuid"
	"os/user"
	"strconv"
)

/**
* @Author: super
* @Date: 2020-09-09 20:33
* @Description:
**/

func GenerateUUIDV1() string {
	// NewV1 returns UUID based on current timestamp and MAC address.
	id := uuid.NewV1()
	return id.String()
}

func GenerateUUIDV2() (string, error) {
	// NewV2 returns DCE Security UUID based on POSIX UID/GID.
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	b, err := strconv.Atoi(u.Uid)
	if err != nil{
		return "", err
	}
	id := uuid.NewV2(byte(b))
	//return id.String(), nil
	return id.String(), nil
}

func GenerateUUIDV3() string {
	// NewV3 returns UUID based on MD5 hash of namespace UUID and name.
	i := uuid.UUID{}
	id := uuid.NewV3(i, "uuidv3")
	return id.String()
}

func GenerateUUIDV4() string {
	// NewV4 returns random generated UUID.
	id := uuid.NewV4()
	return id.String()
}

func GenerateUUIDV5() string {
	// NewV5 returns UUID based on SHA-1 hash of namespace UUID and name.
	i := uuid.UUID{}
	id := uuid.NewV5(i, "uuidv5")
	return id.String()
}
