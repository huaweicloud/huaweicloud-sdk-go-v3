package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CancelingAuthorizationV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 授权关系的ID

	AppAuthId string `json:"app_auth_id"`
}

func (o CancelingAuthorizationV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelingAuthorizationV2Request struct{}"
	}

	return strings.Join([]string{"CancelingAuthorizationV2Request", string(data)}, " ")
}
