package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFlinkJarResponse Response Object
type CreateFlinkJarResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 消息内容
	Message *string `json:"message,omitempty"`

	Job            *CreateJobRespJob `json:"job,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o CreateFlinkJarResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlinkJarResponse struct{}"
	}

	return strings.Join([]string{"CreateFlinkJarResponse", string(data)}, " ")
}
