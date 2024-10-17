package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAuditInstanceRequest Request Object
type UpdateAuditInstanceRequest struct {

	// 实例ID。可在查询实例列表接口的ID字段获取。
	InstanceId string `json:"instance_id"`

	Body *UpdateAuditBean `json:"body,omitempty"`
}

func (o UpdateAuditInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuditInstanceRequest struct{}"
	}

	return strings.Join([]string{"UpdateAuditInstanceRequest", string(data)}, " ")
}
