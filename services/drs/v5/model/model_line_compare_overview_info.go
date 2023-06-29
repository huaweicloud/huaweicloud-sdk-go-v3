package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LineCompareOverviewInfo 行数对比概览信息体。
type LineCompareOverviewInfo struct {

	// 源库库名。
	SourceDbName *string `json:"source_db_name,omitempty"`

	// 目标库库名。
	TargetDbName *string `json:"target_db_name,omitempty"`

	// 行对比结果。取值： - CONSISTENT：一致。 - INCONSISTENT：不一致。 - COMPARING：正在对比。 - WAITING_FOR_COMPARISON：等待对比。 - FAILED_TO_COMPARE：对比失败。 - TARGET_DB_NOT_EXIST：目标库不存在。 - CAN_NOT_COMPARE：无法对比。
	Status *LineCompareOverviewInfoStatus `json:"status,omitempty"`
}

func (o LineCompareOverviewInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LineCompareOverviewInfo struct{}"
	}

	return strings.Join([]string{"LineCompareOverviewInfo", string(data)}, " ")
}

type LineCompareOverviewInfoStatus struct {
	value string
}

type LineCompareOverviewInfoStatusEnum struct {
	CONSISTENT             LineCompareOverviewInfoStatus
	INCONSISTENT           LineCompareOverviewInfoStatus
	COMPARING              LineCompareOverviewInfoStatus
	WAITING_FOR_COMPARISON LineCompareOverviewInfoStatus
	FAILED_TO_COMPARE      LineCompareOverviewInfoStatus
	TARGET_DB_NOT_EXIST    LineCompareOverviewInfoStatus
	CAN_NOT_COMPARE        LineCompareOverviewInfoStatus
}

func GetLineCompareOverviewInfoStatusEnum() LineCompareOverviewInfoStatusEnum {
	return LineCompareOverviewInfoStatusEnum{
		CONSISTENT: LineCompareOverviewInfoStatus{
			value: "CONSISTENT",
		},
		INCONSISTENT: LineCompareOverviewInfoStatus{
			value: "INCONSISTENT",
		},
		COMPARING: LineCompareOverviewInfoStatus{
			value: "COMPARING",
		},
		WAITING_FOR_COMPARISON: LineCompareOverviewInfoStatus{
			value: "WAITING_FOR_COMPARISON",
		},
		FAILED_TO_COMPARE: LineCompareOverviewInfoStatus{
			value: "FAILED_TO_COMPARE",
		},
		TARGET_DB_NOT_EXIST: LineCompareOverviewInfoStatus{
			value: "TARGET_DB_NOT_EXIST",
		},
		CAN_NOT_COMPARE: LineCompareOverviewInfoStatus{
			value: "CAN_NOT_COMPARE",
		},
	}
}

func (c LineCompareOverviewInfoStatus) Value() string {
	return c.value
}

func (c LineCompareOverviewInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LineCompareOverviewInfoStatus) UnmarshalJSON(b []byte) error {
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
