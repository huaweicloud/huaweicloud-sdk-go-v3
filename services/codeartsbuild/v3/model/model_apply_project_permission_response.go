package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyProjectPermissionResponse Response Object
type ApplyProjectPermissionResponse struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 错误
	Error *interface{} `json:"error,omitempty"`

	// 结果
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ApplyProjectPermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyProjectPermissionResponse struct{}"
	}

	return strings.Join([]string{"ApplyProjectPermissionResponse", string(data)}, " ")
}
