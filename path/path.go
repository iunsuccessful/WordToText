package path

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ListFiles(dirPath string) ([]string, error) {
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

func RenameFile(oldTxtPath, newTxtPath string, fileNameMap map[string]string) {
	filenames, err := ListFiles(oldTxtPath)
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
				newPath := filepath.Join(newTxtPath, newName+".txt")
				err := os.Rename(filepath.Join(oldTxtPath, filename), newPath)
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

func EnsureDirectoryExists(filePath string) {
	dirPath, _ := filepath.Split(filePath)
	err := mkdirAll(dirPath, 0755)
	if err != nil {
		log.Fatal("Error creating directory:", err)
	}
}

// mkdirAll is a helper function to create directories recursively.
func mkdirAll(path string, mode os.FileMode) error {
	err := os.MkdirAll(path, mode)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("failed to create directory %q: %v", path, err)
	}
	return nil
}
