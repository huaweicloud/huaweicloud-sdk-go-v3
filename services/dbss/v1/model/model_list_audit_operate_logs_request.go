package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditOperateLogsRequest Request Object
type ListAuditOperateLogsRequest struct {

	// 实例ID
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
