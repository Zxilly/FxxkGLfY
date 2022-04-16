package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"net/http"
	"os"
	"strings"
)

func makeRecord(context *cli.Context) error {
	var user string
	var id string

	path := context.String("config")
	if path == "" {
		path = "config.json"
	}

	if path == "ENV" {
		id = os.Getenv("GLFY_ID")
		user = os.Getenv("GLFY_USER")
		if id == "" {
			return fmt.Errorf("id not found in env")
		}
		if user == "" {
			return fmt.Errorf("user not found in env")
		}
	} else {
		cfgByte, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		var cfg map[string]string
		err = json.Unmarshal(cfgByte, &cfg)
		if err != nil {
			return err
		}

		var found bool
		id, found = cfg["id"]
		if !found {
			return fmt.Errorf("id not found in config")
		}
		user, found = cfg["user"]
		if !found {
			return fmt.Errorf("user not found in config")
		}
	}

	var users []string
	if strings.Contains(user, ",") {
		users = strings.Split(user, ",")
	} else {
		users = append(users, user)
	}

	courseResp, err := http.Get("http://osscache.vol.jxmfkj.com/html/assets/js/course_data.js")
	if err != nil {
		return err
	}

	courseRespBody, err := io.ReadAll(courseResp.Body)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(courseResp.Body)

	courseRespBody = courseRespBody[len("var course_data = "):]

	var courseData map[string]interface{}
	err = json.Unmarshal(courseRespBody, &courseData)
	if err != nil {
		return err
	}
	result := courseData["result"].(map[string]interface{})
	courseID := result["id"].(string)
	title := result["title"].(string)

	for _, cardNo := range users {
		payload := &map[string]string{
			"course": courseID,
			"nid":    id,
			"cardNo": cardNo,
		}
		payloadByte, err := json.Marshal(payload)
		if err != nil {
			return err
		}
		targetURL := "http://osscache.vol.jxmfkj.com/pub/vol/volClass/join?accessToken="
		signResp, err := http.Post(targetURL, "application/json", bytes.NewBuffer(payloadByte))
		if err != nil {
			return err
		}
		signRespBody, err := io.ReadAll(signResp.Body)
		if err != nil {
			return err
		}
		var signResult map[string]interface{}
		err = json.Unmarshal(signRespBody, &signResult)
		if err != nil {
			return err
		}
		if signResult["status"].(float64) != 200 {
			return fmt.Errorf("sign failed\n %s", signRespBody)
		} else {
			fmt.Printf("%s sign %s success\n", cardNo, title)
		}
	}

	fmt.Printf("all task on %s finished\n", title)

	return nil
}
