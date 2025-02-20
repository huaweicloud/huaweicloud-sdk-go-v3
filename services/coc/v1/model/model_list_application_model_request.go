package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplicationModelRequest Request Object
type ListApplicationModelRequest struct {

	// 应用id
	ApplicationId *string `json:"application_id,omitempty"`

	// 组件id
	ComponentId *string `json:"component_id,omitempty"`

	// 最大的返回数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，上一页请求最后一个id
	Marker *string `json:"marker,omitempty"`

	// 分页页码
	PageNo *int32 `json:"page_no,omitempty"`
}

func (o ListApplicationModelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationModelRequest struct{}"
	}

	return strings.Join([]string{"ListApplicationModelRequest", string(data)}, " ")
}
