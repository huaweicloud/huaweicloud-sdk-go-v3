package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowTopObjectsRequest Request Object
type ShowTopObjectsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// top行数
	Top *int32 `json:"top,omitempty"`

	// 数据库名
	DatabaseName *string `json:"database_name,omitempty"`

	// 排序规则
	Order *string `json:"order,omitempty"`

	// 语言
	XLanguage *ShowTopObjectsRequestXLanguage `json:"X-Language,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为1000，不能为负数，最小值为1，最大值为1000。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowTopObjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopObjectsRequest struct{}"
	}

	return strings.Join([]string{"ShowTopObjectsRequest", string(data)}, " ")
}

type ShowTopObjectsRequestXLanguage struct {
	value string
}

type ShowTopObjectsRequestXLanguageEnum struct {
	ZH_CN ShowTopObjectsRequestXLanguage
	EN_US ShowTopObjectsRequestXLanguage
}

func GetShowTopObjectsRequestXLanguageEnum() ShowTopObjectsRequestXLanguageEnum {
	return ShowTopObjectsRequestXLanguageEnum{
		ZH_CN: ShowTopObjectsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowTopObjectsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowTopObjectsRequestXLanguage) Value() string {
	return c.value
}

func (c ShowTopObjectsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTopObjectsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
