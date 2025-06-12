package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDefaultProjectPermissionResponse Response Object
type ShowDefaultProjectPermissionResponse struct {

	// 返回结果
	Result *[]ShowDefaultProjectPermissionResponseBodyResult `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDefaultProjectPermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDefaultProjectPermissionResponse struct{}"
	}

	return strings.Join([]string{"ShowDefaultProjectPermissionResponse", string(data)}, " ")
}
