package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSendDiffSmsResponse Response Object
type BatchSendDiffSmsResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误描述
	Description *string `json:"description,omitempty"`

	// 短信状态
	Result         *[]SmsId `json:"result,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o BatchSendDiffSmsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSendDiffSmsResponse struct{}"
	}

	return strings.Join([]string{"BatchSendDiffSmsResponse", string(data)}, " ")
}
