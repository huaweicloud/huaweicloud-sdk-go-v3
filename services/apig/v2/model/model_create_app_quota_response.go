package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// CreateAppQuotaResponse Response Object
type CreateAppQuotaResponse struct {

	// 凭据配额编号
	AppQuotaId *string `json:"app_quota_id,omitempty"`

	// 配额名称。支持汉字，英文，数字，下划线，且只能以英文和汉字开头，3-255字符
	Name *string `json:"name,omitempty"`

	// 凭据配额的访问次数限制
	CallLimits *int32 `json:"call_limits,omitempty"`

	// 限定时间单位：SECOND:秒、MINUTE:分、HOUR:时、DAY:天
	TimeUnit *CreateAppQuotaResponseTimeUnit `json:"time_unit,omitempty"`

	// 配额的限定时间值
	TimeInterval *int32 `json:"time_interval,omitempty"`

	// 参数说明和描述
	Remark *string `json:"remark,omitempty"`

	// 首次配额重置时间点，不配置默认为首次调用时间计算
	ResetTime *string `json:"reset_time,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 配额策略已绑定应用数量
	BoundAppNum    *int32 `json:"bound_app_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateAppQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppQuotaResponse struct{}"
	}

	return strings.Join([]string{"CreateAppQuotaResponse", string(data)}, " ")
}

type CreateAppQuotaResponseTimeUnit struct {
	value string
}

type CreateAppQuotaResponseTimeUnitEnum struct {
	SECOND CreateAppQuotaResponseTimeUnit
	MINUTE CreateAppQuotaResponseTimeUnit
	HOUR   CreateAppQuotaResponseTimeUnit
	DAY    CreateAppQuotaResponseTimeUnit
}

func GetCreateAppQuotaResponseTimeUnitEnum() CreateAppQuotaResponseTimeUnitEnum {
	return CreateAppQuotaResponseTimeUnitEnum{
		SECOND: CreateAppQuotaResponseTimeUnit{
			value: "SECOND",
		},
		MINUTE: CreateAppQuotaResponseTimeUnit{
			value: "MINUTE",
		},
		HOUR: CreateAppQuotaResponseTimeUnit{
			value: "HOUR",
		},
		DAY: CreateAppQuotaResponseTimeUnit{
			value: "DAY",
		},
	}
}

func (c CreateAppQuotaResponseTimeUnit) Value() string {
	return c.value
}

func (c CreateAppQuotaResponseTimeUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAppQuotaResponseTimeUnit) UnmarshalJSON(b []byte) error {
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
