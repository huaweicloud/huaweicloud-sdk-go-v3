package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ResizeProtectedInstanceRequest struct {
	// 保护实例的ID。

	ProtectedInstanceId string `json:"protected_instance_id"`

	Body *ResizeProtectedInstanceRequestBody `json:"body,omitempty"`
}

func (o ResizeProtectedInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeProtectedInstanceRequest struct{}"
	}

	return strings.Join([]string{"ResizeProtectedInstanceRequest", string(data)}, " ")
}
