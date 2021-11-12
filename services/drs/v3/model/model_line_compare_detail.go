package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

//
type LineCompareDetail struct {
	// 源库的表名称。

	SourceTableName string `json:"source_table_name"`
	// 目标库的表名称。

	TargetTableName string `json:"target_table_name"`
	// 源库的表行数。

	SourceRowNum int32 `json:"source_row_num"`
	// 目标库的表行数。

	TargetRowNum int32 `json:"target_row_num"`
	// 源库的表和目标库的表的差异值。

	DiffRowNum int32 `json:"diff_row_num"`
	// 对比结果。

	LineCompareResult LineCompareDetailLineCompareResult `json:"line_compare_result"`
	// 附加信息。

	Message *string `json:"message,omitempty"`
}

func (o LineCompareDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LineCompareDetail struct{}"
	}

	return strings.Join([]string{"LineCompareDetail", string(data)}, " ")
}

type LineCompareDetailLineCompareResult struct {
	value string
}

type LineCompareDetailLineCompareResultEnum struct {
	CONSISTENT             LineCompareDetailLineCompareResult
	INCONSISTENT           LineCompareDetailLineCompareResult
	COMPARING              LineCompareDetailLineCompareResult
	WAITING_FOR_COMPARISON LineCompareDetailLineCompareResult
	FAILED_TO_COMPARE      LineCompareDetailLineCompareResult
	TARGET_DB_NOT_EXIT     LineCompareDetailLineCompareResult
	CAN_NOT_COMPARE        LineCompareDetailLineCompareResult
}

func GetLineCompareDetailLineCompareResultEnum() LineCompareDetailLineCompareResultEnum {
	return LineCompareDetailLineCompareResultEnum{
		CONSISTENT: LineCompareDetailLineCompareResult{
			value: "CONSISTENT-一致",
		},
		INCONSISTENT: LineCompareDetailLineCompareResult{
			value: "INCONSISTENT-不一致",
		},
		COMPARING: LineCompareDetailLineCompareResult{
			value: "COMPARING-正在对比",
		},
		WAITING_FOR_COMPARISON: LineCompareDetailLineCompareResult{
			value: "WAITING_FOR_COMPARISON-等待对比",
		},
		FAILED_TO_COMPARE: LineCompareDetailLineCompareResult{
			value: "FAILED_TO_COMPARE-对比失败",
		},
		TARGET_DB_NOT_EXIT: LineCompareDetailLineCompareResult{
			value: "TARGET_DB_NOT_EXIT-目标库不存在",
		},
		CAN_NOT_COMPARE: LineCompareDetailLineCompareResult{
			value: "CAN_NOT_COMPARE-无法对比",
		},
	}
}

func (c LineCompareDetailLineCompareResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LineCompareDetailLineCompareResult) UnmarshalJSON(b []byte) error {
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
