package main

import (
	"io/ioutil"
	"os"
	"os/exec"
        "fmt"

	"github.com/Sirupsen/logrus"
)

type (
	Plugin struct {
		Url       string
		Username  string
		Password  string
		ScanImage string
		CaCert    string
		Threshold string
		Security  string
	}
)

func (p Plugin) Exec() error {
	os.Setenv("CLAIR_ADDR", p.Url)
	_, exist := os.LookupEnv("DOCKER_USER")
	if !exist {
		os.Setenv("DOCKER_USER", p.Username)
	}
	_, exist = os.LookupEnv("DOCKER_PASSWORD")
	if !exist {
		os.Setenv("DOCKER_PASSWORD", p.Password)
	}

	if p.Threshold != "" {
		os.Setenv("CLAIR_THRESHOLD", p.Threshold)
                fmt.Printf("set threshold to: %v", p.Threshold) }
	_, exist = os.LookupEnv("PLUGIN_THRESHOLD")
	if exist {
		os.Setenv("CLAIR_THRESHOLD", p.Threshold)
                fmt.Printf("set threshold to: %v", p.Threshold)
	}

	if p.Security != "" {
		os.Setenv("CLAIR_OUTPUT", p.Security)
                fmt.Printf("set security to: %v", p.Security) }
	_, exist = os.LookupEnv("PLUGIN_SECURITY")
	if exist {
		os.Setenv("CLAIR_OUTPUT", p.Security)
                fmt.Printf("set security to: %v", p.Security)
	}

	var commands []*exec.Cmd

	if p.CaCert != "" {
		commands = append(commands, installCaCert(p.CaCert))
	}

	commands = append(commands, scanImage(p.ScanImage))

	for _, command := range commands {
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		err := command.Run()

		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err,
			}).Fatal("Failed to execute a command")
		}
	}

	return nil
}

func installCaCert(cacert string) *exec.Cmd {
	ioutil.WriteFile("/usr/local/share/ca-certificates/ca_cert.crt", []byte(cacert), 0644)
	return exec.Command(
		"update-ca-certificates",
	)
}

func scanImage(image string) *exec.Cmd {
	return exec.Command(
		"klar",
		image,
	)
}
