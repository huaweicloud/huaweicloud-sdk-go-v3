package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFunctionApplicationsRequest Request Object
type ListFunctionApplicationsRequest struct {

	// 本次查询最大返回的数据条数，最大值500，默认值100
	Limit *string `json:"limit,omitempty"`

	// 本次查询起始位置，默认值0
	Marker *string `json:"marker,omitempty"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o ListFunctionApplicationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionApplicationsRequest struct{}"
	}

	return strings.Join([]string{"ListFunctionApplicationsRequest", string(data)}, " ")
}
