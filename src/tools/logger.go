package tools

import (
	log "github.com/sirupsen/logrus"
)

func Log() {
	log.SetFormatter(&log.JSONFormatter{})

	standardFields := log.Fields{
		"hostname": "staging-1",
		"appname":  "foo-app",
		"session":  "1ce3f6v",
	}

	customFields := log.Fields{
		"string": "foo",
		"int": 1,
		"float": 1.1,
	}

	log.WithFields(standardFields).WithFields(customFields).Info("My first ssl event from Golang")

}