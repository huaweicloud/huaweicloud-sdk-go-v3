package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAuditAgentRequest Request Object
type DeleteAuditAgentRequest struct {

	// 实例ID。可在查询实例列表接口的ID字段获取。
	InstanceId string `json:"instance_id"`

	// agent ID。可在查询数据库agent列表接口ID字段获取。
	AgentId string `json:"agent_id"`

	// 数据库ID, 可在查询数据库列表接口ID字段获取。
	DbId string `json:"db_id"`
}

func (o DeleteAuditAgentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAuditAgentRequest struct{}"
	}

	return strings.Join([]string{"DeleteAuditAgentRequest", string(data)}, " ")
}
