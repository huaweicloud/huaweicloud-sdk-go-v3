package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AttachPluginToApiRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// API编号
	ApiId string `json:"api_id"`

	Body *ApiOperPluginInfo `json:"body,omitempty"`
}

func (o AttachPluginToApiRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachPluginToApiRequest struct{}"
	}

	return strings.Join([]string{"AttachPluginToApiRequest", string(data)}, " ")
}
