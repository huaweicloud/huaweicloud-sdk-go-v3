package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditOperateLogsRequest Request Object
type ListAuditOperateLogsRequest struct {

	// 实例ID。可在查询实例列表接口的ID字段获取。
	InstanceId string `json:"instance_id"`

	Body *OperateLogGetRequest `json:"body,omitempty"`
}

func (o ListAuditOperateLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditOperateLogsRequest struct{}"
	}

	return strings.Join([]string{"ListAuditOperateLogsRequest", string(data)}, " ")
}
