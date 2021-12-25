package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowApplicationV3Request struct {
	// 语言类型 中文:zh-cn 英文:en-us

	XLanguage *ShowApplicationV3RequestXLanguage `json:"X-Language,omitempty"`
	// 应用id

	ApplicationId string `json:"application_id"`
}

func (o ShowApplicationV3Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApplicationV3Request struct{}"
	}

	return strings.Join([]string{"ShowApplicationV3Request", string(data)}, " ")
}

type ShowApplicationV3RequestXLanguage struct {
	value string
}

type ShowApplicationV3RequestXLanguageEnum struct {
	ZH_CN ShowApplicationV3RequestXLanguage
	EN_US ShowApplicationV3RequestXLanguage
}

func GetShowApplicationV3RequestXLanguageEnum() ShowApplicationV3RequestXLanguageEnum {
	return ShowApplicationV3RequestXLanguageEnum{
		ZH_CN: ShowApplicationV3RequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowApplicationV3RequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowApplicationV3RequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApplicationV3RequestXLanguage) UnmarshalJSON(b []byte) error {
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
