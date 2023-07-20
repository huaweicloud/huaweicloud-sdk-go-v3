package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuthInfoRequest Request Object
type ListAuthInfoRequest struct {

	// 认证信息名称
	AuthInfoName *string `json:"auth_info_name,omitempty"`

	// 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 默认为100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAuthInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuthInfoRequest struct{}"
	}

	return strings.Join([]string{"ListAuthInfoRequest", string(data)}, " ")
}
