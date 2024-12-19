package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListRecycleInstancesDetailsRequest Request Object
type ListRecycleInstancesDetailsRequest struct {

	// 语言。默认值：en-us。
	XLanguage *ListRecycleInstancesDetailsRequestXLanguage `json:"X-Language,omitempty"`

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为50，不能为负数，最小值为1，最大值为50。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListRecycleInstancesDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecycleInstancesDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListRecycleInstancesDetailsRequest", string(data)}, " ")
}

type ListRecycleInstancesDetailsRequestXLanguage struct {
	value string
}

type ListRecycleInstancesDetailsRequestXLanguageEnum struct {
	ZH_CN ListRecycleInstancesDetailsRequestXLanguage
	EN_US ListRecycleInstancesDetailsRequestXLanguage
}

func GetListRecycleInstancesDetailsRequestXLanguageEnum() ListRecycleInstancesDetailsRequestXLanguageEnum {
	return ListRecycleInstancesDetailsRequestXLanguageEnum{
		ZH_CN: ListRecycleInstancesDetailsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListRecycleInstancesDetailsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListRecycleInstancesDetailsRequestXLanguage) Value() string {
	return c.value
}

func (c ListRecycleInstancesDetailsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRecycleInstancesDetailsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
