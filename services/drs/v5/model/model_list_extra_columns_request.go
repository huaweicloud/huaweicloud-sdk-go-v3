package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListExtraColumnsRequest Request Object
type ListExtraColumnsRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ListExtraColumnsRequestXLanguage `json:"X-Language,omitempty"`

	// 是否仅查询已下发的加工对象，默认为否。
	IsOnlyShowSent *bool `json:"is_only_show_sent,omitempty"`

	// 偏移量，表示查询该偏移量后面的记录，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询返回记录的数量限制，默认为10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListExtraColumnsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExtraColumnsRequest struct{}"
	}

	return strings.Join([]string{"ListExtraColumnsRequest", string(data)}, " ")
}

type ListExtraColumnsRequestXLanguage struct {
	value string
}

type ListExtraColumnsRequestXLanguageEnum struct {
	EN_US ListExtraColumnsRequestXLanguage
	ZH_CN ListExtraColumnsRequestXLanguage
}

func GetListExtraColumnsRequestXLanguageEnum() ListExtraColumnsRequestXLanguageEnum {
	return ListExtraColumnsRequestXLanguageEnum{
		EN_US: ListExtraColumnsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListExtraColumnsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListExtraColumnsRequestXLanguage) Value() string {
	return c.value
}

func (c ListExtraColumnsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListExtraColumnsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
