package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 行数对比任务表级详情。
type TableLineCompareDetailInfo struct {

	// 源库表名称。
	SourceTableName *string `json:"source_table_name,omitempty"`

	// 源库表行数。
	SourceRowNum *string `json:"source_row_num,omitempty"`

	// 目标库表名称。
	TargetTableName *string `json:"target_table_name,omitempty"`

	// 目标库表行数。
	TargetRowNum *string `json:"target_row_num,omitempty"`

	// 差异值。
	DifferenceRowNum *string `json:"difference_row_num,omitempty"`

	// 对比结果。取值： - CONSISTENT：一致。 - INCONSISTENT：不一致。 - COMPARING：正在对比。 - WAITING_FOR_COMPARISON：等待对比。 - FAILED_TO_COMPARE：对比失败。 - TARGET_DB_NOT_EXIST：目标库不存在。 - CAN_NOT_COMPARE：无法对比。
	Status *TableLineCompareDetailInfoStatus `json:"status,omitempty"`

	// 信息。
	Message *string `json:"message,omitempty"`
}

func (o TableLineCompareDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableLineCompareDetailInfo struct{}"
	}

	return strings.Join([]string{"TableLineCompareDetailInfo", string(data)}, " ")
}

type TableLineCompareDetailInfoStatus struct {
	value string
}

type TableLineCompareDetailInfoStatusEnum struct {
	CONSISTENT             TableLineCompareDetailInfoStatus
	INCONSISTENT           TableLineCompareDetailInfoStatus
	COMPARING              TableLineCompareDetailInfoStatus
	WAITING_FOR_COMPARISON TableLineCompareDetailInfoStatus
	FAILED_TO_COMPARE      TableLineCompareDetailInfoStatus
	TARGET_DB_NOT_EXIST    TableLineCompareDetailInfoStatus
	CAN_NOT_COMPARE        TableLineCompareDetailInfoStatus
}

func GetTableLineCompareDetailInfoStatusEnum() TableLineCompareDetailInfoStatusEnum {
	return TableLineCompareDetailInfoStatusEnum{
		CONSISTENT: TableLineCompareDetailInfoStatus{
			value: "CONSISTENT",
		},
		INCONSISTENT: TableLineCompareDetailInfoStatus{
			value: "INCONSISTENT",
		},
		COMPARING: TableLineCompareDetailInfoStatus{
			value: "COMPARING",
		},
		WAITING_FOR_COMPARISON: TableLineCompareDetailInfoStatus{
			value: "WAITING_FOR_COMPARISON",
		},
		FAILED_TO_COMPARE: TableLineCompareDetailInfoStatus{
			value: "FAILED_TO_COMPARE",
		},
		TARGET_DB_NOT_EXIST: TableLineCompareDetailInfoStatus{
			value: "TARGET_DB_NOT_EXIST",
		},
		CAN_NOT_COMPARE: TableLineCompareDetailInfoStatus{
			value: "CAN_NOT_COMPARE",
		},
	}
}

func (c TableLineCompareDetailInfoStatus) Value() string {
	return c.value
}

func (c TableLineCompareDetailInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TableLineCompareDetailInfoStatus) UnmarshalJSON(b []byte) error {
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
