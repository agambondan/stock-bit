package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

func WriteFile(write string) {
	//Logging
	file, err := os.OpenFile("info.log", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	if _, err := file.WriteString(fmt.Sprintf("%s - %s", time.Now().String()[:19], write)); err != nil {
		log.Fatal(err)
	}
}
