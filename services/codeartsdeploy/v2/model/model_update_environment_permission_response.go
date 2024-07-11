package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEnvironmentPermissionResponse Response Object
type UpdateEnvironmentPermissionResponse struct {

	// 请求成功失败状态
	Status *string `json:"status,omitempty"`

	Result         *DevUcEnvironmentPermission `json:"result,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o UpdateEnvironmentPermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEnvironmentPermissionResponse struct{}"
	}

	return strings.Join([]string{"UpdateEnvironmentPermissionResponse", string(data)}, " ")
}
