package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListCloudDbaInstancesRequest Request Object
type ListCloudDbaInstancesRequest struct {

	// 数据库类型。
	DatastoreType string `json:"datastore_type"`

	// 偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 每页记录数，默认为20，最大取值200。
	Limit *int32 `json:"limit,omitempty"`

	// 请求语言类型。
	XLanguage *ListCloudDbaInstancesRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListCloudDbaInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudDbaInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListCloudDbaInstancesRequest", string(data)}, " ")
}

type ListCloudDbaInstancesRequestXLanguage struct {
	value string
}

type ListCloudDbaInstancesRequestXLanguageEnum struct {
	EN_US ListCloudDbaInstancesRequestXLanguage
	ZH_CN ListCloudDbaInstancesRequestXLanguage
}

func GetListCloudDbaInstancesRequestXLanguageEnum() ListCloudDbaInstancesRequestXLanguageEnum {
	return ListCloudDbaInstancesRequestXLanguageEnum{
		EN_US: ListCloudDbaInstancesRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListCloudDbaInstancesRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListCloudDbaInstancesRequestXLanguage) Value() string {
	return c.value
}

func (c ListCloudDbaInstancesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCloudDbaInstancesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
