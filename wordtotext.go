package main

import (
	"github.com/iunsuccessful/WordToText/docx"
	"github.com/iunsuccessful/WordToText/excel"
	"github.com/iunsuccessful/WordToText/path"
	"log"
	"path/filepath"
	"strings"
)

func main() {

	excelPath := "/Users/iunsuccessful/Downloads/20240521文档处理/UC分发文汇总表格.xlsx"

	docPath := "/Users/iunsuccessful/Downloads/20240521文档处理/docx/"

	txtPath := "/Users/iunsuccessful/Downloads/20240521文档处理/txt1/"
	newTxtPath := "/Users/iunsuccessful/Downloads/20240521文档处理/txt2/"

	filenames, err := path.ListFiles(docPath)

	if err != nil {
		log.Fatal(err)
	}

	// 确保目录存在
	path.EnsureDirectoryExists(txtPath)

	// .docx -> txt
	for _, filename := range filenames {
		newFileName := strings.Replace(filename, ".docx", ".txt", -1)
		docx.ConvertToText(filepath.Join(docPath, filename), filepath.Join(txtPath, newFileName))
		log.Printf("%s -> %s\n", filename, newFileName)
	}

	// rename
	// 读取 excel
	fileNameMap := excel.ReadExcel(excelPath)

	// 确保目录存在
	path.EnsureDirectoryExists(newTxtPath)

	// 根据 excel 重命名
	path.RenameFile(txtPath, newTxtPath, fileNameMap)

}
