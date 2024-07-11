package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnvironmentRequest struct {

	// 环境名称
	Name *string `json:"name,omitempty"`

	// 环境描述信息
	Description *string `json:"description,omitempty"`
}

func (o EnvironmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvironmentRequest struct{}"
	}

	return strings.Join([]string{"EnvironmentRequest", string(data)}, " ")
}
