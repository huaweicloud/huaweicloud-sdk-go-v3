package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAspInfosRequest Request Object
type ListAspInfosRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ListAspInfosRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询。 **约束限制**: 不涉及。 **取值范围**: [0，2^31-1] **默认取值**: 默认为0（偏移0条数据，表示从第一条数据开始查询）。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 查询记录数。 **约束限制**: 不涉及。 **取值范围**: [1，100] **默认取值**: 默认为10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 查询开始时间。格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始，Z指时区偏移量，时区中的+号需要进行URL编码，编码为%2B，时区中的-号无需编码。 例如北京时间偏移显示为+0800，start_time=2024-03-16T17:20:33+0800传参时编码为start_time=2024-03-16T17:20:33%2B0800。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**: 查询结束时间。格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始，Z指时区偏移量，时区中的+号需要进行URL编码，编码为%2B，时区中的-号无需编码。 例如北京时间偏移显示为+0800，start_time=2024-03-16T17:20:33+0800传参时编码为start_time=2024-03-16T17:20:33%2B0800。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**: 采集任务ID。填写下发asp采集接口返回的job_id，则可过滤查询对应采集任务的结果。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	JobId *string `json:"job_id,omitempty"`
}

func (o ListAspInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAspInfosRequest struct{}"
	}

	return strings.Join([]string{"ListAspInfosRequest", string(data)}, " ")
}

type ListAspInfosRequestXLanguage struct {
	value string
}

type ListAspInfosRequestXLanguageEnum struct {
	ZH_CN ListAspInfosRequestXLanguage
	EN_US ListAspInfosRequestXLanguage
}

func GetListAspInfosRequestXLanguageEnum() ListAspInfosRequestXLanguageEnum {
	return ListAspInfosRequestXLanguageEnum{
		ZH_CN: ListAspInfosRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListAspInfosRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListAspInfosRequestXLanguage) Value() string {
	return c.value
}

func (c ListAspInfosRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAspInfosRequestXLanguage) UnmarshalJSON(b []byte) error {
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
