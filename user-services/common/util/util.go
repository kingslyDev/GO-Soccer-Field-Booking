package util

import (
	"os"
	"reflect"
	"strconv"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func BindFromJson(dest any, file, path string) error {
	v := viper.New()

	v.SetConfigType("json")
	v.AddConfigPath(path)
	v.SetConfigName(file)

	err := v.ReadInConfig()
	if err != nil {
		return err
	}

	err = v.Unmarshal(dest)
	if err != nil {
		logrus.Errorf("Failed to Unmarshal: %v", err)
		return err
	}
	return nil
}

func setEnvFromConsulKV(v *viper.Viper) error {
	env := make(map[string]any)

	err := v.Unmarshal(&env)
	if err != nil {
		logrus.Errorf("Failed to Unmarshal: %v", err)
		return err
	}

	for k, v := range env {
		var (
			valOf = reflect.ValueOf(v)
			val   string
		)

		switch valOf.Kind() {
		case reflect.String:
			val = valOf.String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val = strconv.FormatInt(valOf.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val = strconv.FormatUint(valOf.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			val = strconv.FormatFloat(valOf.Float(), 'f', -1, 64)
		case reflect.Bool:
			val = strconv.FormatBool(valOf.Bool())
		default:
			panic("Invalid type")
		}

		err := os.Setenv(k, val)
		if err != nil {
			logrus.Errorf("Failed to set env: %v", err)
			return err
		}
		
	}

	return nil
}

func BindFromConsulKV(dest any,endPoint, path string) error {
	v := viper.New()

	v.SetConfigType("json")
	v.AddRemoteProvider("consul", endPoint, path)
	v.OnConfigChange(func(e fsnotify.Event) {
		if err := setEnvFromConsulKV(v); err != nil {
			logrus.Errorf("Failed to set env from consul: %v", err)
		}
	})

	if err := v.ReadRemoteConfig(); err != nil {
		logrus.Errorf("Failed to read remote config: %v", err)
		return err
	}

	if err := setEnvFromConsulKV(v); err != nil {
		logrus.Errorf("Failed to set env from consul: %v", err)
		return err
	}

	err := v.Unmarshal(dest)
	if err != nil {
		logrus.Errorf("Failed to Unmarshal: %v", err)
		return err
	}

	return nil
}
