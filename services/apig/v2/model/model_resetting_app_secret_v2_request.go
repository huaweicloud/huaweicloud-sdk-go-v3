package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ResettingAppSecretV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// APP的编号，可通过查询APP列表获取

	AppId string `json:"app_id"`

	Body *AppSecretReq `json:"body,omitempty"`
}

func (o ResettingAppSecretV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResettingAppSecretV2Request struct{}"
	}

	return strings.Join([]string{"ResettingAppSecretV2Request", string(data)}, " ")
}
