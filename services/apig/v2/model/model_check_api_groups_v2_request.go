package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckApiGroupsV2Request Request Object
type CheckApiGroupsV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	Body *ApiGroupCheck `json:"body,omitempty"`
}

func (o CheckApiGroupsV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckApiGroupsV2Request struct{}"
	}

	return strings.Join([]string{"CheckApiGroupsV2Request", string(data)}, " ")
}
