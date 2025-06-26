package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectJobPermissionResponse Response Object
type ShowProjectJobPermissionResponse struct {
	Result *ShowJobPermissionResult `json:"result,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowProjectJobPermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectJobPermissionResponse struct{}"
	}

	return strings.Join([]string{"ShowProjectJobPermissionResponse", string(data)}, " ")
}
