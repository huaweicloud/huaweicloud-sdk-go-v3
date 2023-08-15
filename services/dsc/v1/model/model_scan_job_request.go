package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ScanJobRequest struct {

	// 资产ID列表
	AssetIds *[]string `json:"asset_ids,omitempty"`

	// 扫描周期，日(DAY)，周(WEEK)，月(MONTH)，单次扫描(ONCE)
	Cycle *ScanJobRequestCycle `json:"cycle,omitempty"`

	// 扫描任务名
	Name *string `json:"name,omitempty"`

	// 是否开启任务
	Open *bool `json:"open,omitempty"`

	// 规则组ID列表
	RuleGroupIds *[]string `json:"rule_group_ids,omitempty"`

	// 扫描任务开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 时区
	TimeZone *string `json:"time_zone,omitempty"`

	// 主题的唯一资源标识符
	TopicUrn *string `json:"topic_urn,omitempty"`

	// 是否用nlp
	UseNlp *bool `json:"use_nlp,omitempty"`
}

func (o ScanJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScanJobRequest struct{}"
	}

	return strings.Join([]string{"ScanJobRequest", string(data)}, " ")
}

type ScanJobRequestCycle struct {
	value string
}

type ScanJobRequestCycleEnum struct {
	ONCE  ScanJobRequestCycle
	DAY   ScanJobRequestCycle
	WEEK  ScanJobRequestCycle
	MONTH ScanJobRequestCycle
}

func GetScanJobRequestCycleEnum() ScanJobRequestCycleEnum {
	return ScanJobRequestCycleEnum{
		ONCE: ScanJobRequestCycle{
			value: "ONCE",
		},
		DAY: ScanJobRequestCycle{
			value: "DAY",
		},
		WEEK: ScanJobRequestCycle{
			value: "WEEK",
		},
		MONTH: ScanJobRequestCycle{
			value: "MONTH",
		},
	}
}

func (c ScanJobRequestCycle) Value() string {
	return c.value
}

func (c ScanJobRequestCycle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScanJobRequestCycle) UnmarshalJSON(b []byte) error {
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
