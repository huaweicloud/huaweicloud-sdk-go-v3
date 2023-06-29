package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ContentCompareResultOverview struct {

	// 源库名称。
	SourceDbName string `json:"source_db_name"`

	// 目标库名称。
	TargetDbName string `json:"target_db_name"`

	// 对比结果。 - CONSISTENT-一致 - INCONSISTENT-不一致 - COMPARING-正在对比 - WAITING_FOR_COMPARISON-等待对比 - FAILED_TO_COMPARE-对比失败 - TARGET_DB_NOT_EXIT-目标库不存在 - CAN_NOT_COMPARE-无法对比
	ContentCompareResult ContentCompareResultOverviewContentCompareResult `json:"content_compare_result"`
}

func (o ContentCompareResultOverview) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContentCompareResultOverview struct{}"
	}

	return strings.Join([]string{"ContentCompareResultOverview", string(data)}, " ")
}

type ContentCompareResultOverviewContentCompareResult struct {
	value string
}

type ContentCompareResultOverviewContentCompareResultEnum struct {
	CONSISTENT             ContentCompareResultOverviewContentCompareResult
	INCONSISTENT           ContentCompareResultOverviewContentCompareResult
	COMPARING              ContentCompareResultOverviewContentCompareResult
	WAITING_FOR_COMPARISON ContentCompareResultOverviewContentCompareResult
	FAILED_TO_COMPARE      ContentCompareResultOverviewContentCompareResult
	TARGET_DB_NOT_EXIT     ContentCompareResultOverviewContentCompareResult
	CAN_NOT_COMPARE        ContentCompareResultOverviewContentCompareResult
}

func GetContentCompareResultOverviewContentCompareResultEnum() ContentCompareResultOverviewContentCompareResultEnum {
	return ContentCompareResultOverviewContentCompareResultEnum{
		CONSISTENT: ContentCompareResultOverviewContentCompareResult{
			value: "CONSISTENT",
		},
		INCONSISTENT: ContentCompareResultOverviewContentCompareResult{
			value: "INCONSISTENT",
		},
		COMPARING: ContentCompareResultOverviewContentCompareResult{
			value: "COMPARING",
		},
		WAITING_FOR_COMPARISON: ContentCompareResultOverviewContentCompareResult{
			value: "WAITING_FOR_COMPARISON",
		},
		FAILED_TO_COMPARE: ContentCompareResultOverviewContentCompareResult{
			value: "FAILED_TO_COMPARE",
		},
		TARGET_DB_NOT_EXIT: ContentCompareResultOverviewContentCompareResult{
			value: "TARGET_DB_NOT_EXIT",
		},
		CAN_NOT_COMPARE: ContentCompareResultOverviewContentCompareResult{
			value: "CAN_NOT_COMPARE",
		},
	}
}

func (c ContentCompareResultOverviewContentCompareResult) Value() string {
	return c.value
}

func (c ContentCompareResultOverviewContentCompareResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ContentCompareResultOverviewContentCompareResult) UnmarshalJSON(b []byte) error {
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
