package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListRequestThrottlingPoliciesBindedToApiV2Response struct {
	// 本次返回的列表长度

	Size int32 `json:"size"`
	// 满足条件的记录数

	Total int64 `json:"total"`
	// 本次查询返回的流控策略列表

	Throttles      *[]ThrottleBindingThrottleResp `json:"throttles,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListRequestThrottlingPoliciesBindedToApiV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListRequestThrottlingPoliciesBindedToApiV2Response struct{}"
	}

	return strings.Join([]string{"ListRequestThrottlingPoliciesBindedToApiV2Response", string(data)}, " ")
}
