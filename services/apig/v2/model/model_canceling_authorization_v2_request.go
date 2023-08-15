package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelingAuthorizationV2Request Request Object
type CancelingAuthorizationV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 授权关系的编号
	AppAuthId string `json:"app_auth_id"`
}

func (o CancelingAuthorizationV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelingAuthorizationV2Request struct{}"
	}

	return strings.Join([]string{"CancelingAuthorizationV2Request", string(data)}, " ")
}
