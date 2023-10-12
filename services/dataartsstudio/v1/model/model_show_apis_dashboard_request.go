package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowApisDashboardRequest Request Object
type ShowApisDashboardRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	// dlm版本类型
	DlmType ShowApisDashboardRequestDlmType `json:"Dlm-Type"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	// 集群编号
	InstanceId *string `json:"instance_id,omitempty"`

	// 开始时间（13位时间戳）
	StartTime int64 `json:"start_time"`

	// 结束时间（13位时间戳）
	EndTime int64 `json:"end_time"`

	// 时间单位
	TimeUnit ShowApisDashboardRequestTimeUnit `json:"time_unit"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowApisDashboardRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApisDashboardRequest struct{}"
	}

	return strings.Join([]string{"ShowApisDashboardRequest", string(data)}, " ")
}

type ShowApisDashboardRequestDlmType struct {
	value string
}

type ShowApisDashboardRequestDlmTypeEnum struct {
	SHARED    ShowApisDashboardRequestDlmType
	EXCLUSIVE ShowApisDashboardRequestDlmType
}

func GetShowApisDashboardRequestDlmTypeEnum() ShowApisDashboardRequestDlmTypeEnum {
	return ShowApisDashboardRequestDlmTypeEnum{
		SHARED: ShowApisDashboardRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ShowApisDashboardRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ShowApisDashboardRequestDlmType) Value() string {
	return c.value
}

func (c ShowApisDashboardRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApisDashboardRequestDlmType) UnmarshalJSON(b []byte) error {
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

type ShowApisDashboardRequestTimeUnit struct {
	value string
}

type ShowApisDashboardRequestTimeUnitEnum struct {
	HOUR ShowApisDashboardRequestTimeUnit
	DAY  ShowApisDashboardRequestTimeUnit
}

func GetShowApisDashboardRequestTimeUnitEnum() ShowApisDashboardRequestTimeUnitEnum {
	return ShowApisDashboardRequestTimeUnitEnum{
		HOUR: ShowApisDashboardRequestTimeUnit{
			value: "HOUR",
		},
		DAY: ShowApisDashboardRequestTimeUnit{
			value: "DAY",
		},
	}
}

func (c ShowApisDashboardRequestTimeUnit) Value() string {
	return c.value
}

func (c ShowApisDashboardRequestTimeUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApisDashboardRequestTimeUnit) UnmarshalJSON(b []byte) error {
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
