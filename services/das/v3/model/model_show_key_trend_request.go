package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowKeyTrendRequest Request Object
type ShowKeyTrendRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 统计维度。size/storage/qps/traffic，默认size
	StatType *ShowKeyTrendRequestStatType `json:"stat_type,omitempty"`

	// key名称（需URL编码）
	Key string `json:"key"`

	// 开始时间（毫秒时间戳）
	StartAt int64 `json:"start_at"`

	// 结束时间（毫秒时间戳）
	EndAt int64 `json:"end_at"`

	// 请求语言类型。en-us：英文。zh-cn：中文。
	XLanguage *ShowKeyTrendRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowKeyTrendRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKeyTrendRequest struct{}"
	}

	return strings.Join([]string{"ShowKeyTrendRequest", string(data)}, " ")
}

type ShowKeyTrendRequestStatType struct {
	value string
}

type ShowKeyTrendRequestStatTypeEnum struct {
	SIZE    ShowKeyTrendRequestStatType
	STORAGE ShowKeyTrendRequestStatType
	QPS     ShowKeyTrendRequestStatType
	TRAFFIC ShowKeyTrendRequestStatType
}

func GetShowKeyTrendRequestStatTypeEnum() ShowKeyTrendRequestStatTypeEnum {
	return ShowKeyTrendRequestStatTypeEnum{
		SIZE: ShowKeyTrendRequestStatType{
			value: "size",
		},
		STORAGE: ShowKeyTrendRequestStatType{
			value: "storage",
		},
		QPS: ShowKeyTrendRequestStatType{
			value: "qps",
		},
		TRAFFIC: ShowKeyTrendRequestStatType{
			value: "traffic",
		},
	}
}

func (c ShowKeyTrendRequestStatType) Value() string {
	return c.value
}

func (c ShowKeyTrendRequestStatType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowKeyTrendRequestStatType) UnmarshalJSON(b []byte) error {
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

type ShowKeyTrendRequestXLanguage struct {
	value string
}

type ShowKeyTrendRequestXLanguageEnum struct {
	ZH_CN ShowKeyTrendRequestXLanguage
	EN_US ShowKeyTrendRequestXLanguage
}

func GetShowKeyTrendRequestXLanguageEnum() ShowKeyTrendRequestXLanguageEnum {
	return ShowKeyTrendRequestXLanguageEnum{
		ZH_CN: ShowKeyTrendRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowKeyTrendRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowKeyTrendRequestXLanguage) Value() string {
	return c.value
}

func (c ShowKeyTrendRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowKeyTrendRequestXLanguage) UnmarshalJSON(b []byte) error {
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
