package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobAuthInfosResponse Response Object
type ListJobAuthInfosResponse struct {

	// 是否成功
	IsSuccess *bool `json:"is_success,omitempty"`

	// 请求消息
	Message *string `json:"message,omitempty"`

	// 认证信息个数
	Count *int32 `json:"count,omitempty"`

	// 认证信息列表
	AuthInfos      *[]AuthInfo `json:"auth_infos,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListJobAuthInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobAuthInfosResponse struct{}"
	}

	return strings.Join([]string{"ListJobAuthInfosResponse", string(data)}, " ")
}
