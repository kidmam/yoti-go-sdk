package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/getyoti/yoti-go-sdk"

	_ "github.com/joho/godotenv/autoload"
)

var (
	sdkID  string
	key    []byte
	client *yoti.Client
)

func main() {
	var sdkID = os.Getenv("YOTI_CLIENT_SDK_ID")
	var key, err = ioutil.ReadFile(os.Getenv("YOTI_KEY_FILE_PATH"))

	if err != nil {
		log.Printf("Unable to retrieve `YOTI_KEY_FILE_PATH`. Error: `%s`", err)
		return
	}

	var client = yoti.Client{
		SdkID: sdkID,
		Key:   key}

	givenNames := "Edward Richard George"
	familyName := "Heath"

	amlAddress := yoti.AmlAddress{
		Country: "GBR"}

	amlProfile := yoti.AmlProfile{
		GivenNames: givenNames,
		FamilyName: familyName,
		Address:    amlAddress}

	result, err := client.PerformAmlCheck(amlProfile)

	if err != nil {
		log.Printf(
			"Unable to retrieve AML result. Error: %s", err)
	} else {
		log.Printf(
			"AML Result for %s %s:",
			givenNames,
			familyName)
		log.Printf(
			"On PEP list: %s",
			strconv.FormatBool(result.OnPEPList))
		log.Printf(
			"On Fraud list: %s",
			strconv.FormatBool(result.OnFraudList))
		log.Printf(
			"On Watch list: %s",
			strconv.FormatBool(result.OnWatchList))
	}
}