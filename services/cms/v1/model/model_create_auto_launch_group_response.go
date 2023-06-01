package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateAutoLaunchGroupResponse struct {

	// 创建好的智能购买组
	AutoLaunchGroupId *string `json:"auto_launch_group_id,omitempty"`

	// 错误码。 请求失败时，响应体中包含错误。 请求成功时，不在响应消息体中包含错误码。
	ErrorCode *int32 `json:"error_code,omitempty"`

	// 错误描述 请求失败时，响应体中包含错误 请求成功时，不在响应消息体中包含错误描述
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAutoLaunchGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAutoLaunchGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateAutoLaunchGroupResponse", string(data)}, " ")
}
