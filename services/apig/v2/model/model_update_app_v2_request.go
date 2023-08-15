package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAppV2Request Request Object
type UpdateAppV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 应用编号
	AppId string `json:"app_id"`

	Body *AppCreate `json:"body,omitempty"`
}

func (o UpdateAppV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppV2Request struct{}"
	}

	return strings.Join([]string{"UpdateAppV2Request", string(data)}, " ")
}
