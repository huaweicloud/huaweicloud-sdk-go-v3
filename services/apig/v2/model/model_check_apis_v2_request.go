package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckApisV2Request Request Object
type CheckApisV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	Body *ApiCheckInfoV2 `json:"body,omitempty"`
}

func (o CheckApisV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckApisV2Request struct{}"
	}

	return strings.Join([]string{"CheckApisV2Request", string(data)}, " ")
}
