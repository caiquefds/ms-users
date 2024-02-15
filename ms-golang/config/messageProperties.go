package config

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadMessagesPropertiesFile() (map[string]string, error) {
	messages := make(map[string]string)

	if len("./message.properties") == 0 {
		return messages, nil
	}
	file, err := os.Open("./message.properties")
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
				messages[key] = value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return messages, nil
}
