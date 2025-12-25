package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowPipeConsumptionResponse Response Object
type ShowPipeConsumptionResponse struct {

	// 接入点域名信息(格式：{dataspace}.{endpoint})
	AccessPoint *string `json:"access_point,omitempty"`

	// 数据管道名称
	PipeName *string `json:"pipe_name,omitempty"`

	// 实时消费的开关状态；enable-开启，disable-关闭
	Status *ShowPipeConsumptionResponseStatus `json:"status,omitempty"`

	// 订阅名称
	SubscriptionName *string `json:"subscription_name,omitempty"`

	// 表id
	TableId        *string `json:"table_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPipeConsumptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPipeConsumptionResponse struct{}"
	}

	return strings.Join([]string{"ShowPipeConsumptionResponse", string(data)}, " ")
}

type ShowPipeConsumptionResponseStatus struct {
	value string
}

type ShowPipeConsumptionResponseStatusEnum struct {
	ENABLE  ShowPipeConsumptionResponseStatus
	DISABLE ShowPipeConsumptionResponseStatus
}

func GetShowPipeConsumptionResponseStatusEnum() ShowPipeConsumptionResponseStatusEnum {
	return ShowPipeConsumptionResponseStatusEnum{
		ENABLE: ShowPipeConsumptionResponseStatus{
			value: "enable",
		},
		DISABLE: ShowPipeConsumptionResponseStatus{
			value: "disable",
		},
	}
}

func (c ShowPipeConsumptionResponseStatus) Value() string {
	return c.value
}

func (c ShowPipeConsumptionResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowPipeConsumptionResponseStatus) UnmarshalJSON(b []byte) error {
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
