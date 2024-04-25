package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowApisDetailRequest Request Object
type ShowApisDetailRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 数据服务的版本类型，指定SHARED共享版或EXCLUSIVE专享版。
	DlmType *ShowApisDetailRequestDlmType `json:"Dlm-Type,omitempty"`

	// 消息体的类型（格式），有Body体的情况下必选，没有Body体无需填写。如果请求消息体中含有中文字符，则需要通过charset=utf8指定中文字符集，例如取值为：application/json;charset=utf8。
	ContentType string `json:"Content-Type"`

	// api编号。
	ApiId string `json:"api_id"`

	// 集群编号。
	InstanceId *string `json:"instance_id,omitempty"`

	// 开始时间（13位时间戳）。
	StartTime int64 `json:"start_time"`

	// 结束时间（13位时间戳）。
	EndTime int64 `json:"end_time"`

	// 时间单位。
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
