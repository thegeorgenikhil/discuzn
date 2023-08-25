package config

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

func covertToString(name ConfigName) string {
	return string(name)
}

func mustHave(key ConfigName) {
	if !viper.IsSet(covertToString(key)) {
		panic(fmt.Sprintf("key %s is not set", key))
	}
}

func mustGetString(key ConfigName) string {
	mustHave(key)
	return viper.GetString(covertToString(key))
}

func mustGetInt(key ConfigName) int {
	mustHave(key)
	v, err := strconv.Atoi(viper.GetString(covertToString(key)))
	if err != nil {
		panic(fmt.Sprintf("key %s is not a valid Integer value", key))
	}

	return v
}
