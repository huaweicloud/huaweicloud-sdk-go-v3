package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowApisDetailRequest Request Object
type ShowApisDetailRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// dlm版本类型
	DlmType ShowApisDetailRequestDlmType `json:"Dlm-Type"`

	// api编号
	ApiId string `json:"api_id"`

	// 集群编号
	InstanceId *string `json:"instance_id,omitempty"`

	// 开始时间（13位时间戳）
	StartTime int64 `json:"start_time"`

	// 结束时间（13位时间戳）
	EndTime int64 `json:"end_time"`

	// 时间单位
	TimeUnit ShowApisDetailRequestTimeUnit `json:"time_unit"`
}

func (o ShowApisDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApisDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowApisDetailRequest", string(data)}, " ")
}

type ShowApisDetailRequestDlmType struct {
	value string
}

type ShowApisDetailRequestDlmTypeEnum struct {
	SHARED    ShowApisDetailRequestDlmType
	EXCLUSIVE ShowApisDetailRequestDlmType
}

func GetShowApisDetailRequestDlmTypeEnum() ShowApisDetailRequestDlmTypeEnum {
	return ShowApisDetailRequestDlmTypeEnum{
		SHARED: ShowApisDetailRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ShowApisDetailRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ShowApisDetailRequestDlmType) Value() string {
	return c.value
}

func (c ShowApisDetailRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApisDetailRequestDlmType) UnmarshalJSON(b []byte) error {
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

type ShowApisDetailRequestTimeUnit struct {
	value string
}

type ShowApisDetailRequestTimeUnitEnum struct {
	HOUR ShowApisDetailRequestTimeUnit
	DAY  ShowApisDetailRequestTimeUnit
}

func GetShowApisDetailRequestTimeUnitEnum() ShowApisDetailRequestTimeUnitEnum {
	return ShowApisDetailRequestTimeUnitEnum{
		HOUR: ShowApisDetailRequestTimeUnit{
			value: "HOUR",
		},
		DAY: ShowApisDetailRequestTimeUnit{
			value: "DAY",
		},
	}
}

func (c ShowApisDetailRequestTimeUnit) Value() string {
	return c.value
}

func (c ShowApisDetailRequestTimeUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApisDetailRequestTimeUnit) UnmarshalJSON(b []byte) error {
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
