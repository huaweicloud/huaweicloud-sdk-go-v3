package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowUpgradeCandidateVersionsDetailsRequest Request Object
type ShowUpgradeCandidateVersionsDetailsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 语言[zh-cn, en-us]。
	XLanguage *ShowUpgradeCandidateVersionsDetailsRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowUpgradeCandidateVersionsDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUpgradeCandidateVersionsDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowUpgradeCandidateVersionsDetailsRequest", string(data)}, " ")
}

type ShowUpgradeCandidateVersionsDetailsRequestXLanguage struct {
	value string
}

type ShowUpgradeCandidateVersionsDetailsRequestXLanguageEnum struct {
	ZH_CN ShowUpgradeCandidateVersionsDetailsRequestXLanguage
	EN_US ShowUpgradeCandidateVersionsDetailsRequestXLanguage
}

func GetShowUpgradeCandidateVersionsDetailsRequestXLanguageEnum() ShowUpgradeCandidateVersionsDetailsRequestXLanguageEnum {
	return ShowUpgradeCandidateVersionsDetailsRequestXLanguageEnum{
		ZH_CN: ShowUpgradeCandidateVersionsDetailsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowUpgradeCandidateVersionsDetailsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowUpgradeCandidateVersionsDetailsRequestXLanguage) Value() string {
	return c.value
}

func (c ShowUpgradeCandidateVersionsDetailsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowUpgradeCandidateVersionsDetailsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
