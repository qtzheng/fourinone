package utils

import (
	"github.com/qtzheng/fourinone/utils"
	"os"
	"path/filepath"
)

func CheckFile(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
func CreateFile(path string, data string) bool {
	f := os.NewFile(os.O_CREATE|os.O_WRONLY, name)
	defer f.Close()
	_, err := f.WriteString(data)
	if utils.CheckError(err) {
		utils.LogInfo(utils.Exit, err.Error())
	}

}
func SelfPath() string {
	path, _ := filepath.Abs(os.Args[0])
	return path
}
func SelfDir() string {
	return filepath.Dir(SelfPath())
}
