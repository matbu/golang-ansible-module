package main

import (
	"encoding/json"
	"fmt"
	"os"

	"golang-ansible-module/plugins/module_utils"
)

var args struct {
        Name     string `json:"name"`
}

func main() {
	var response moduleutils.Response

	if len(os.Args) != 2 {
		response.Msg = "No argument file provided"
		moduleutils.FailJson(response)
	}

	argsFile := os.Args[1]
	argsData, err := os.ReadFile(argsFile)
	if err != nil {
		response.Msg = fmt.Sprintf("Failed to read argument file: %v", err)
		moduleutils.FailJson(response)
	}

	if err := json.Unmarshal(argsData, &args); err != nil {
		response.Msg = fmt.Sprintf("Failed to parse argument file: %v", err)
		moduleutils.FailJson(response)
	}

	response.Msg = fmt.Sprintf("Hello %s", args.Name)
	moduleutils.ExitJson(response)

}
