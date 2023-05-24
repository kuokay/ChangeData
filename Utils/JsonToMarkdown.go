package Utils

import (
	"encoding/json"
	"fmt"
	"strings"

)



func Mainchange(jsonStr []byte) map[string]string {
	
	var data map[string]interface{}
	err := json.Unmarshal(jsonStr, &data)
	if err != nil {
		fmt.Println("Error while unmarshaling JSON:", err)
		return map[string]string{"msg": "err"}
	}

	markdownOutput := jsonToMarkdown(data)
	fmt.Println(markdownOutput)
	return map[string]string{"msg": markdownOutput}
}

func jsonToMarkdown(jsonData map[string]interface{}) string {
	var markdownBuilder strings.Builder

	var processNode func(node interface{}, indentLevel int)
	processNode = func(node interface{}, indentLevel int) {
		switch node := node.(type) {
		case map[string]interface{}:
			for key, value := range node {
				markdownBuilder.WriteString(fmt.Sprintf("%s- **%s**: ", strings.Repeat("  ", indentLevel), key))
				processNode(value, indentLevel+1)
			}
		case []interface{}:
			for _, item := range node {
				processNode(item, indentLevel)
			}
		default:
			markdownBuilder.WriteString(fmt.Sprintf("%s- %v\n", strings.Repeat("  ", indentLevel), node))
		}
	}

	processNode(jsonData, 0)

	return markdownBuilder.String()
}