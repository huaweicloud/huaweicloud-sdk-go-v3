package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListApplicableInstancesRequest Request Object
type ListApplicableInstancesRequest struct {

	// 语言。
	XLanguage *ListApplicableInstancesRequestXLanguage `json:"X-Language,omitempty"`

	// 参数配置模板ID。
	ConfigId string `json:"config_id"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListApplicableInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicableInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListApplicableInstancesRequest", string(data)}, " ")
}

type ListApplicableInstancesRequestXLanguage struct {
	value string
}

type ListApplicableInstancesRequestXLanguageEnum struct {
	ZH_CN ListApplicableInstancesRequestXLanguage
	EN_US ListApplicableInstancesRequestXLanguage
}

func GetListApplicableInstancesRequestXLanguageEnum() ListApplicableInstancesRequestXLanguageEnum {
	return ListApplicableInstancesRequestXLanguageEnum{
		ZH_CN: ListApplicableInstancesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListApplicableInstancesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListApplicableInstancesRequestXLanguage) Value() string {
	return c.value
}

func (c ListApplicableInstancesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApplicableInstancesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
