package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateLogStreamParams 创建日志流参数。
type CreateLogStreamParams struct {

	// 需要创建的日志流名称。
	LogStreamName CreateLogStreamParamsLogStreamName `json:"log_stream_name"`

	// 企业项目名称。
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	// 日志存储时间  最小值：1 最大值：365 说明： 该参数仅对华东-上海一、华北-北京四、华南-广州用户开放。
	TtlInDays *CreateLogStreamParamsTtlInDays `json:"ttl_in_days,omitempty"`

	// 标签字段信息
	Tags *[]TagsBody `json:"tags,omitempty"`
}

func (o CreateLogStreamParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogStreamParams struct{}"
	}

	return strings.Join([]string{"CreateLogStreamParams", string(data)}, " ")
}

type CreateLogStreamParamsLogStreamName struct {
	value string
}

type CreateLogStreamParamsLogStreamNameEnum struct {
	LTS_STREAM_13CI CreateLogStreamParamsLogStreamName
}

func GetCreateLogStreamParamsLogStreamNameEnum() CreateLogStreamParamsLogStreamNameEnum {
	return CreateLogStreamParamsLogStreamNameEnum{
		LTS_STREAM_13CI: CreateLogStreamParamsLogStreamName{
			value: "lts-stream-13ci",
		},
	}
}

func (c CreateLogStreamParamsLogStreamName) Value() string {
	return c.value
}

func (c CreateLogStreamParamsLogStreamName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLogStreamParamsLogStreamName) UnmarshalJSON(b []byte) error {
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

type CreateLogStreamParamsTtlInDays struct {
	value string
}

type CreateLogStreamParamsTtlInDaysEnum struct {
	E_7 CreateLogStreamParamsTtlInDays
}

func GetCreateLogStreamParamsTtlInDaysEnum() CreateLogStreamParamsTtlInDaysEnum {
	return CreateLogStreamParamsTtlInDaysEnum{
		E_7: CreateLogStreamParamsTtlInDays{
			value: "7",
		},
	}
}

func (c CreateLogStreamParamsTtlInDays) Value() string {
	return c.value
}

func (c CreateLogStreamParamsTtlInDays) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLogStreamParamsTtlInDays) UnmarshalJSON(b []byte) error {
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
