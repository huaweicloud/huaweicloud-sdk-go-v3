package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListHealthReportTaskRequest Request Object
type ListHealthReportTaskRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 开始时间（Unix timestamp），单位：毫秒。
	StartAt int64 `json:"start_at"`

	// 结束时间（Unix timestamp），单位：毫秒。
	EndAt int64 `json:"end_at"`

	// 偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 每页记录数，默认为10，最大取值200。
	Limit *int32 `json:"limit,omitempty"`

	// 请求语言类型。
	XLanguage *ListHealthReportTaskRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListHealthReportTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHealthReportTaskRequest struct{}"
	}

	return strings.Join([]string{"ListHealthReportTaskRequest", string(data)}, " ")
}

type ListHealthReportTaskRequestXLanguage struct {
	value string
}

type ListHealthReportTaskRequestXLanguageEnum struct {
	EN_US ListHealthReportTaskRequestXLanguage
	ZH_CN ListHealthReportTaskRequestXLanguage
}

func GetListHealthReportTaskRequestXLanguageEnum() ListHealthReportTaskRequestXLanguageEnum {
	return ListHealthReportTaskRequestXLanguageEnum{
		EN_US: ListHealthReportTaskRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListHealthReportTaskRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListHealthReportTaskRequestXLanguage) Value() string {
	return c.value
}

func (c ListHealthReportTaskRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHealthReportTaskRequestXLanguage) UnmarshalJSON(b []byte) error {
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
