package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSpecialThrottlingConfigurationsV2Response Response Object
type ListSpecialThrottlingConfigurationsV2Response struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 本次查询返回的特殊配置列表
	ThrottleSpecials *[]ThrottleSpecialInfo `json:"throttle_specials,omitempty"`
	HttpStatusCode   int                    `json:"-"`
}

func (o ListSpecialThrottlingConfigurationsV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSpecialThrottlingConfigurationsV2Response struct{}"
	}

	return strings.Join([]string{"ListSpecialThrottlingConfigurationsV2Response", string(data)}, " ")
}
