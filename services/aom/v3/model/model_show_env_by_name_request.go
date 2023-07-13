package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEnvByNameRequest Request Object
type ShowEnvByNameRequest struct {

	// 环境名称
	EnvironmentName string `json:"environment_name"`

	// 环境region
	Region string `json:"region"`

	// 组件id
	ComponentId string `json:"component_id"`
}

func (o ShowEnvByNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEnvByNameRequest struct{}"
	}

	return strings.Join([]string{"ShowEnvByNameRequest", string(data)}, " ")
}
