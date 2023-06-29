package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppCodeAutoV2Request Request Object
type CreateAppCodeAutoV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 应用编号
	AppId string `json:"app_id"`
}

func (o CreateAppCodeAutoV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppCodeAutoV2Request struct{}"
	}

	return strings.Join([]string{"CreateAppCodeAutoV2Request", string(data)}, " ")
}
