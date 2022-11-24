package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowResJobResponse struct {

	// 是否成功
	IsSuccess *bool `json:"is_success,omitempty"`

	// 作业详情
	Jobs *[]Jobs `json:"jobs,omitempty"`

	// 返回消息（请求成功时，不返回此字段）
	Message *string `json:"message,omitempty"`

	// 错误码（请求成功时，不返回此字段）
	ErrorCode      *string `json:"error_code,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowResJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResJobResponse struct{}"
	}

	return strings.Join([]string{"ShowResJobResponse", string(data)}, " ")
}
