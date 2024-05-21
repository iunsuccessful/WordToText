package main

import (
	"github.com/iunsuccessful/WordToText/excel"
	"github.com/iunsuccessful/WordToText/path"
)

func main() {

	// 读取 excel
	fileNameMap := excel.ReadExcel()
	//fmt.Print(fileNameMap)
	// 读取文件夹下面的文件列表
	//path.ReadFiles()
	// 修改名称
	path.RenameFile(fileNameMap)

}
