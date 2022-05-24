package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"net/http"
	"os"
)

func configure(context *cli.Context) error {
	fmt.Println("Generating configuration file")

	id, err := find("N0013")
	if err != nil {
		return err
	}
	id, err = find(id)
	if err != nil {
		return err
	}
	id, err = find(id)
	if err != nil {
		return err
	}
	fmt.Println("Please enter your username, multiple usernames can be separated by a comma")
	var user string
	_, err = fmt.Scan(&user)
	if err != nil {
		return err
	}
	if user == "" {
		return fmt.Errorf("no username entered")
	}

	config := map[string]string{
		"id":   id,
		"user": user,
	}
	configFile, err := json.Marshal(config)
	if err != nil {
		return err
	}
	// write to file
	path := context.String("config")
	if path == "" {
		path = "config.json"
	}
	if path == "ENV" {
		fmt.Println("GLFY_ID=" + id)
		fmt.Println("GLFY_USER=" + user)
		fmt.Println("Environment variables shown. Please set them in your shell at runtime")
	} else {
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				panic(err)
			}
		}(f)
		_, err = io.Copy(f, bytes.NewReader(configFile))
		if err != nil {
			return err
		}
		fmt.Println("Configuration file generated")
	}
	return nil
}

func find(id string) (string, error) {
	resp, err := http.Get(host + "/pub/vol/config/organization?pid=" + id)
	if err != nil {
		return "", err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var orgData map[string]interface{}
	err = json.Unmarshal(data, &orgData)
	if err != nil {
		return "", err
	}
	result := orgData["result"].([]interface{})
	if len(result) == 0 {
		return "", fmt.Errorf("no organization found")
	}

	for i, v := range result {
		fmt.Printf("[%d] %s\n", i, v.(map[string]interface{})["title"])
	}

	choice := -1
	fmt.Printf("Please select organization: ")
	_, err = fmt.Scan(&choice)
	if err != nil {
		return "", err
	}
	if choice < 0 || choice >= len(result) {
		return "", fmt.Errorf("invalid choice")
	}

	return result[choice].(map[string]interface{})["id"].(string), nil
}
