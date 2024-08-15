package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowUpgradeCandidateVersionsRequest Request Object
type ShowUpgradeCandidateVersionsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 语言[zh-cn, en-us]。
	XLanguage *ShowUpgradeCandidateVersionsRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowUpgradeCandidateVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUpgradeCandidateVersionsRequest struct{}"
	}

	return strings.Join([]string{"ShowUpgradeCandidateVersionsRequest", string(data)}, " ")
}

type ShowUpgradeCandidateVersionsRequestXLanguage struct {
	value string
}

type ShowUpgradeCandidateVersionsRequestXLanguageEnum struct {
	ZH_CN ShowUpgradeCandidateVersionsRequestXLanguage
	EN_US ShowUpgradeCandidateVersionsRequestXLanguage
}

func GetShowUpgradeCandidateVersionsRequestXLanguageEnum() ShowUpgradeCandidateVersionsRequestXLanguageEnum {
	return ShowUpgradeCandidateVersionsRequestXLanguageEnum{
		ZH_CN: ShowUpgradeCandidateVersionsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowUpgradeCandidateVersionsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowUpgradeCandidateVersionsRequestXLanguage) Value() string {
	return c.value
}

func (c ShowUpgradeCandidateVersionsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowUpgradeCandidateVersionsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
