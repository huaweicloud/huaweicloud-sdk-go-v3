package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeletePipeConsumptionResponse Response Object
type DeletePipeConsumptionResponse struct {

	// 接入点域名信息(格式：{dataspace}.{endpoint})
	AccessPoint *string `json:"access_point,omitempty"`

	// 数据管道名称
	PipeName *string `json:"pipe_name,omitempty"`

	// 实时消费的开关状态；enable-开启，disable-关闭
	Status *DeletePipeConsumptionResponseStatus `json:"status,omitempty"`

	// 订阅名称
	SubscriptionName *string `json:"subscription_name,omitempty"`

	// 表id
	TableId        *string `json:"table_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeletePipeConsumptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePipeConsumptionResponse struct{}"
	}

	return strings.Join([]string{"DeletePipeConsumptionResponse", string(data)}, " ")
}

type DeletePipeConsumptionResponseStatus struct {
	value string
}

type DeletePipeConsumptionResponseStatusEnum struct {
	ENABLE  DeletePipeConsumptionResponseStatus
	DISABLE DeletePipeConsumptionResponseStatus
}

func GetDeletePipeConsumptionResponseStatusEnum() DeletePipeConsumptionResponseStatusEnum {
	return DeletePipeConsumptionResponseStatusEnum{
		ENABLE: DeletePipeConsumptionResponseStatus{
			value: "enable",
		},
		DISABLE: DeletePipeConsumptionResponseStatus{
			value: "disable",
		},
	}
}

func (c DeletePipeConsumptionResponseStatus) Value() string {
	return c.value
}

func (c DeletePipeConsumptionResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeletePipeConsumptionResponseStatus) UnmarshalJSON(b []byte) error {
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
