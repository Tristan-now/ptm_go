package deepl

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"net/url"
)

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
