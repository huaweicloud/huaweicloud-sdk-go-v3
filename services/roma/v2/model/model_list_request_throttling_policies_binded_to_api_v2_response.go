package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRequestThrottlingPoliciesBindedToApiV2Response struct {

	// 本次返回的列表长度
	Size int32 `json:"size" xml:"size"`

	// 满足条件的记录数
	Total int64 `json:"total" xml:"total"`

	// 本次查询返回的流控策略列表
	Throttles      *[]ThrottleForApi `json:"throttles,omitempty" xml:"throttles"`
	HttpStatusCode int               `json:"-"`
}

func (o ListRequestThrottlingPoliciesBindedToApiV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRequestThrottlingPoliciesBindedToApiV2Response struct{}"
	}

	return strings.Join([]string{"ListRequestThrottlingPoliciesBindedToApiV2Response", string(data)}, " ")
}
