package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAppsDetailRequest Request Object
type ShowAppsDetailRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// dlm版本类型
	DlmType ShowAppsDetailRequestDlmType `json:"Dlm-Type"`

	// app编号
	AppId string `json:"app_id"`

	// 开始时间（13位时间戳）
	StartTime int64 `json:"start_time"`

	// 结束时间（13位时间戳）
	EndTime int64 `json:"end_time"`

	// 时间单位
	TimeUnit ShowAppsDetailRequestTimeUnit `json:"time_unit"`
}

func (o ShowAppsDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppsDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowAppsDetailRequest", string(data)}, " ")
}

type ShowAppsDetailRequestDlmType struct {
	value string
}

type ShowAppsDetailRequestDlmTypeEnum struct {
	SHARED    ShowAppsDetailRequestDlmType
	EXCLUSIVE ShowAppsDetailRequestDlmType
}

func GetShowAppsDetailRequestDlmTypeEnum() ShowAppsDetailRequestDlmTypeEnum {
	return ShowAppsDetailRequestDlmTypeEnum{
		SHARED: ShowAppsDetailRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ShowAppsDetailRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ShowAppsDetailRequestDlmType) Value() string {
	return c.value
}

func (c ShowAppsDetailRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAppsDetailRequestDlmType) UnmarshalJSON(b []byte) error {
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

type ShowAppsDetailRequestTimeUnit struct {
	value string
}

type ShowAppsDetailRequestTimeUnitEnum struct {
	HOUR ShowAppsDetailRequestTimeUnit
	DAY  ShowAppsDetailRequestTimeUnit
}

func GetShowAppsDetailRequestTimeUnitEnum() ShowAppsDetailRequestTimeUnitEnum {
	return ShowAppsDetailRequestTimeUnitEnum{
		HOUR: ShowAppsDetailRequestTimeUnit{
			value: "HOUR",
		},
		DAY: ShowAppsDetailRequestTimeUnit{
			value: "DAY",
		},
	}
}

func (c ShowAppsDetailRequestTimeUnit) Value() string {
	return c.value
}

func (c ShowAppsDetailRequestTimeUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAppsDetailRequestTimeUnit) UnmarshalJSON(b []byte) error {
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
