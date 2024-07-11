package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEnvironmentResponse Response Object
type UpdateEnvironmentResponse struct {

	// 请求成功失败状态
	Status *string `json:"status,omitempty"`

	// 环境id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateEnvironmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEnvironmentResponse struct{}"
	}

	return strings.Join([]string{"UpdateEnvironmentResponse", string(data)}, " ")
}
