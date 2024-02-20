package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// AnalyzerSummary 包含有关分析器的信息。
type AnalyzerSummary struct {

	// 分析器创建的时间。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 分析器的唯一标识符。
	Id string `json:"id"`

	// 唯一的资源名称。
	LastAnalyzedResource *string `json:"last_analyzed_resource,omitempty"`

	// 分析最近分析的资源的时间。
	LastResourceAnalyzedAt *sdktime.SdkTime `json:"last_resource_analyzed_at,omitempty"`

	// 分析器的名称。
	Name string `json:"name"`

	// 分析器的状态
	Status AnalyzerSummaryStatus `json:"status"`

	StatusReason *StatusReason `json:"status_reason,omitempty"`

	Tags *[]Tag `json:"tags,omitempty"`

	Type *AnalyzerType `json:"type"`

	// 唯一的资源名称。
	Urn string `json:"urn"`
}

func (o AnalyzerSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnalyzerSummary struct{}"
	}

	return strings.Join([]string{"AnalyzerSummary", string(data)}, " ")
}

type AnalyzerSummaryStatus struct {
	value string
}

type AnalyzerSummaryStatusEnum struct {
	ACTIVE   AnalyzerSummaryStatus
	CREATING AnalyzerSummaryStatus
	FAILED   AnalyzerSummaryStatus
}

func GetAnalyzerSummaryStatusEnum() AnalyzerSummaryStatusEnum {
	return AnalyzerSummaryStatusEnum{
		ACTIVE: AnalyzerSummaryStatus{
			value: "active",
		},
		CREATING: AnalyzerSummaryStatus{
			value: "creating",
		},
		FAILED: AnalyzerSummaryStatus{
			value: "failed",
		},
	}
}

func (c AnalyzerSummaryStatus) Value() string {
	return c.value
}

func (c AnalyzerSummaryStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AnalyzerSummaryStatus) UnmarshalJSON(b []byte) error {
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
