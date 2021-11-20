package env

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)


func GetStringWithDefault(name string, defaultValue string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		logrus.WithFields(logrus.Fields{
			name: defaultValue,
		}).Warn("Did not find environment variable, using default value.")
		return defaultValue
	}
	logrus.WithFields(logrus.Fields{
		name:         value,
		"value_type": fmt.Sprintf("%T", value),
	}).Info("Found environment variable")
	return value
}


func GetString(name string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		logrus.WithField("env_name", name).Fatal("Cannot proceed did not find string parameter.")
	}
	logrus.WithFields(logrus.Fields{
		name:         value,
		"value_type": fmt.Sprintf("%T", value),
	}).Info("Found environment variable")
	return value
}

func GetIntWithDefault(name string, defaultValue int) int {
	value, ok := os.LookupEnv(name)
	if !ok {
		logrus.WithFields(logrus.Fields{
			name: defaultValue,
		}).Warn("Did not find environment variable, using default value.")
		return defaultValue
	}
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		logrus.WithError(err).WithField(name, value).Fatal("Cannot proceed, failed to parse int")
	}
	logrus.WithFields(logrus.Fields{
		name:         value,
		"value_type": fmt.Sprintf("%T", value),
	}).Info("Found environment variable")
	return valueInt
}

