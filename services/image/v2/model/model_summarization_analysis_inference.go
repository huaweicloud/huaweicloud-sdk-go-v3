package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SummarizationAnalysisInference struct {

	// 摘要片段输出总时长
	OutputDuration int32 `json:"output_duration"`

	// 输出类型，不填摘要片段和集锦都输出。填summary只输出集锦；填fragment只输出片段。
	OutcomeType *SummarizationAnalysisInferenceOutcomeType `json:"outcome_type,omitempty"`
}

func (o SummarizationAnalysisInference) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SummarizationAnalysisInference struct{}"
	}

	return strings.Join([]string{"SummarizationAnalysisInference", string(data)}, " ")
}

type SummarizationAnalysisInferenceOutcomeType struct {
	value string
}

type SummarizationAnalysisInferenceOutcomeTypeEnum struct {
	FRAGMENT SummarizationAnalysisInferenceOutcomeType
	SUMMARY  SummarizationAnalysisInferenceOutcomeType
}

func GetSummarizationAnalysisInferenceOutcomeTypeEnum() SummarizationAnalysisInferenceOutcomeTypeEnum {
	return SummarizationAnalysisInferenceOutcomeTypeEnum{
		FRAGMENT: SummarizationAnalysisInferenceOutcomeType{
			value: "fragment",
		},
		SUMMARY: SummarizationAnalysisInferenceOutcomeType{
			value: "summary",
		},
	}
}

func (c SummarizationAnalysisInferenceOutcomeType) Value() string {
	return c.value
}

func (c SummarizationAnalysisInferenceOutcomeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SummarizationAnalysisInferenceOutcomeType) UnmarshalJSON(b []byte) error {
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
