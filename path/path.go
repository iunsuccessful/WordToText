package path

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func listFiles(dirPath string) ([]string, error) {
	var filenames []string

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			filenames = append(filenames, info.Name())
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error walking the path: %w", err)
	}
	return filenames, nil
}

func ReadFiles() {
	dirPath := "D:\\Users\\Jactitator\\Downloads\\20240521 叶丹丹文档处理\\20240521 叶丹丹文档处理\\txt" // 替换为实际目录路径
	filenames, err := listFiles(dirPath)
	if err != nil {
		fmt.Println("Error listing files:", err)
		return
	}
	for _, filename := range filenames {
		fmt.Println(filename)
	}
}

func RenameFile(fileNameMap map[string]string) {
	dirPath := "D:\\Users\\Jactitator\\Downloads\\20240521 叶丹丹文档处理\\20240521 叶丹丹文档处理\\txt"
	newDirPath := "D:\\Users\\Jactitator\\Downloads\\20240521 叶丹丹文档处理\\20240521 叶丹丹文档处理\\txtnew"
	filenames, err := listFiles(dirPath)
	if err != nil {
		fmt.Println("Error listing files:", err)
		return
	}

	unprocessedMap := make(map[string]string)

	// fileNameMap 复制给 unprocessedMap
	for key, value := range fileNameMap {
		unprocessedMap[key] = value
	}

	for _, filename := range filenames {

		// 遍历 map, 找当前文件名包含的 key
		for key, newName := range fileNameMap {
			if strings.Contains(filename, key) {
				newPath := filepath.Join(newDirPath, newName+".txt")
				err := os.Rename(filepath.Join(dirPath, filename), newPath)
				if err != nil {
					fmt.Println("Error renaming file:", err)
				}
				fmt.Printf("Renamed %s to %s\n", filename, newName)
				delete(unprocessedMap, key)
			}
		}
		//if newName, ok := fileNameMap[filename]; ok {
		//	newPath := filepath.Join(dirPath, newName)
		//	err := os.Rename(filepath.Join(dirPath, filename), newPath)
		//	if err != nil {
		//		fmt.Println("Error renaming file:", err)
		//	}
		//}
		//fmt.Println(filename)
	}
	// 打印 unprocessedMap
	for key, value := range unprocessedMap {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}
