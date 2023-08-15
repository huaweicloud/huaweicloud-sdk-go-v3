package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteApiAclBindingV2Request Request Object
type BatchDeleteApiAclBindingV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 必须为delete
	Action string `json:"action"`

	Body *AclBindingBatchDelete `json:"body,omitempty"`
}

func (o BatchDeleteApiAclBindingV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteApiAclBindingV2Request struct{}"
	}

	return strings.Join([]string{"BatchDeleteApiAclBindingV2Request", string(data)}, " ")
}
