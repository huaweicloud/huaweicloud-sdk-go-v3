package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDbObjectCollectionStatusRequest Request Object
type ShowDbObjectCollectionStatusRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ShowDbObjectCollectionStatusRequestXLanguage `json:"X-Language,omitempty"`

	// 对象信息采集的ID，提交查询数据库对象信息接口返回的ID。
	QueryId string `json:"query_id"`
}

func (o ShowDbObjectCollectionStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDbObjectCollectionStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowDbObjectCollectionStatusRequest", string(data)}, " ")
}

type ShowDbObjectCollectionStatusRequestXLanguage struct {
	value string
}

type ShowDbObjectCollectionStatusRequestXLanguageEnum struct {
	EN_US ShowDbObjectCollectionStatusRequestXLanguage
	ZH_CN ShowDbObjectCollectionStatusRequestXLanguage
}

func GetShowDbObjectCollectionStatusRequestXLanguageEnum() ShowDbObjectCollectionStatusRequestXLanguageEnum {
	return ShowDbObjectCollectionStatusRequestXLanguageEnum{
		EN_US: ShowDbObjectCollectionStatusRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowDbObjectCollectionStatusRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowDbObjectCollectionStatusRequestXLanguage) Value() string {
	return c.value
}

func (c ShowDbObjectCollectionStatusRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDbObjectCollectionStatusRequestXLanguage) UnmarshalJSON(b []byte) error {
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
