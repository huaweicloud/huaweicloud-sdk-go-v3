package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowApisOverviewRequest Request Object
type ShowApisOverviewRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	// dlm版本类型
	DlmType ShowApisOverviewRequestDlmType `json:"Dlm-Type"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	// 开始时间（13位时间戳）
	StartTime int64 `json:"start_time"`

	// 结束时间（13位时间戳）
	EndTime int64 `json:"end_time"`

	// 时间单位
	TimeUnit ShowApisOverviewRequestTimeUnit `json:"time_unit"`
}

func (o ShowApisOverviewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApisOverviewRequest struct{}"
	}

	return strings.Join([]string{"ShowApisOverviewRequest", string(data)}, " ")
}

type ShowApisOverviewRequestDlmType struct {
	value string
}

type ShowApisOverviewRequestDlmTypeEnum struct {
	SHARED    ShowApisOverviewRequestDlmType
	EXCLUSIVE ShowApisOverviewRequestDlmType
}

func GetShowApisOverviewRequestDlmTypeEnum() ShowApisOverviewRequestDlmTypeEnum {
	return ShowApisOverviewRequestDlmTypeEnum{
		SHARED: ShowApisOverviewRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ShowApisOverviewRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ShowApisOverviewRequestDlmType) Value() string {
	return c.value
}

func (c ShowApisOverviewRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApisOverviewRequestDlmType) UnmarshalJSON(b []byte) error {
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

type ShowApisOverviewRequestTimeUnit struct {
	value string
}

type ShowApisOverviewRequestTimeUnitEnum struct {
	HOUR ShowApisOverviewRequestTimeUnit
	DAY  ShowApisOverviewRequestTimeUnit
}

func GetShowApisOverviewRequestTimeUnitEnum() ShowApisOverviewRequestTimeUnitEnum {
	return ShowApisOverviewRequestTimeUnitEnum{
		HOUR: ShowApisOverviewRequestTimeUnit{
			value: "HOUR",
		},
		DAY: ShowApisOverviewRequestTimeUnit{
			value: "DAY",
		},
	}
}

func (c ShowApisOverviewRequestTimeUnit) Value() string {
	return c.value
}

func (c ShowApisOverviewRequestTimeUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApisOverviewRequestTimeUnit) UnmarshalJSON(b []byte) error {
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
