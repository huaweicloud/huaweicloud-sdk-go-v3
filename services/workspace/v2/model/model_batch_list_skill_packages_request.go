package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchListSkillPackagesRequest Request Object
type BatchListSkillPackagesRequest struct {

	// 语言，用于国际化。 - en-us：英文 - zh-cn：中文
	XLanguage *BatchListSkillPackagesRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchListSkillPackagesReq `json:"body,omitempty"`
}

func (o BatchListSkillPackagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListSkillPackagesRequest struct{}"
	}

	return strings.Join([]string{"BatchListSkillPackagesRequest", string(data)}, " ")
}

type BatchListSkillPackagesRequestXLanguage struct {
	value string
}

type BatchListSkillPackagesRequestXLanguageEnum struct {
	EN_US BatchListSkillPackagesRequestXLanguage
	ZH_CN BatchListSkillPackagesRequestXLanguage
}

func GetBatchListSkillPackagesRequestXLanguageEnum() BatchListSkillPackagesRequestXLanguageEnum {
	return BatchListSkillPackagesRequestXLanguageEnum{
		EN_US: BatchListSkillPackagesRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchListSkillPackagesRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchListSkillPackagesRequestXLanguage) Value() string {
	return c.value
}

func (c BatchListSkillPackagesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchListSkillPackagesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
