package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

//
type ContentCompareResultOverview struct {
	// 源库名称。

	SourceDbName string `json:"source_db_name"`
	// 目标库名称。

	TargetDbName string `json:"target_db_name"`
	// 对比结果。

	ContentCompareResult ContentCompareResultOverviewContentCompareResult `json:"content_compare_result"`
}

func (o ContentCompareResultOverview) String() string {
	data, err := json.Marshal(o)
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
			value: "CONSISTENT-一致",
		},
		INCONSISTENT: ContentCompareResultOverviewContentCompareResult{
			value: "INCONSISTENT-不一致",
		},
		COMPARING: ContentCompareResultOverviewContentCompareResult{
			value: "COMPARING-正在对比",
		},
		WAITING_FOR_COMPARISON: ContentCompareResultOverviewContentCompareResult{
			value: "WAITING_FOR_COMPARISON-等待对比",
		},
		FAILED_TO_COMPARE: ContentCompareResultOverviewContentCompareResult{
			value: "FAILED_TO_COMPARE-对比失败",
		},
		TARGET_DB_NOT_EXIT: ContentCompareResultOverviewContentCompareResult{
			value: "TARGET_DB_NOT_EXIT-目标库不存在",
		},
		CAN_NOT_COMPARE: ContentCompareResultOverviewContentCompareResult{
			value: "CAN_NOT_COMPARE-无法对比",
		},
	}
}

func (c ContentCompareResultOverviewContentCompareResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ContentCompareResultOverviewContentCompareResult) UnmarshalJSON(b []byte) error {
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
