package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ListHbaInfoHistoryRequest Request Object
type ListHbaInfoHistoryRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ListHbaInfoHistoryRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 开始时间。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 默认当天0点（UTC时区）。
	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`

	// **参数解释**: 结束时间。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 默认当前时间（UTC时区）。
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	// **参数描述** 偏移量。 **约束限制**: 不涉及。 **取值范围** 大于等于0。 **默认值** 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数描述** 每页显示的条目数量。 **约束限制**: 不涉及。 **取值范围** [1, 100] **默认值** 10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListHbaInfoHistoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHbaInfoHistoryRequest struct{}"
	}

	return strings.Join([]string{"ListHbaInfoHistoryRequest", string(data)}, " ")
}

type ListHbaInfoHistoryRequestXLanguage struct {
	value string
}

type ListHbaInfoHistoryRequestXLanguageEnum struct {
	ZH_CN ListHbaInfoHistoryRequestXLanguage
	EN_US ListHbaInfoHistoryRequestXLanguage
}

func GetListHbaInfoHistoryRequestXLanguageEnum() ListHbaInfoHistoryRequestXLanguageEnum {
	return ListHbaInfoHistoryRequestXLanguageEnum{
		ZH_CN: ListHbaInfoHistoryRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListHbaInfoHistoryRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListHbaInfoHistoryRequestXLanguage) Value() string {
	return c.value
}

func (c ListHbaInfoHistoryRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHbaInfoHistoryRequestXLanguage) UnmarshalJSON(b []byte) error {
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
