package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuditResultSystemAuditResultErrors struct {

	// 音频文件名。
	AudioName *string `json:"audio_name,omitempty"`

	// 文本文件名。
	TextName *string `json:"text_name,omitempty"`

	// 异常错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 异常错误信息。
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o AuditResultSystemAuditResultErrors) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuditResultSystemAuditResultErrors struct{}"
	}

	return strings.Join([]string{"AuditResultSystemAuditResultErrors", string(data)}, " ")
}
