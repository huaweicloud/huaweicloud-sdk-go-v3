package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchAuditDatabaseRequest Request Object
type SwitchAuditDatabaseRequest struct {

	// 实例ID。可在查询实例列表接口的ID字段获取。
	InstanceId string `json:"instance_id"`

	Body *SwitchAuditDbRequest `json:"body,omitempty"`
}

func (o SwitchAuditDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchAuditDatabaseRequest struct{}"
	}

	return strings.Join([]string{"SwitchAuditDatabaseRequest", string(data)}, " ")
}
