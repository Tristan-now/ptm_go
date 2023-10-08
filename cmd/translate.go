/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"net/url"
	"os/exec"
	"regexp"
)

// translateCmd represents the translate command

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

// 3.翻译
func getURL(ispro bool) string {
	if ispro {
		return "https://api.deepl.com/v2/translate"
	}
	return "https://api-free.deepl.com/v2/translate"
}

var errorCode = map[int]string{
	400: "Bad request. Please check error message and your parameters.",
	403: "Authorization failed. Please supply a valid auth_key parameter.",
	404: "The requested resource could not be found.",
	413: "The request size exceeds the limit.",
	414: "The request URL is too long. You can avoid this error by using a POST request instead of a GET request, and sending the parameters in the HTTP body.",
	429: "Too many requests. Please wait and resend your request.",
	456: "Quota exceeded. The character limit has been reached.",
	503: "Resource currently unavailable. Try again later.",
	529: "Too many requests. Please wait and resend your request.",
}

func Translate(text string, target_language string, source_language string, ispro bool, token string) (string, error) {
	URL := getURL(ispro)
	//创建了一个 url.Values 类型的变量 key，用于构建HTTP POST请求的参数。这包括身份验证密钥、要翻译的文本、目标语言和源语言。
	key := url.Values{}
	key.Add("auth_key", token)
	key.Add("text", text)
	key.Add("target_lang", target_language)
	key.Add("source_lang", source_language)
	//使用http.PostForm()函数发送HTTP POST请求，并将响应结构（*Response）返回给resp变量。
	resp, err := http.PostForm(URL, key)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	//如果响应状态码不是200，则返回错误消息。
	msg, ok := errorCode[resp.StatusCode]
	if ok {
		return fmt.Sprintf("%d: %s", resp.StatusCode, msg), nil
	}

	//body, err := ioutil.ReadAll(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {

		return "", err
	}

	str_json := string(body)
	t := gjson.Get(str_json, "translations.0.text").String()
	return t, nil
}

var t = &cobra.Command{
	Use:   "t",
	Short: "",
	Long:  `translate papers`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("translate called")
		fmt.Println("select the text you want to translate ")
		DEEPL_TOKEN := "ff5294c4-2874-22da-c873-5419f08c2b08:fx"
		for {
			output, err := Get_select_text()
			if output == "" {
				continue
			}
			re_output := Regular_output(output)
			if err != nil {
				fmt.Println(err)
			}
			re, err := Translate(re_output, "zh", "en", false, DEEPL_TOKEN)
			if err != nil {
				fmt.Println("Please install xsel")
			}
			fmt.Println(output)
			fmt.Println()
			fmt.Println(re)
			fmt.Println()
			fmt.Println("---------------------------------------------")
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(t)
}
