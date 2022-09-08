package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRequestThrottlingPolicyV2Response struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 本次查询到的流控策略列表
	Throttles      *[]ThrottlesInfo `json:"throttles,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListRequestThrottlingPolicyV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRequestThrottlingPolicyV2Response struct{}"
	}

	return strings.Join([]string{"ListRequestThrottlingPolicyV2Response", string(data)}, " ")
}
