package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RotateAuditLogRequestBody 轮转审计日志请求体
type RotateAuditLogRequestBody struct {

	// **参数解释**：  实例id。  **约束限制**：  不涉及。
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o RotateAuditLogRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RotateAuditLogRequestBody struct{}"
	}

	return strings.Join([]string{"RotateAuditLogRequestBody", string(data)}, " ")
}
