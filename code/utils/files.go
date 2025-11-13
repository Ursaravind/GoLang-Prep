package utils
import (
	"fmt"
	"path/filepath"
)
func FilePath(){
	// join
	path := filepath.Join("dir1","dir2","text.html")
	fmt.Println(path)
	// Dir
	dir_path := filepath.Dir(path)
	fmt.Println(dir_path)
	// extracting the file with filepath.Base
	fmt.Println("File name : "+filepath.Base(path))

}