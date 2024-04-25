package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowApisDashboardRequest Request Object
type ShowApisDashboardRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 数据服务的版本类型，指定SHARED共享版或EXCLUSIVE专享版。
	DlmType *ShowApisDashboardRequestDlmType `json:"Dlm-Type,omitempty"`

	// 消息体的类型（格式），有Body体的情况下必选，没有Body体无需填写。如果请求消息体中含有中文字符，则需要通过charset=utf8指定中文字符集，例如取值为：application/json;charset=utf8。
	ContentType string `json:"Content-Type"`

	// 集群编号。
	InstanceId *string `json:"instance_id,omitempty"`

	// 开始时间（13位时间戳）。
	StartTime int64 `json:"start_time"`

	// 结束时间（13位时间戳）。
	EndTime int64 `json:"end_time"`

	// 时间单位。
	TimeUnit ShowApisDashboardRequestTimeUnit `json:"time_unit"`

	// limit。
	Limit *int32 `json:"limit,omitempty"`

	// offset。
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
