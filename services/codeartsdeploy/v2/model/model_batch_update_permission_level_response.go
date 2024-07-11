package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdatePermissionLevelResponse Response Object
type BatchUpdatePermissionLevelResponse struct {

	// 请求成功失败状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchUpdatePermissionLevelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdatePermissionLevelResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdatePermissionLevelResponse", string(data)}, " ")
}
