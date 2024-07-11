package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateApplicationPermissionsResponse Response Object
type BatchUpdateApplicationPermissionsResponse struct {

	// 请求成功失败状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchUpdateApplicationPermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateApplicationPermissionsResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateApplicationPermissionsResponse", string(data)}, " ")
}
