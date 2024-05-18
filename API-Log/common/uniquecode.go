package common

import (
	"crypto/rand"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

func UniqueCode() (encryptString string, requestCode string, currentYear string, err error) {

	log.Println("UniqueCode(+)")
	KeyValue := ""

	session := uuid.NewV4()
	sessionSHA256 := session.String()
	KeyValue = strings.ReplaceAll(sessionSHA256, "-", "")
	//to get random number
	b := make([]byte, 8)
	rand.Read(b)
	requestcode := fmt.Sprintf("%x", b)

	// to get current Year
	t := time.Now()
	year := t.Year()
	CurrentYear := strconv.Itoa(year)
	log.Println("UniqueCode(-)")

	return KeyValue, requestcode, CurrentYear, nil
}

func CallReference() string {

	log.Println("CallReference(+)")
	KeyValue := ""

	session := uuid.NewV4()
	sessionSHA256 := session.String()
	KeyValue = strings.ReplaceAll(sessionSHA256, "-", "")
	log.Println("CallReference(-)")

	return KeyValue
}
