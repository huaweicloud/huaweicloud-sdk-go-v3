package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type SwitchShardRequest struct {

	// 语言
	XLanguage *SwitchShardRequestXLanguage `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *SwitchShardRequestBody `json:"body,omitempty"`
}

func (o SwitchShardRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchShardRequest struct{}"
	}

	return strings.Join([]string{"SwitchShardRequest", string(data)}, " ")
}

type SwitchShardRequestXLanguage struct {
	value string
}

type SwitchShardRequestXLanguageEnum struct {
	ZH_CN SwitchShardRequestXLanguage
	EN_US SwitchShardRequestXLanguage
}

func GetSwitchShardRequestXLanguageEnum() SwitchShardRequestXLanguageEnum {
	return SwitchShardRequestXLanguageEnum{
		ZH_CN: SwitchShardRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: SwitchShardRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c SwitchShardRequestXLanguage) Value() string {
	return c.value
}

func (c SwitchShardRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SwitchShardRequestXLanguage) UnmarshalJSON(b []byte) error {
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
