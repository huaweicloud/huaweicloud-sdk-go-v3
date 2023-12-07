package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobAuthInfosRequest Request Object
type ListJobAuthInfosRequest struct {

	// 认证信息名称
	AuthInfoName *string `json:"auth_info_name,omitempty"`

	// 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 默认为100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListJobAuthInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobAuthInfosRequest struct{}"
	}

	return strings.Join([]string{"ListJobAuthInfosRequest", string(data)}, " ")
}
