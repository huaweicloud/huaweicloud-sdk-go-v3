package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditAgentRequest Request Object
type ListAuditAgentRequest struct {

	// 实例ID。可在查询实例列表接口的ID字段获取。
	InstanceId string `json:"instance_id"`

	// 数据库ID,可在查询数据库列表接口ID字段获取。
	DbId string `json:"db_id"`

	// 偏移量
	Offset *string `json:"offset,omitempty"`

	// 查询记录数
	Limit *string `json:"limit,omitempty"`
}

func (o ListAuditAgentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditAgentRequest struct{}"
	}

	return strings.Join([]string{"ListAuditAgentRequest", string(data)}, " ")
}
