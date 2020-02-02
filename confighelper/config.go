package confighelper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Configuration struct {
	Port              int    `json:"Port"`
	Static_Variable   string `json:"Static_Variable"`
	Connection_String string `json:"Connection_String"`
}

func GetConfiguration() Configuration {

	configuration := Configuration{}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	content, err := ioutil.ReadFile(dir + "/confighelper/config.json")
	if err != nil {
		log.Fatal("===>", err)
	}
	fmt.Println(content)
	err1 := json.Unmarshal(content, &configuration)
	if err1 != nil {

		log.Fatal("errrr==============>", err1)
	}
	fmt.Println("configuration struct", configuration)
	return configuration
}
