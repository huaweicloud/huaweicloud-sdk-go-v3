package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListApisBindedToRequestThrottlingPolicyV2Response struct {
	// 满足条件的API总数
	Total *int32 `json:"total,omitempty"`
	// 本次返回的API列表长度
	Size *int32 `json:"size,omitempty"`
	// 本次查询返回的API列表
	Apis           *[]ThrottleBindingApiResp `json:"apis,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListApisBindedToRequestThrottlingPolicyV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListApisBindedToRequestThrottlingPolicyV2Response struct{}"
	}

	return strings.Join([]string{"ListApisBindedToRequestThrottlingPolicyV2Response", string(data)}, " ")
}
