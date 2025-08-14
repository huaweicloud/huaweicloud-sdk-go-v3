package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplicationProvidersRequest Request Object
type ListApplicationProvidersRequest struct {

	// 每个请求返回的最大结果数。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记
	Marker *string `json:"marker,omitempty"`

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`
}

func (o ListApplicationProvidersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationProvidersRequest struct{}"
	}

	return strings.Join([]string{"ListApplicationProvidersRequest", string(data)}, " ")
}
