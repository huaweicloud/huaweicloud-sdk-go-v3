package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchDeleteAclV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 必须为delete
	Action string `json:"action"`

	Body *AclBatchDelete `json:"body,omitempty"`
}

func (o BatchDeleteAclV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAclV2Request struct{}"
	}

	return strings.Join([]string{"BatchDeleteAclV2Request", string(data)}, " ")
}
