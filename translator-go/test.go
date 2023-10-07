package main

import (
	"fmt"
	"translator-go/deepl"
	"translator-go/utils"
)

func main() {
	DEEPL_TOKEN := "ff5294c4-2874-22da-c873-5419f08c2b08:fx"

	output, err := utils.Get_select_text()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("原始文本为：", output)
	fmt.Println("-----------")
	re_output := utils.Regular_output(output)

	fmt.Println("正则化后文本为:", re_output)
	fmt.Println("-----------")
	res, err := deepl.Translate(re_output, "zh", "en", false, DEEPL_TOKEN)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("翻译结果为：", res)

}
