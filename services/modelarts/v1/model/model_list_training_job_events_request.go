package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTrainingJobEventsRequest Request Object
type ListTrainingJobEventsRequest struct {

	// 训练作业ID。获取方法请参见[查询训练作业列表](ListTrainingJobs.xml)。
	TrainingJobId string `json:"training_job_id"`

	// **参数解释**：分页列表的起始页。 **约束限制**：最小为0。例如设置为0，则表示从第一页开始查询。 **取值范围**：不涉及。 **默认取值**：默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 指定每一页返回的最大条目数，取值范围[1,100]，默认为50。
	Limit *int32 `json:"limit,omitempty"`

	// instance order
	Order *ListTrainingJobEventsRequestOrder `json:"order,omitempty"`

	// 开始时间，需要与结束时间一起传入。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间，需要与开始时间一起传入。
	EndTime *string `json:"end_time,omitempty"`

	// 语言。
	XLanguage *ListTrainingJobEventsRequestXLanguage `json:"X-Language,omitempty"`

	// 指定返回的事件级别，取值范围[Info Error Warning]。
	Level *string `json:"level,omitempty"`

	// 指定事件信息包含的内容，最长256个字符。
	Pattern *string `json:"pattern,omitempty"`
}

func (o ListTrainingJobEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTrainingJobEventsRequest struct{}"
	}

	return strings.Join([]string{"ListTrainingJobEventsRequest", string(data)}, " ")
}

type ListTrainingJobEventsRequestOrder struct {
	value string
}

type ListTrainingJobEventsRequestOrderEnum struct {
	ASC  ListTrainingJobEventsRequestOrder
	DESC ListTrainingJobEventsRequestOrder
}

func GetListTrainingJobEventsRequestOrderEnum() ListTrainingJobEventsRequestOrderEnum {
	return ListTrainingJobEventsRequestOrderEnum{
		ASC: ListTrainingJobEventsRequestOrder{
			value: "asc",
		},
		DESC: ListTrainingJobEventsRequestOrder{
			value: "desc",
		},
	}
}

func (c ListTrainingJobEventsRequestOrder) Value() string {
	return c.value
}

func (c ListTrainingJobEventsRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTrainingJobEventsRequestOrder) UnmarshalJSON(b []byte) error {
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

type ListTrainingJobEventsRequestXLanguage struct {
	value string
}

type ListTrainingJobEventsRequestXLanguageEnum struct {
	ZH_CN ListTrainingJobEventsRequestXLanguage
	EN_US ListTrainingJobEventsRequestXLanguage
}

func GetListTrainingJobEventsRequestXLanguageEnum() ListTrainingJobEventsRequestXLanguageEnum {
	return ListTrainingJobEventsRequestXLanguageEnum{
		ZH_CN: ListTrainingJobEventsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListTrainingJobEventsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListTrainingJobEventsRequestXLanguage) Value() string {
	return c.value
}

func (c ListTrainingJobEventsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTrainingJobEventsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
