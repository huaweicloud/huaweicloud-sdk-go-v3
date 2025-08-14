package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectResourcesPolicyStatesSummaryRequest Request Object
type CollectResourcesPolicyStatesSummaryRequest struct {

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 最大的返回数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty"`
}

func (o CollectResourcesPolicyStatesSummaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectResourcesPolicyStatesSummaryRequest struct{}"
	}

	return strings.Join([]string{"CollectResourcesPolicyStatesSummaryRequest", string(data)}, " ")
}
