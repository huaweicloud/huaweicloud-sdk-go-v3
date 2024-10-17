package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadAuditAgentRequest Request Object
type DownloadAuditAgentRequest struct {

	// 实例ID。可在查询实例列表接口的ID字段获取。
	InstanceId string `json:"instance_id"`

	// agent ID。可在查询数据库agent列表接口ID字段获取。
	AgentId string `json:"agent_id"`
}

func (o DownloadAuditAgentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadAuditAgentRequest struct{}"
	}

	return strings.Join([]string{"DownloadAuditAgentRequest", string(data)}, " ")
}
