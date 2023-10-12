package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowApiDashboardRequest Request Object
type ShowApiDashboardRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	// dlm版本类型
	DlmType ShowApiDashboardRequestDlmType `json:"Dlm-Type"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	// api编号
	ApiId string `json:"api_id"`

	// 集群编号
	InstanceId *string `json:"instance_id,omitempty"`

	// 开始时间（13位时间戳）
	StartTime int64 `json:"start_time"`

	// 结束时间（13位时间戳）
	EndTime int64 `json:"end_time"`

	// 时间单位
	TimeUnit ShowApiDashboardRequestTimeUnit `json:"time_unit"`
}

func (o ShowApiDashboardRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApiDashboardRequest struct{}"
	}

	return strings.Join([]string{"ShowApiDashboardRequest", string(data)}, " ")
}

type ShowApiDashboardRequestDlmType struct {
	value string
}

type ShowApiDashboardRequestDlmTypeEnum struct {
	SHARED    ShowApiDashboardRequestDlmType
	EXCLUSIVE ShowApiDashboardRequestDlmType
}

func GetShowApiDashboardRequestDlmTypeEnum() ShowApiDashboardRequestDlmTypeEnum {
	return ShowApiDashboardRequestDlmTypeEnum{
		SHARED: ShowApiDashboardRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ShowApiDashboardRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ShowApiDashboardRequestDlmType) Value() string {
	return c.value
}

func (c ShowApiDashboardRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApiDashboardRequestDlmType) UnmarshalJSON(b []byte) error {
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

type ShowApiDashboardRequestTimeUnit struct {
	value string
}

type ShowApiDashboardRequestTimeUnitEnum struct {
	HOUR ShowApiDashboardRequestTimeUnit
	DAY  ShowApiDashboardRequestTimeUnit
}

func GetShowApiDashboardRequestTimeUnitEnum() ShowApiDashboardRequestTimeUnitEnum {
	return ShowApiDashboardRequestTimeUnitEnum{
		HOUR: ShowApiDashboardRequestTimeUnit{
			value: "HOUR",
		},
		DAY: ShowApiDashboardRequestTimeUnit{
			value: "DAY",
		},
	}
}

func (c ShowApiDashboardRequestTimeUnit) Value() string {
	return c.value
}

func (c ShowApiDashboardRequestTimeUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApiDashboardRequestTimeUnit) UnmarshalJSON(b []byte) error {
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
