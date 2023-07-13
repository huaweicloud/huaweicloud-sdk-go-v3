package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CustomizeSourceCreateReq struct {

	// 自定义事件源名称，租户下唯一，由小写字母、数字、点、下划线和中划线组成，必须以字母或数字开头，且不能以hc.开头
	Name string `json:"name"`

	// 事件源描述
	Description *string `json:"description,omitempty"`

	// 指导事件源归属的事件通道ID
	ChannelId string `json:"channel_id"`

	// 事件源类型
	Type *CustomizeSourceCreateReqType `json:"type,omitempty"`

	// json格式封装消息实例链接信息：如RabbitMQ实例的instance_id字段、虚拟主机vhost字段、队列queue字段、用户名、密码等
	Detail *interface{} `json:"detail,omitempty"`
}

func (o CustomizeSourceCreateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomizeSourceCreateReq struct{}"
	}

	return strings.Join([]string{"CustomizeSourceCreateReq", string(data)}, " ")
}

type CustomizeSourceCreateReqType struct {
	value string
}

type CustomizeSourceCreateReqTypeEnum struct {
	APPLICATION CustomizeSourceCreateReqType
	RABBITMQ    CustomizeSourceCreateReqType
	ROCKETMQ    CustomizeSourceCreateReqType
}

func GetCustomizeSourceCreateReqTypeEnum() CustomizeSourceCreateReqTypeEnum {
	return CustomizeSourceCreateReqTypeEnum{
		APPLICATION: CustomizeSourceCreateReqType{
			value: "APPLICATION",
		},
		RABBITMQ: CustomizeSourceCreateReqType{
			value: "RABBITMQ",
		},
		ROCKETMQ: CustomizeSourceCreateReqType{
			value: "ROCKETMQ",
		},
	}
}

func (c CustomizeSourceCreateReqType) Value() string {
	return c.value
}

func (c CustomizeSourceCreateReqType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CustomizeSourceCreateReqType) UnmarshalJSON(b []byte) error {
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
