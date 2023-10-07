package utils

import (
	"errors"
	"os/exec"
	"regexp"
)

// 1.将选中文字放入output
func Get_select_text() (string, error) {
	//打开新的终端，输入xsel，查看是否安装
	output, err := exec.Command("xsel").Output()
	if err != nil {
		//fmt.Println("Please install xsel")
		errors.New("Please install xsel")
		return "", err
	}

	return string(output), nil
}

// 2.对output进行格式化，去掉标点符号和换行等
func Regular_output(output string) string {

	regex := regexp.MustCompile("[^a-zA-Z0-9 ]+")
	re_output := regex.ReplaceAllString(output, "")
	return re_output
}
