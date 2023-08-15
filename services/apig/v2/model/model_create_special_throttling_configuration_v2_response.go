package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSpecialThrottlingConfigurationV2Response Response Object
type CreateSpecialThrottlingConfigurationV2Response struct {

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

	// 特殊对象的身份标识
	ObjectId *string `json:"object_id,omitempty"`

	// 特殊对象类型：APP、USER
	ObjectType *string `json:"object_type,omitempty"`

	// 作用的APP或租户的名称
	ObjectName *string `json:"object_name,omitempty"`

	// 流控策略编号
	ThrottleId     *string `json:"throttle_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSpecialThrottlingConfigurationV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSpecialThrottlingConfigurationV2Response struct{}"
	}

	return strings.Join([]string{"CreateSpecialThrottlingConfigurationV2Response", string(data)}, " ")
}
