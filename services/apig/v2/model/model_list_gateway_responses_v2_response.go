package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListGatewayResponsesV2Response struct {
	// 本次返回的列表长度

	Size int32 `json:"size"`
	// 满足条件的记录数

	Total int64 `json:"total"`
	// 响应列表

	Responses      *[]ResponseInfoResp `json:"responses,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListGatewayResponsesV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListGatewayResponsesV2Response struct{}"
	}

	return strings.Join([]string{"ListGatewayResponsesV2Response", string(data)}, " ")
}
