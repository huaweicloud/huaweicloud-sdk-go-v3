package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadFileResponse Response Object
type UploadFileResponse struct {

	// 参数解释： 对话ID。 约束限制： 不涉及 取值范围： 不涉及 默认取值： 不涉及
	ChatId         *string `json:"chat_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadFileResponse struct{}"
	}

	return strings.Join([]string{"UploadFileResponse", string(data)}, " ")
}
