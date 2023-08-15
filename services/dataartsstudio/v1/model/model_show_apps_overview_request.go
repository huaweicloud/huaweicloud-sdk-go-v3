package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAppsOverviewRequest Request Object
type ShowAppsOverviewRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// dlm版本类型
	DlmType ShowAppsOverviewRequestDlmType `json:"Dlm-Type"`

	// 开始时间（13位时间戳）
	StartTime int64 `json:"start_time"`

	// 结束时间（13位时间戳）
	EndTime int64 `json:"end_time"`

	// 时间单位
	TimeUnit ShowAppsOverviewRequestTimeUnit `json:"time_unit"`
}

func (o ShowAppsOverviewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppsOverviewRequest struct{}"
	}

	return strings.Join([]string{"ShowAppsOverviewRequest", string(data)}, " ")
}

type ShowAppsOverviewRequestDlmType struct {
	value string
}

type ShowAppsOverviewRequestDlmTypeEnum struct {
	SHARED    ShowAppsOverviewRequestDlmType
	EXCLUSIVE ShowAppsOverviewRequestDlmType
}

func GetShowAppsOverviewRequestDlmTypeEnum() ShowAppsOverviewRequestDlmTypeEnum {
	return ShowAppsOverviewRequestDlmTypeEnum{
		SHARED: ShowAppsOverviewRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ShowAppsOverviewRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ShowAppsOverviewRequestDlmType) Value() string {
	return c.value
}

func (c ShowAppsOverviewRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAppsOverviewRequestDlmType) UnmarshalJSON(b []byte) error {
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

type ShowAppsOverviewRequestTimeUnit struct {
	value string
}

type ShowAppsOverviewRequestTimeUnitEnum struct {
	HOUR ShowAppsOverviewRequestTimeUnit
	DAY  ShowAppsOverviewRequestTimeUnit
}

func GetShowAppsOverviewRequestTimeUnitEnum() ShowAppsOverviewRequestTimeUnitEnum {
	return ShowAppsOverviewRequestTimeUnitEnum{
		HOUR: ShowAppsOverviewRequestTimeUnit{
			value: "HOUR",
		},
		DAY: ShowAppsOverviewRequestTimeUnit{
			value: "DAY",
		},
	}
}

func (c ShowAppsOverviewRequestTimeUnit) Value() string {
	return c.value
}

func (c ShowAppsOverviewRequestTimeUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAppsOverviewRequestTimeUnit) UnmarshalJSON(b []byte) error {
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
