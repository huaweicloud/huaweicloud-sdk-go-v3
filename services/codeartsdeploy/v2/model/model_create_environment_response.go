package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEnvironmentResponse Response Object
type CreateEnvironmentResponse struct {

	// 请求成功失败状态
	Status *string `json:"status,omitempty"`

	// 环境id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEnvironmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnvironmentResponse struct{}"
	}

	return strings.Join([]string{"CreateEnvironmentResponse", string(data)}, " ")
}
