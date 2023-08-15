package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGatewayResponsesV2Response Response Object
type ListGatewayResponsesV2Response struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 响应列表
	Responses      *[]ResponsesInfo `json:"responses,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListGatewayResponsesV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGatewayResponsesV2Response struct{}"
	}

	return strings.Join([]string{"ListGatewayResponsesV2Response", string(data)}, " ")
}
