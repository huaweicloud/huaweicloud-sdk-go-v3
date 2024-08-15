package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowBatchUpgradeCandidateVersionsRequest Request Object
type ShowBatchUpgradeCandidateVersionsRequest struct {

	// 语言[zh-cn, en-us]。
	XLanguage *ShowBatchUpgradeCandidateVersionsRequestXLanguage `json:"X-Language,omitempty"`

	Body *UpgradeInstancesRequestBody `json:"body,omitempty"`
}

func (o ShowBatchUpgradeCandidateVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchUpgradeCandidateVersionsRequest struct{}"
	}

	return strings.Join([]string{"ShowBatchUpgradeCandidateVersionsRequest", string(data)}, " ")
}

type ShowBatchUpgradeCandidateVersionsRequestXLanguage struct {
	value string
}

type ShowBatchUpgradeCandidateVersionsRequestXLanguageEnum struct {
	ZH_CN ShowBatchUpgradeCandidateVersionsRequestXLanguage
	EN_US ShowBatchUpgradeCandidateVersionsRequestXLanguage
}

func GetShowBatchUpgradeCandidateVersionsRequestXLanguageEnum() ShowBatchUpgradeCandidateVersionsRequestXLanguageEnum {
	return ShowBatchUpgradeCandidateVersionsRequestXLanguageEnum{
		ZH_CN: ShowBatchUpgradeCandidateVersionsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowBatchUpgradeCandidateVersionsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowBatchUpgradeCandidateVersionsRequestXLanguage) Value() string {
	return c.value
}

func (c ShowBatchUpgradeCandidateVersionsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowBatchUpgradeCandidateVersionsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
