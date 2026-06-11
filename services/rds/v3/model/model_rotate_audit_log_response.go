package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RotateAuditLogResponse Response Object
type RotateAuditLogResponse struct {

	// **参数解释**：  实例id。  **约束限制**：  不涉及。
	InstanceId     *string `json:"instance_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RotateAuditLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RotateAuditLogResponse struct{}"
	}

	return strings.Join([]string{"RotateAuditLogResponse", string(data)}, " ")
}
