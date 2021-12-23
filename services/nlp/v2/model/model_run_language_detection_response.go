package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunLanguageDetectionResponse struct {
	// 调用成功时表示调用结果，语种缩写对应名称如下： zh    中文 en    英文 ru    俄语 ja    日文 de    德文 fr    法文 es    西班牙文 pt    葡萄牙文 it    意大利文 tr    土耳其文 ar    阿拉伯文 ko    韩语 th    泰国语 ms    马来语 vi    越南语 当输入文本过短或不明确时，识别结果可能不准确； 当输入文本包含多种语言时，会返回占比最高的语种。 调用失败时无此字段。

	DetectedLanguage *string `json:"detected_language,omitempty"`
	// 调用失败时的错误码，具体请参见错误码。调用成功时无此字段。

	ErrorCode *string `json:"error_code,omitempty"`
	// 调用失败时的错误信息。调用成功时无此字段。

	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunLanguageDetectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunLanguageDetectionResponse struct{}"
	}

	return strings.Join([]string{"RunLanguageDetectionResponse", string(data)}, " ")
}
