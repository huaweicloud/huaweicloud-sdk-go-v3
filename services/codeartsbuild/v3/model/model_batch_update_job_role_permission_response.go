package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateJobRolePermissionResponse Response Object
type BatchUpdateJobRolePermissionResponse struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 错误
	Error *interface{} `json:"error,omitempty"`

	// 结果
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchUpdateJobRolePermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateJobRolePermissionResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateJobRolePermissionResponse", string(data)}, " ")
}
