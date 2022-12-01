package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListDbObjectsRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ListDbObjectsRequestXLanguage `json:"X-Language,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`

	// 查询对象信息类型。取值： - source：查询源库对象信息。 - modified：查询已选择的（未下发与已同步的）对象信息。 - synchronized：查询已同步的（已下发的）对象信息 。
	Type string `json:"type"`
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
