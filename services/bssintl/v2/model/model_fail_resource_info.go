package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FailResourceInfo struct {

	// 状态码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述信息。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 资源ID。
	ResourceId *string `json:"resource_id,omitempty"`
}

func (o FailResourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailResourceInfo struct{}"
	}

	return strings.Join([]string{"FailResourceInfo", string(data)}, " ")
}
