package config

import (
	"fmt"
	"os"
	"strconv"
)

func envStr(name string, def ...string) (*string, error) {
	env := os.Getenv(name)

	if env == "" {
		for _, v := range def {
			return &v, nil
		}

		return nil, fmt.Errorf("env %s is empty", name)
	}

	return &env, nil
}

func envBool(name string, def ...bool) (*bool, error) {
	env := os.Getenv(name)

	if env == "" {
		for _, v := range def {
			return &v, nil
		}

		return nil, fmt.Errorf("env %s is empty", name)
	}

	value, err := strconv.ParseBool(env)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

func envInt(name string, def ...int) (*int, error) {
	env := os.Getenv(name)

	if env == "" {
		for _, v := range def {
			return &v, nil
		}

		return nil, fmt.Errorf("env %s is empty", name)
	}

	value, err := strconv.Atoi(env)
	if err != nil {
		return nil, err
	}

	return &value, nil
}
