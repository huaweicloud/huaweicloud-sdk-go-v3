package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccountsRequest Request Object
type ListAccountsRequest struct {

	// 创建令牌接口调用签发的访问令牌
	AccessToken string `json:"access-token"`

	// 页面中最大结果数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记。非分页的接口，不使用该值。
	Marker *string `json:"marker,omitempty"`
}

func (o ListAccountsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccountsRequest struct{}"
	}

	return strings.Join([]string{"ListAccountsRequest", string(data)}, " ")
}
