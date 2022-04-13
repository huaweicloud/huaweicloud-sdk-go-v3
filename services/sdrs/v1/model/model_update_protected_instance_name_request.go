package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateProtectedInstanceNameRequest struct {
	// 保护实例的ID。

	ProtectedInstanceId string `json:"protected_instance_id"`

	Body *UpdateProtectedInstanceNameRequestBody `json:"body,omitempty"`
}

func (o UpdateProtectedInstanceNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProtectedInstanceNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateProtectedInstanceNameRequest", string(data)}, " ")
}
