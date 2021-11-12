package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 诊断项
type DiagnosisItem struct {
	// 诊断项名称

	Name DiagnosisItemName `json:"name"`
	// 原因ID列表

	CauseIds *[]ConclusionItem `json:"cause_ids,omitempty"`
	// 影响ID列表

	ImpactIds *[]ConclusionItem `json:"impact_ids,omitempty"`
	// 建议ID列表

	AdviceIds *[]ConclusionItem `json:"advice_ids,omitempty"`
	// 诊断结果

	Result DiagnosisItemResult `json:"result"`
}

func (o DiagnosisItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiagnosisItem struct{}"
	}

	return strings.Join([]string{"DiagnosisItem", string(data)}, " ")
}

type DiagnosisItemName struct {
	value string
}

type DiagnosisItemNameEnum struct {
	CONNECTION_NUM             DiagnosisItemName
	RX_CONTROLLED              DiagnosisItemName
	PERSISTENCE                DiagnosisItemName
	CENTRALIZED_EXPIRATION     DiagnosisItemName
	INNER_MEMORY_FRAGMENTATION DiagnosisItemName
	TIME_CONSUMING_COMMANDS    DiagnosisItemName
	HIT_RATIO                  DiagnosisItemName
	MEMORY_USAGE               DiagnosisItemName
	CPU_USAGE                  DiagnosisItemName
}

func GetDiagnosisItemNameEnum() DiagnosisItemNameEnum {
	return DiagnosisItemNameEnum{
		CONNECTION_NUM: DiagnosisItemName{
			value: "connection_num",
		},
		RX_CONTROLLED: DiagnosisItemName{
			value: "rx_controlled",
		},
		PERSISTENCE: DiagnosisItemName{
			value: "persistence",
		},
		CENTRALIZED_EXPIRATION: DiagnosisItemName{
			value: "centralized_expiration",
		},
		INNER_MEMORY_FRAGMENTATION: DiagnosisItemName{
			value: "inner_memory_fragmentation",
		},
		TIME_CONSUMING_COMMANDS: DiagnosisItemName{
			value: "time_consuming_commands",
		},
		HIT_RATIO: DiagnosisItemName{
			value: "hit_ratio",
		},
		MEMORY_USAGE: DiagnosisItemName{
			value: "memory_usage",
		},
		CPU_USAGE: DiagnosisItemName{
			value: "cpu_usage",
		},
	}
}

func (c DiagnosisItemName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DiagnosisItemName) UnmarshalJSON(b []byte) error {
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

type DiagnosisItemResult struct {
	value string
}

type DiagnosisItemResultEnum struct {
	FAILED   DiagnosisItemResult
	ABNORMAL DiagnosisItemResult
	NORMAL   DiagnosisItemResult
}

func GetDiagnosisItemResultEnum() DiagnosisItemResultEnum {
	return DiagnosisItemResultEnum{
		FAILED: DiagnosisItemResult{
			value: "failed",
		},
		ABNORMAL: DiagnosisItemResult{
			value: "abnormal",
		},
		NORMAL: DiagnosisItemResult{
			value: "normal",
		},
	}
}

func (c DiagnosisItemResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DiagnosisItemResult) UnmarshalJSON(b []byte) error {
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
