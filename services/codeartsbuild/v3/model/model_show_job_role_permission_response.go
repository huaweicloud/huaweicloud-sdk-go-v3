package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobRolePermissionResponse Response Object
type ShowJobRolePermissionResponse struct {

	// 结果
	Result *[]ShowJobRolePermissionResult `json:"result,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowJobRolePermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobRolePermissionResponse struct{}"
	}

	return strings.Join([]string{"ShowJobRolePermissionResponse", string(data)}, " ")
}
