package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAppsDashboardRequest Request Object
type ShowAppsDashboardRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// dlm版本类型
	DlmType ShowAppsDashboardRequestDlmType `json:"Dlm-Type"`

	// 开始时间（13位时间戳）
	StartTime int64 `json:"start_time"`

	// 结束时间（13位时间戳）
	EndTime int64 `json:"end_time"`

	// 时间单位
	TimeUnit ShowAppsDashboardRequestTimeUnit `json:"time_unit"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowAppsDashboardRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppsDashboardRequest struct{}"
	}

	return strings.Join([]string{"ShowAppsDashboardRequest", string(data)}, " ")
}

type ShowAppsDashboardRequestDlmType struct {
	value string
}

type ShowAppsDashboardRequestDlmTypeEnum struct {
	SHARED    ShowAppsDashboardRequestDlmType
	EXCLUSIVE ShowAppsDashboardRequestDlmType
}

func GetShowAppsDashboardRequestDlmTypeEnum() ShowAppsDashboardRequestDlmTypeEnum {
	return ShowAppsDashboardRequestDlmTypeEnum{
		SHARED: ShowAppsDashboardRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ShowAppsDashboardRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ShowAppsDashboardRequestDlmType) Value() string {
	return c.value
}

func (c ShowAppsDashboardRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAppsDashboardRequestDlmType) UnmarshalJSON(b []byte) error {
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

type ShowAppsDashboardRequestTimeUnit struct {
	value string
}

type ShowAppsDashboardRequestTimeUnitEnum struct {
	HOUR ShowAppsDashboardRequestTimeUnit
	DAY  ShowAppsDashboardRequestTimeUnit
}

func GetShowAppsDashboardRequestTimeUnitEnum() ShowAppsDashboardRequestTimeUnitEnum {
	return ShowAppsDashboardRequestTimeUnitEnum{
		HOUR: ShowAppsDashboardRequestTimeUnit{
			value: "HOUR",
		},
		DAY: ShowAppsDashboardRequestTimeUnit{
			value: "DAY",
		},
	}
}

func (c ShowAppsDashboardRequestTimeUnit) Value() string {
	return c.value
}

func (c ShowAppsDashboardRequestTimeUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAppsDashboardRequestTimeUnit) UnmarshalJSON(b []byte) error {
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
