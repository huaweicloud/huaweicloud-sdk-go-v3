package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ContentCompareOverviewInfo 内容对比概览信息体。
type ContentCompareOverviewInfo struct {

	// 源库库名。
	SourceDb *string `json:"source_db,omitempty"`

	// 目标库库名。
	TargetDb *string `json:"target_db,omitempty"`

	// 对比结果。取值： - CONSISTENT：一致。 - INCONSISTENT：不一致。 - COMPARING：正在对比。 - WAITING_FOR_COMPARISON：等待对比。 - FAILED_TO_COMPARE：对比失败。 - TARGET_DB_NOT_EXIST：目标库不存在。 - CAN_NOT_COMPARE：无法对比。 - WAIT_FOR_COMPARE：命令已下发，等待对比结果。 - CANCELED：已取消。
	Status *ContentCompareOverviewInfoStatus `json:"status,omitempty"`

	// 对比结果。
	CompareResult *bool `json:"compare_result,omitempty"`
}

func (o ContentCompareOverviewInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContentCompareOverviewInfo struct{}"
	}

	return strings.Join([]string{"ContentCompareOverviewInfo", string(data)}, " ")
}

type ContentCompareOverviewInfoStatus struct {
	value string
}

type ContentCompareOverviewInfoStatusEnum struct {
	CONSISTENT             ContentCompareOverviewInfoStatus
	INCONSISTENT           ContentCompareOverviewInfoStatus
	COMPARING              ContentCompareOverviewInfoStatus
	WAITING_FOR_COMPARISON ContentCompareOverviewInfoStatus
	FAILED_TO_COMPARE      ContentCompareOverviewInfoStatus
	TARGET_DB_NOT_EXIST    ContentCompareOverviewInfoStatus
	CAN_NOT_COMPARE        ContentCompareOverviewInfoStatus
	WAIT_FOR_COMPARE       ContentCompareOverviewInfoStatus
	CANCELED               ContentCompareOverviewInfoStatus
}

func GetContentCompareOverviewInfoStatusEnum() ContentCompareOverviewInfoStatusEnum {
	return ContentCompareOverviewInfoStatusEnum{
		CONSISTENT: ContentCompareOverviewInfoStatus{
			value: "CONSISTENT",
		},
		INCONSISTENT: ContentCompareOverviewInfoStatus{
			value: "INCONSISTENT",
		},
		COMPARING: ContentCompareOverviewInfoStatus{
			value: "COMPARING",
		},
		WAITING_FOR_COMPARISON: ContentCompareOverviewInfoStatus{
			value: "WAITING_FOR_COMPARISON",
		},
		FAILED_TO_COMPARE: ContentCompareOverviewInfoStatus{
			value: "FAILED_TO_COMPARE",
		},
		TARGET_DB_NOT_EXIST: ContentCompareOverviewInfoStatus{
			value: "TARGET_DB_NOT_EXIST",
		},
		CAN_NOT_COMPARE: ContentCompareOverviewInfoStatus{
			value: "CAN_NOT_COMPARE",
		},
		WAIT_FOR_COMPARE: ContentCompareOverviewInfoStatus{
			value: "WAIT_FOR_COMPARE",
		},
		CANCELED: ContentCompareOverviewInfoStatus{
			value: "CANCELED",
		},
	}
}

func (c ContentCompareOverviewInfoStatus) Value() string {
	return c.value
}

func (c ContentCompareOverviewInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ContentCompareOverviewInfoStatus) UnmarshalJSON(b []byte) error {
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
