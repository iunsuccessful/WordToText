package excel

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"regexp"
)

func ReadExcel(excelPath string) map[string]string {
	// 获取当前目录
	// 打开 Excel 文件
	file, err := xlsx.OpenFile(excelPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}

	// 获取第一个工作表
	sheet := file.Sheets[0]

	fileNameMap := make(map[string]string)
	re := regexp.MustCompile(`[《》\n]`) // 编译正则表达式，匹配 '《' 或 '》'
	// 遍历每一行
	for _, row := range sheet.Rows {
		// 第一行不处理
		// 遍历每列并打印单元格内容
		//for col, cell := range row.Cells {
		//	fmt.Print(col, cell.Value, "\t")
		//}
		if len(row.Cells) < 2 {
			continue
		}
		// 去掉《》

		oldName := re.ReplaceAllString(row.Cells[0].Value, "")
		newName := re.ReplaceAllString(row.Cells[1].Value, "")
		fileNameMap[oldName] = newName
	}
	return fileNameMap
}
