package docx

import (
	"code.sajari.com/docconv/v2"
	"log"
	"os"
)

func GetDocxContent(path string) string {
	res, err := docconv.ConvertPath(path)
	if err != nil {
		log.Println(path, " ", err)
		return ""
	}
	return res.Body
}

func ConvertToText(docPath, txtPath string) {
	content := GetDocxContent(docPath)
	if len(content) <= 0 {
		return
	}
	// 写入 .txt
	err := os.WriteFile(txtPath, []byte(content), 0644)
	if err != nil {
		log.Fatal(docPath, err)
	}
}
