package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunTextTranslationResponse Response Object
type RunTextTranslationResponse struct {

	// 翻译原文，编码格式为UTF-8。调用失败时无此字段。
	SrcText *string `json:"src_text,omitempty"`

	// 翻译译文，编码格式为UTF-8。调用失败时无此字段。
	TranslatedText *string `json:"translated_text,omitempty"`

	// 调用失败时的错误码，具体请参见错误码。调用成功时无此字段。
	ErrorCode *string `json:"error_code,omitempty"`

	// 调用失败时的错误信息。调用成功时无此字段。
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunTextTranslationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunTextTranslationResponse struct{}"
	}

	return strings.Join([]string{"RunTextTranslationResponse", string(data)}, " ")
}
