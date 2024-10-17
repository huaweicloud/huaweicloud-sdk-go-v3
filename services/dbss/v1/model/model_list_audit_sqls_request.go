package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditSqlsRequest Request Object
type ListAuditSqlsRequest struct {

	// 实例ID。可在查询实例列表接口的ID字段获取。
	InstanceId string `json:"instance_id"`

	Body *AuditSqlRequest `json:"body,omitempty"`
}

func (o ListAuditSqlsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditSqlsRequest struct{}"
	}

	return strings.Join([]string{"ListAuditSqlsRequest", string(data)}, " ")
}
