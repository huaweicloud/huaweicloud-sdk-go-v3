package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateSpecialThrottlingConfigurationV2Response struct {

	// 特殊配置的编号
	Id *string `json:"id,omitempty" xml:"id"`

	// 特殊对象在流控时间内能够访问API的最大次数限制
	CallLimits *int32 `json:"call_limits,omitempty" xml:"call_limits"`

	// 设置时间
	ApplyTime *sdktime.SdkTime `json:"apply_time,omitempty" xml:"apply_time"`

	// 作用的APP名称
	AppName *string `json:"app_name,omitempty" xml:"app_name"`

	// 作用的APP编号
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	// 特殊对象的身份标识
	ObjectId *string `json:"object_id,omitempty" xml:"object_id"`

	// 特殊对象类型：APP、USER
	ObjectType *string `json:"object_type,omitempty" xml:"object_type"`

	// 作用的APP或租户的名称
	ObjectName *string `json:"object_name,omitempty" xml:"object_name"`

	// 流控策略编号
	ThrottleId     *string `json:"throttle_id,omitempty" xml:"throttle_id"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateSpecialThrottlingConfigurationV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSpecialThrottlingConfigurationV2Response struct{}"
	}

	return strings.Join([]string{"UpdateSpecialThrottlingConfigurationV2Response", string(data)}, " ")
}
