package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"errors"
	"strings"
)

type AppQuotaInfo struct {
	// 客户端配额编号

	AppQuotaId *string `json:"app_quota_id,omitempty"`
	// 配额名称。支持汉字，英文，数字，下划线，且只能以英文和汉字开头，3-255字符

	Name *string `json:"name,omitempty"`
	// 客户端配额的访问次数限制

	CallLimits *int32 `json:"call_limits,omitempty"`
	// 限定时间单位：SECOND:秒、MINUTE:分、HOURE:时、DAY:天

	TimeUnit *AppQuotaInfoTimeUnit `json:"time_unit,omitempty"`
	// 配额的限定时间值

	TimeInterval *int32 `json:"time_interval,omitempty"`
	// 参数说明和描述

	Remark *string `json:"remark,omitempty"`
	// 首次配额重置时间点，不配置默认为首次调用时间计算

	ResetTime *string `json:"reset_time,omitempty"`
	// 创建时间

	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`
	// 配额策略已绑定应用数量

	BoundAppNum *int32 `json:"bound_app_num,omitempty"`
}

func (o AppQuotaInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppQuotaInfo struct{}"
	}

	return strings.Join([]string{"AppQuotaInfo", string(data)}, " ")
}

type AppQuotaInfoTimeUnit struct {
	value string
}

type AppQuotaInfoTimeUnitEnum struct {
	SECOND AppQuotaInfoTimeUnit
	MINUTE AppQuotaInfoTimeUnit
	HOUR   AppQuotaInfoTimeUnit
	DAY    AppQuotaInfoTimeUnit
}

func GetAppQuotaInfoTimeUnitEnum() AppQuotaInfoTimeUnitEnum {
	return AppQuotaInfoTimeUnitEnum{
		SECOND: AppQuotaInfoTimeUnit{
			value: "SECOND",
		},
		MINUTE: AppQuotaInfoTimeUnit{
			value: "MINUTE",
		},
		HOUR: AppQuotaInfoTimeUnit{
			value: "HOUR",
		},
		DAY: AppQuotaInfoTimeUnit{
			value: "DAY",
		},
	}
}

func (c AppQuotaInfoTimeUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AppQuotaInfoTimeUnit) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
