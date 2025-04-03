package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccountAgenciesRequest Request Object
type ListAccountAgenciesRequest struct {

	// 创建令牌接口调用签发的访问令牌
	AccessToken string `json:"access-token"`

	// 页面中最大结果数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记。非分页的接口，不使用该值。
	Marker *string `json:"marker,omitempty"`

	// 帐户的全局唯一标识符（ID）
	AccountId string `json:"account_id"`
}

func (o ListAccountAgenciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccountAgenciesRequest struct{}"
	}

	return strings.Join([]string{"ListAccountAgenciesRequest", string(data)}, " ")
}
