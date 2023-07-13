package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppCodeV2Request Request Object
type CreateAppCodeV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 应用编号
	AppId string `json:"app_id"`

	Body *AppCodeCreate `json:"body,omitempty"`
}

func (o CreateAppCodeV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppCodeV2Request struct{}"
	}

	return strings.Join([]string{"CreateAppCodeV2Request", string(data)}, " ")
}
