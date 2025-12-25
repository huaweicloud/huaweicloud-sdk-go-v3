package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DisableDataConsumptionResponse Response Object
type DisableDataConsumptionResponse struct {

	// 实时消费的开关状态；enable-开启，disable-关闭
	Status *DisableDataConsumptionResponseStatus `json:"status,omitempty"`

	// 接入点域名信息(格式：{dataspace}.{endpoint})
	AccessPoint *string `json:"access_point,omitempty"`

	// 数据管道名称
	PipeName *string `json:"pipe_name,omitempty"`

	// 订阅名称
	SubscriptionName *string `json:"subscription_name,omitempty"`

	// table表id
	TableId        *string `json:"table_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DisableDataConsumptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableDataConsumptionResponse struct{}"
	}

	return strings.Join([]string{"DisableDataConsumptionResponse", string(data)}, " ")
}

type DisableDataConsumptionResponseStatus struct {
	value string
}

type DisableDataConsumptionResponseStatusEnum struct {
	ENABLE  DisableDataConsumptionResponseStatus
	DISABLE DisableDataConsumptionResponseStatus
}

func GetDisableDataConsumptionResponseStatusEnum() DisableDataConsumptionResponseStatusEnum {
	return DisableDataConsumptionResponseStatusEnum{
		ENABLE: DisableDataConsumptionResponseStatus{
			value: "enable",
		},
		DISABLE: DisableDataConsumptionResponseStatus{
			value: "disable",
		},
	}
}

func (c DisableDataConsumptionResponseStatus) Value() string {
	return c.value
}

func (c DisableDataConsumptionResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DisableDataConsumptionResponseStatus) UnmarshalJSON(b []byte) error {
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
