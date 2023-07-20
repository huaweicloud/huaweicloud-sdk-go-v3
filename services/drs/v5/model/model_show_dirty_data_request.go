package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDirtyDataRequest Request Object
type ShowDirtyDataRequest struct {

	// 请求语言类型。
	XLanguage *ShowDirtyDataRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`

	// 开始时间，UTC时间，例如：2020-09-01T18:50:20Z
	BeginTime *string `json:"begin_time,omitempty"`

	// 结束时间，UTC时间，例如：2020-09-01T19:50:20Z
	EndTime *string `json:"end_time,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset 大于等于 0。默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。默认为10，取值范围【1-1000】
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowDirtyDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDirtyDataRequest struct{}"
	}

	return strings.Join([]string{"ShowDirtyDataRequest", string(data)}, " ")
}

type ShowDirtyDataRequestXLanguage struct {
	value string
}

type ShowDirtyDataRequestXLanguageEnum struct {
	EN_US ShowDirtyDataRequestXLanguage
	ZH_CN ShowDirtyDataRequestXLanguage
}

func GetShowDirtyDataRequestXLanguageEnum() ShowDirtyDataRequestXLanguageEnum {
	return ShowDirtyDataRequestXLanguageEnum{
		EN_US: ShowDirtyDataRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowDirtyDataRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowDirtyDataRequestXLanguage) Value() string {
	return c.value
}

func (c ShowDirtyDataRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDirtyDataRequestXLanguage) UnmarshalJSON(b []byte) error {
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
