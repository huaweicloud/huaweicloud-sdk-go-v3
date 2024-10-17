package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddAuditAgentRequest Request Object
type AddAuditAgentRequest struct {

	// 实例ID。可在查询实例列表接口的ID字段获取。
	InstanceId string `json:"instance_id"`

	Body *AuditAgentRequest `json:"body,omitempty"`
}

func (o AddAuditAgentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAuditAgentRequest struct{}"
	}

	return strings.Join([]string{"AddAuditAgentRequest", string(data)}, " ")
}
