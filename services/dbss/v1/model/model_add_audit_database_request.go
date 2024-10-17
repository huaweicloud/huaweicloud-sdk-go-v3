package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddAuditDatabaseRequest Request Object
type AddAuditDatabaseRequest struct {

	// 实例ID。可在查询实例列表接口的ID字段获取。
	InstanceId string `json:"instance_id"`

	Body *CreateDatabaseRequest `json:"body,omitempty"`
}

func (o AddAuditDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAuditDatabaseRequest struct{}"
	}

	return strings.Join([]string{"AddAuditDatabaseRequest", string(data)}, " ")
}
