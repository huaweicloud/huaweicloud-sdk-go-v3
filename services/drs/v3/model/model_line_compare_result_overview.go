package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type LineCompareResultOverview struct {

	// 源库名称。
	SourceDbName string `json:"source_db_name"`

	// 目标库名称。
	TargetDbName string `json:"target_db_name"`

	// 对比结果。 - CONSISTENT-一致 - INCONSISTENT-不一致 - COMPARING-正在对比 - WAITING_FOR_COMPARISON-等待对比 - FAILED_TO_COMPARE-对比失败 - TARGET_DB_NOT_EXIT-目标库不存在 - CAN_NOT_COMPARE-无法对比
	LineCompareResult LineCompareResultOverviewLineCompareResult `json:"line_compare_result"`
}

func (o LineCompareResultOverview) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LineCompareResultOverview struct{}"
	}

	return strings.Join([]string{"LineCompareResultOverview", string(data)}, " ")
}

type LineCompareResultOverviewLineCompareResult struct {
	value string
}

type LineCompareResultOverviewLineCompareResultEnum struct {
	CONSISTENT             LineCompareResultOverviewLineCompareResult
	INCONSISTENT           LineCompareResultOverviewLineCompareResult
	COMPARING              LineCompareResultOverviewLineCompareResult
	WAITING_FOR_COMPARISON LineCompareResultOverviewLineCompareResult
	FAILED_TO_COMPARE      LineCompareResultOverviewLineCompareResult
	TARGET_DB_NOT_EXIT     LineCompareResultOverviewLineCompareResult
	CAN_NOT_COMPARE        LineCompareResultOverviewLineCompareResult
}

func GetLineCompareResultOverviewLineCompareResultEnum() LineCompareResultOverviewLineCompareResultEnum {
	return LineCompareResultOverviewLineCompareResultEnum{
		CONSISTENT: LineCompareResultOverviewLineCompareResult{
			value: "CONSISTENT",
		},
		INCONSISTENT: LineCompareResultOverviewLineCompareResult{
			value: "INCONSISTENT",
		},
		COMPARING: LineCompareResultOverviewLineCompareResult{
			value: "COMPARING",
		},
		WAITING_FOR_COMPARISON: LineCompareResultOverviewLineCompareResult{
			value: "WAITING_FOR_COMPARISON",
		},
		FAILED_TO_COMPARE: LineCompareResultOverviewLineCompareResult{
			value: "FAILED_TO_COMPARE",
		},
		TARGET_DB_NOT_EXIT: LineCompareResultOverviewLineCompareResult{
			value: "TARGET_DB_NOT_EXIT",
		},
		CAN_NOT_COMPARE: LineCompareResultOverviewLineCompareResult{
			value: "CAN_NOT_COMPARE",
		},
	}
}

func (c LineCompareResultOverviewLineCompareResult) Value() string {
	return c.value
}

func (c LineCompareResultOverviewLineCompareResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LineCompareResultOverviewLineCompareResult) UnmarshalJSON(b []byte) error {
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
