package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportPlaybookResponse Response Object
type ImportPlaybookResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 响应消息
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ImportPlaybookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportPlaybookResponse struct{}"
	}

	return strings.Join([]string{"ImportPlaybookResponse", string(data)}, " ")
}
