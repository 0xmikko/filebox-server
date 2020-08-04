/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package main

import (
	"fmt"
	"github.com/MikaelLazarev/filebox-server/config"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

func loadHerokuEnv(appname string) {
	var cfg config.Config

	filename := ".env"
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Cant get working dir")
	}

	if strings.Contains(cwd, "/server/") {
		serverDir := "/server/"
		lastIndex := strings.Index(cwd, serverDir) + len(serverDir)
		filename = cwd[:lastIndex] + strings.TrimPrefix(filename, "./")
	}

	err = godotenv.Load(filename)
	if err != nil {
		log.Printf("Cant read .env cfg file %s\n%s ", filename, err)
	} else {
		log.Println("Getting configuration from " + filename)
	}

	rv := reflect.ValueOf(&cfg).Elem()
	num := rv.NumField()
	for i := 0; i < num; i++ {
		envValue := rv.Type().Field(i).Tag.Get("env")
		defaultValue := rv.Type().Field(i).Tag.Get("default")
		if envValue != "" && envValue != "PORT" {
			value := strings.Replace(config.GetEnv(envValue, defaultValue), "\\n", "\n", -1)
			commandLine := "heroku"
			params := []string{"config:set " + envValue + "=" + value + " -a " + appname}
			//commandLine := "ls -la"

			cmd := exec.Command(commandLine, params...)
			fmt.Println(cmd.String())
			//cmd.Dir = "."
			//var out bytes.Buffer
			//var stderr bytes.Buffer
			//cmd.Stdout = &out
			//cmd.Stderr = &stderr
			if err := cmd.Run(); err != nil {
				log.Println("Error: f", err)
			}

			cmd.Wait()

		}
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("use heroku_conf <appname>")
	}

	appname := os.Args[1]
	loadHerokuEnv(appname)

}
