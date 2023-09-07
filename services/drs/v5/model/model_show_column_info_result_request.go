package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowColumnInfoResultRequest Request Object
type ShowColumnInfoResultRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ShowColumnInfoResultRequestXLanguage `json:"X-Language,omitempty"`

	// 指定数据库表的列信息采集的ID，提交采集指定数据库表的列信息接口返回的ID。
	QueryId string `json:"query_id"`

	// 偏移量，表示查询该偏移量后面的记录。
	Offset *int32 `json:"offset,omitempty"`

	// 查询返回记录的数量限制。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowColumnInfoResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowColumnInfoResultRequest struct{}"
	}

	return strings.Join([]string{"ShowColumnInfoResultRequest", string(data)}, " ")
}

type ShowColumnInfoResultRequestXLanguage struct {
	value string
}

type ShowColumnInfoResultRequestXLanguageEnum struct {
	EN_US ShowColumnInfoResultRequestXLanguage
	ZH_CN ShowColumnInfoResultRequestXLanguage
}

func GetShowColumnInfoResultRequestXLanguageEnum() ShowColumnInfoResultRequestXLanguageEnum {
	return ShowColumnInfoResultRequestXLanguageEnum{
		EN_US: ShowColumnInfoResultRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowColumnInfoResultRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowColumnInfoResultRequestXLanguage) Value() string {
	return c.value
}

func (c ShowColumnInfoResultRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowColumnInfoResultRequestXLanguage) UnmarshalJSON(b []byte) error {
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
