package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDbObjectsRequest Request Object
type ListDbObjectsRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ListDbObjectsRequestXLanguage `json:"X-Language,omitempty"`

	// 偏移量，表示查询该偏移量后面的记录。
	Offset *int32 `json:"offset,omitempty"`

	// 查询返回记录的数量限制。
	Limit *int32 `json:"limit,omitempty"`

	// 查询对象信息类型。 取值： - source：查询源库对象信息。 - modified：查询已选择的（已同步的和未下发的）对象信息。 - synchronized：查询已同步的（已下发的）对象信息 ， 使用场景在任务处于全量中或者增量中。
	Type string `json:"type"`

	// 查询指定库的信息。
	DbNames *[]string `json:"db_names,omitempty"`
}

func (o ListDbObjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbObjectsRequest struct{}"
	}

	return strings.Join([]string{"ListDbObjectsRequest", string(data)}, " ")
}

type ListDbObjectsRequestXLanguage struct {
	value string
}

type ListDbObjectsRequestXLanguageEnum struct {
	EN_US ListDbObjectsRequestXLanguage
	ZH_CN ListDbObjectsRequestXLanguage
}

func GetListDbObjectsRequestXLanguageEnum() ListDbObjectsRequestXLanguageEnum {
	return ListDbObjectsRequestXLanguageEnum{
		EN_US: ListDbObjectsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListDbObjectsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListDbObjectsRequestXLanguage) Value() string {
	return c.value
}

func (c ListDbObjectsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDbObjectsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
