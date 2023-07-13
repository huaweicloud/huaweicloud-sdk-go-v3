package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AppQuotaCreate struct {

	// 配额名称。支持汉字，英文，数字，下划线，且只能以英文和汉字开头
	Name string `json:"name"`

	// 客户端配额的访问次数限制
	CallLimits int32 `json:"call_limits"`

	// 限定时间单位：SECOND:秒、MINUTE:分、HOURE:时、DAY:天
	TimeUnit AppQuotaCreateTimeUnit `json:"time_unit"`

	// 流控的限定时间值
	TimeInterval int32 `json:"time_interval"`

	// 首次配额重置时间点，不配置默认为首次调用时间计算
	ResetTime *string `json:"reset_time,omitempty"`

	// 参数说明和描述。  不支持<，>字符
	Remark *string `json:"remark,omitempty"`
}

func (o AppQuotaCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppQuotaCreate struct{}"
	}

	return strings.Join([]string{"AppQuotaCreate", string(data)}, " ")
}

type AppQuotaCreateTimeUnit struct {
	value string
}

type AppQuotaCreateTimeUnitEnum struct {
	SECOND AppQuotaCreateTimeUnit
	MINUTE AppQuotaCreateTimeUnit
	HOUR   AppQuotaCreateTimeUnit
	DAY    AppQuotaCreateTimeUnit
}

func GetAppQuotaCreateTimeUnitEnum() AppQuotaCreateTimeUnitEnum {
	return AppQuotaCreateTimeUnitEnum{
		SECOND: AppQuotaCreateTimeUnit{
			value: "SECOND",
		},
		MINUTE: AppQuotaCreateTimeUnit{
			value: "MINUTE",
		},
		HOUR: AppQuotaCreateTimeUnit{
			value: "HOUR",
		},
		DAY: AppQuotaCreateTimeUnit{
			value: "DAY",
		},
	}
}

func (c AppQuotaCreateTimeUnit) Value() string {
	return c.value
}

func (c AppQuotaCreateTimeUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AppQuotaCreateTimeUnit) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
