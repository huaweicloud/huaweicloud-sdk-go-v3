package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type CreateSpaceAnalysisTaskRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 语言

	XLanguage *CreateSpaceAnalysisTaskRequestXLanguage `json:"X-Language,omitempty"`

	Body *CreateSpaceAnalysisTaskBody `json:"body,omitempty"`
}

func (o CreateSpaceAnalysisTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateSpaceAnalysisTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateSpaceAnalysisTaskRequest", string(data)}, " ")
}

type CreateSpaceAnalysisTaskRequestXLanguage struct {
	value string
}

type CreateSpaceAnalysisTaskRequestXLanguageEnum struct {
	ZH_CN CreateSpaceAnalysisTaskRequestXLanguage
	EN_US CreateSpaceAnalysisTaskRequestXLanguage
}

func GetCreateSpaceAnalysisTaskRequestXLanguageEnum() CreateSpaceAnalysisTaskRequestXLanguageEnum {
	return CreateSpaceAnalysisTaskRequestXLanguageEnum{
		ZH_CN: CreateSpaceAnalysisTaskRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: CreateSpaceAnalysisTaskRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c CreateSpaceAnalysisTaskRequestXLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateSpaceAnalysisTaskRequestXLanguage) UnmarshalJSON(b []byte) error {
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
