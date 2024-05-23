package config

import (
	"errors"
	"github.com/subosito/gotenv"
	"os"
)

func LoadDefaultEnv() {
	LoadEnv("", "", false)
	LoadEnv("", os.Getenv("APP_ENV"), true)
}

func LoadEnv(basePath string, env string, override bool) {
	if env != "" {
		env = "." + env
	}
	load(
		[]string{
			basePath + ".env" + env,
		}, override,
	)
}

func load(filenames []string, override bool) {
	for _, filename := range filenames {
		f, err := os.Open(filename)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return
			}
			panic(err)
		}
		if override {
			err = gotenv.OverApply(f)
		} else {
			err = gotenv.Apply(f)
		}
		f.Close()
		if err != nil {
			panic(err)
		}
	}
}
