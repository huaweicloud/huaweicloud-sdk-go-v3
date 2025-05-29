package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectPermissionResponse Response Object
type ShowProjectPermissionResponse struct {
	Result *ShowUserProjectPermissionResult `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowProjectPermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectPermissionResponse struct{}"
	}

	return strings.Join([]string{"ShowProjectPermissionResponse", string(data)}, " ")
}
