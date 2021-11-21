package src

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type URLProperties map[string]string

func ReadURLs(filename string) (URLProperties, error) {
	urlProps := URLProperties{}

	if len(filename) == 0 {
		return urlProps, nil
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				urlProps[key] = value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return urlProps, nil
}

func GetOriginalURL(key string) (string, error) {
	urlProps, err := ReadURLs("url.properties")
	if err != nil {
		log.Fatal(err)
		return "", err
	} else {
		return urlProps[key], nil
	}
}
