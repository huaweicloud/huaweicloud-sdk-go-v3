package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ThrottleSpecialBase struct {

	// 特殊配置的编号
	Id *string `json:"id,omitempty"`

	// 特殊对象在流控时间内能够访问API的最大次数限制
	CallLimits *int32 `json:"call_limits,omitempty"`

	// 设置时间
	ApplyTime *sdktime.SdkTime `json:"apply_time,omitempty"`

	// 作用的APP名称
	AppName *string `json:"app_name,omitempty"`

	// 作用的APP编号
	AppId *string `json:"app_id,omitempty"`
}

func (o ThrottleSpecialBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThrottleSpecialBase struct{}"
	}

	return strings.Join([]string{"ThrottleSpecialBase", string(data)}, " ")
}
