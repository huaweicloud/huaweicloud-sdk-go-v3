package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSpaceAnalysisTaskRequest Request Object
type CreateSpaceAnalysisTaskRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 语言
	XLanguage *CreateSpaceAnalysisTaskRequestXLanguage `json:"X-Language,omitempty"`

	Body *CreateSpaceAnalysisTaskBody `json:"body,omitempty"`
}

func (o CreateSpaceAnalysisTaskRequest) String() string {
	data, err := utils.Marshal(o)
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

func (c CreateSpaceAnalysisTaskRequestXLanguage) Value() string {
	return c.value
}

func (c CreateSpaceAnalysisTaskRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSpaceAnalysisTaskRequestXLanguage) UnmarshalJSON(b []byte) error {
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
