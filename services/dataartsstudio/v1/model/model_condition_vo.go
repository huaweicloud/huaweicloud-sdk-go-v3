package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type ConditionVo struct {

	// 编码，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 字段名
	NameEn string `json:"name_en"`

	// 业务属性。
	NameCh string `json:"name_ch"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 创建人。
	CreateBy *string `json:"create_by,omitempty"`

	// 限定计算方法。 枚举值：   - LAST_YEAR: 前一年   - CURRENT_YEAR: 本年   - BETWEEN_YEAR: 自定义年区间   - LAST_MONTH: 前一月   - CURRENT_MONTH: 本月   - BETWEEN_MONTH: 自定义月区间   - LAST_DAY: 前一天   - CURRENT_DAY: 本日   - BETWEEN_DAY: 自定义日区间   - LAST_HOUR: 上一小时   - CURRENT_HOUR: 当前小时   - BETWEEN_HOUR: 自定义小时区间   - LAST_MINUTE: 上一分钟   - CURRENT_MINUTE: 当前分钟   - BETWEEN_MINUTE: 自定义分钟区间
	ConditionFn *ConditionVoConditionFn `json:"condition_fn,omitempty"`

	// 限定计算参数。
	ConditionFnParam *string `json:"condition_fn_param,omitempty"`

	Status *BizStatusEnum `json:"status,omitempty"`

	// 创建时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	ApprovalInfo *ApprovalVo `json:"approval_info,omitempty"`

	NewBiz *BizVersionManageVo `json:"new_biz,omitempty"`

	// 基准时间。
	BaseTime *int32 `json:"base_time,omitempty"`
}

func (o ConditionVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConditionVo struct{}"
	}

	return strings.Join([]string{"ConditionVo", string(data)}, " ")
}

type ConditionVoConditionFn struct {
	value string
}

type ConditionVoConditionFnEnum struct {
	LAST_YEAR      ConditionVoConditionFn
	CURRENT_YEAR   ConditionVoConditionFn
	BETWEEN_YEAR   ConditionVoConditionFn
	LAST_MONTH     ConditionVoConditionFn
	CURRENT_MONTH  ConditionVoConditionFn
	BETWEEN_MONTH  ConditionVoConditionFn
	LAST_DAY       ConditionVoConditionFn
	CURRENT_DAY    ConditionVoConditionFn
	BETWEEN_DAY    ConditionVoConditionFn
	LAST_HOUR      ConditionVoConditionFn
	CURRENT_HOUR   ConditionVoConditionFn
	BETWEEN_HOUR   ConditionVoConditionFn
	LAST_MINUTE    ConditionVoConditionFn
	CURRENT_MINUTE ConditionVoConditionFn
	BETWEEN_MINUTE ConditionVoConditionFn
}

func GetConditionVoConditionFnEnum() ConditionVoConditionFnEnum {
	return ConditionVoConditionFnEnum{
		LAST_YEAR: ConditionVoConditionFn{
			value: "LAST_YEAR",
		},
		CURRENT_YEAR: ConditionVoConditionFn{
			value: "CURRENT_YEAR",
		},
		BETWEEN_YEAR: ConditionVoConditionFn{
			value: "BETWEEN_YEAR",
		},
		LAST_MONTH: ConditionVoConditionFn{
			value: "LAST_MONTH",
		},
		CURRENT_MONTH: ConditionVoConditionFn{
			value: "CURRENT_MONTH",
		},
		BETWEEN_MONTH: ConditionVoConditionFn{
			value: "BETWEEN_MONTH",
		},
		LAST_DAY: ConditionVoConditionFn{
			value: "LAST_DAY",
		},
		CURRENT_DAY: ConditionVoConditionFn{
			value: "CURRENT_DAY",
		},
		BETWEEN_DAY: ConditionVoConditionFn{
			value: "BETWEEN_DAY",
		},
		LAST_HOUR: ConditionVoConditionFn{
			value: "LAST_HOUR",
		},
		CURRENT_HOUR: ConditionVoConditionFn{
			value: "CURRENT_HOUR",
		},
		BETWEEN_HOUR: ConditionVoConditionFn{
			value: "BETWEEN_HOUR",
		},
		LAST_MINUTE: ConditionVoConditionFn{
			value: "LAST_MINUTE",
		},
		CURRENT_MINUTE: ConditionVoConditionFn{
			value: "CURRENT_MINUTE",
		},
		BETWEEN_MINUTE: ConditionVoConditionFn{
			value: "BETWEEN_MINUTE",
		},
	}
}

func (c ConditionVoConditionFn) Value() string {
	return c.value
}

func (c ConditionVoConditionFn) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConditionVoConditionFn) UnmarshalJSON(b []byte) error {
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
