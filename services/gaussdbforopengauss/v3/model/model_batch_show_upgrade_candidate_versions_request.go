package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchShowUpgradeCandidateVersionsRequest Request Object
type BatchShowUpgradeCandidateVersionsRequest struct {

	// 语言[zh-cn, en-us]。
	XLanguage *BatchShowUpgradeCandidateVersionsRequestXLanguage `json:"X-Language,omitempty"`

	Body *UpgradeInstancesRequestBody `json:"body,omitempty"`
}

func (o BatchShowUpgradeCandidateVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowUpgradeCandidateVersionsRequest struct{}"
	}

	return strings.Join([]string{"BatchShowUpgradeCandidateVersionsRequest", string(data)}, " ")
}

type BatchShowUpgradeCandidateVersionsRequestXLanguage struct {
	value string
}

type BatchShowUpgradeCandidateVersionsRequestXLanguageEnum struct {
	ZH_CN BatchShowUpgradeCandidateVersionsRequestXLanguage
	EN_US BatchShowUpgradeCandidateVersionsRequestXLanguage
}

func GetBatchShowUpgradeCandidateVersionsRequestXLanguageEnum() BatchShowUpgradeCandidateVersionsRequestXLanguageEnum {
	return BatchShowUpgradeCandidateVersionsRequestXLanguageEnum{
		ZH_CN: BatchShowUpgradeCandidateVersionsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: BatchShowUpgradeCandidateVersionsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c BatchShowUpgradeCandidateVersionsRequestXLanguage) Value() string {
	return c.value
}

func (c BatchShowUpgradeCandidateVersionsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchShowUpgradeCandidateVersionsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
