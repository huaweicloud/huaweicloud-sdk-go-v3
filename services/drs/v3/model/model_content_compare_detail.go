package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 内容对比结果详情。
type ContentCompareDetail struct {

	// 源库名称。
	SourceDbName string `json:"source_db_name" xml:"source_db_name"`

	// 目标库名称。
	TargetDbName string `json:"target_db_name" xml:"target_db_name"`

	// 源库的表名称。
	SourceTableName string `json:"source_table_name" xml:"source_table_name"`

	// 目标库的表名称。
	TargetTableName string `json:"target_table_name" xml:"target_table_name"`

	// 源库的表的行数。
	SourceRowNum int32 `json:"source_row_num" xml:"source_row_num"`

	// 目标库的表的行数。
	TargetRowNum int32 `json:"target_row_num" xml:"target_row_num"`

	// 源库的表和目标库的表的差异值。
	DiffRowNum int32 `json:"diff_row_num" xml:"diff_row_num"`

	// 行对比结果。
	LineCompareResult *ContentCompareDetailLineCompareResult `json:"line_compare_result,omitempty" xml:"line_compare_result"`

	// 内容对比结果。
	ContentCompareResult ContentCompareDetailContentCompareResult `json:"content_compare_result" xml:"content_compare_result"`

	// 附加信息。
	Message *string `json:"message,omitempty" xml:"message"`
}

func (o ContentCompareDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContentCompareDetail struct{}"
	}

	return strings.Join([]string{"ContentCompareDetail", string(data)}, " ")
}

type ContentCompareDetailLineCompareResult struct {
	value string
}

type ContentCompareDetailLineCompareResultEnum struct {
	CONSISTENT             ContentCompareDetailLineCompareResult
	INCONSISTENT           ContentCompareDetailLineCompareResult
	COMPARING              ContentCompareDetailLineCompareResult
	WAITING_FOR_COMPARISON ContentCompareDetailLineCompareResult
	FAILED_TO_COMPARE      ContentCompareDetailLineCompareResult
	TARGET_DB_NOT_EXIT     ContentCompareDetailLineCompareResult
	CAN_NOT_COMPARE        ContentCompareDetailLineCompareResult
}

func GetContentCompareDetailLineCompareResultEnum() ContentCompareDetailLineCompareResultEnum {
	return ContentCompareDetailLineCompareResultEnum{
		CONSISTENT: ContentCompareDetailLineCompareResult{
			value: "CONSISTENT-一致",
		},
		INCONSISTENT: ContentCompareDetailLineCompareResult{
			value: "INCONSISTENT-不一致",
		},
		COMPARING: ContentCompareDetailLineCompareResult{
			value: "COMPARING-正在对比",
		},
		WAITING_FOR_COMPARISON: ContentCompareDetailLineCompareResult{
			value: "WAITING_FOR_COMPARISON-等待对比",
		},
		FAILED_TO_COMPARE: ContentCompareDetailLineCompareResult{
			value: "FAILED_TO_COMPARE-对比失败",
		},
		TARGET_DB_NOT_EXIT: ContentCompareDetailLineCompareResult{
			value: "TARGET_DB_NOT_EXIT-目标库不存在",
		},
		CAN_NOT_COMPARE: ContentCompareDetailLineCompareResult{
			value: "CAN_NOT_COMPARE-无法对比",
		},
	}
}

func (c ContentCompareDetailLineCompareResult) Value() string {
	return c.value
}

func (c ContentCompareDetailLineCompareResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ContentCompareDetailLineCompareResult) UnmarshalJSON(b []byte) error {
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

type ContentCompareDetailContentCompareResult struct {
	value string
}

type ContentCompareDetailContentCompareResultEnum struct {
	CONSISTENT             ContentCompareDetailContentCompareResult
	INCONSISTENT           ContentCompareDetailContentCompareResult
	COMPARING              ContentCompareDetailContentCompareResult
	WAITING_FOR_COMPARISON ContentCompareDetailContentCompareResult
	FAILED_TO_COMPARE      ContentCompareDetailContentCompareResult
	TARGET_DB_NOT_EXIT     ContentCompareDetailContentCompareResult
	CAN_NOT_COMPARE        ContentCompareDetailContentCompareResult
}

func GetContentCompareDetailContentCompareResultEnum() ContentCompareDetailContentCompareResultEnum {
	return ContentCompareDetailContentCompareResultEnum{
		CONSISTENT: ContentCompareDetailContentCompareResult{
			value: "CONSISTENT-一致",
		},
		INCONSISTENT: ContentCompareDetailContentCompareResult{
			value: "INCONSISTENT-不一致",
		},
		COMPARING: ContentCompareDetailContentCompareResult{
			value: "COMPARING-正在对比",
		},
		WAITING_FOR_COMPARISON: ContentCompareDetailContentCompareResult{
			value: "WAITING_FOR_COMPARISON-等待对比",
		},
		FAILED_TO_COMPARE: ContentCompareDetailContentCompareResult{
			value: "FAILED_TO_COMPARE-对比失败",
		},
		TARGET_DB_NOT_EXIT: ContentCompareDetailContentCompareResult{
			value: "TARGET_DB_NOT_EXIT-目标库不存在",
		},
		CAN_NOT_COMPARE: ContentCompareDetailContentCompareResult{
			value: "CAN_NOT_COMPARE-无法对比",
		},
	}
}

func (c ContentCompareDetailContentCompareResult) Value() string {
	return c.value
}

func (c ContentCompareDetailContentCompareResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ContentCompareDetailContentCompareResult) UnmarshalJSON(b []byte) error {
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
