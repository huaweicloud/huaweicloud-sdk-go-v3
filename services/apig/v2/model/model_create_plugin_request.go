package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePluginRequest Request Object
type CreatePluginRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	Body *PluginCreate `json:"body,omitempty"`
}

func (o CreatePluginRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePluginRequest struct{}"
	}

	return strings.Join([]string{"CreatePluginRequest", string(data)}, " ")
}
