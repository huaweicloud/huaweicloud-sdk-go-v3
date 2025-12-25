package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EnableDataConsumptionResponse Response Object
type EnableDataConsumptionResponse struct {

	// 表名称
	TableName *string `json:"table_name,omitempty"`

	// 数据空间ID
	DataspaceId *string `json:"dataspace_id,omitempty"`

	// 管道ID
	PipeId *string `json:"pipe_id,omitempty"`

	// 管道名称
	PipeName *string `json:"pipe_name,omitempty"`

	// **参数解释**: 状态 - ENABLED 启用 - DISABLED 禁用  **约束限制** 不涉及 **取值范围**: - ENABLED - DISABLED  **默认值** 不涉及
	Status *EnableDataConsumptionResponseStatus `json:"status,omitempty"`

	// **参数解释**: 网络类型 - INTERNET 互联网 - INTRANET 内联网  **约束限制** 不涉及 **取值范围**: - INTERNET - INTRANET  **默认值** 不涉及
	Type *EnableDataConsumptionResponseType `json:"type,omitempty"`

	// 接入点域名信息(格式：{dataspace}.{endpoint})
	AccessPoint *string `json:"access_point,omitempty"`

	// 订阅名称
	Subscription *string `json:"subscription,omitempty"`

	// 表Id
	TableId        *string `json:"table_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o EnableDataConsumptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableDataConsumptionResponse struct{}"
	}

	return strings.Join([]string{"EnableDataConsumptionResponse", string(data)}, " ")
}

type EnableDataConsumptionResponseStatus struct {
	value string
}

type EnableDataConsumptionResponseStatusEnum struct {
	ENABLED  EnableDataConsumptionResponseStatus
	DISABLED EnableDataConsumptionResponseStatus
}

func GetEnableDataConsumptionResponseStatusEnum() EnableDataConsumptionResponseStatusEnum {
	return EnableDataConsumptionResponseStatusEnum{
		ENABLED: EnableDataConsumptionResponseStatus{
			value: "ENABLED",
		},
		DISABLED: EnableDataConsumptionResponseStatus{
			value: "DISABLED",
		},
	}
}

func (c EnableDataConsumptionResponseStatus) Value() string {
	return c.value
}

func (c EnableDataConsumptionResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EnableDataConsumptionResponseStatus) UnmarshalJSON(b []byte) error {
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

type EnableDataConsumptionResponseType struct {
	value string
}

type EnableDataConsumptionResponseTypeEnum struct {
	INTERNET EnableDataConsumptionResponseType
	INTRANET EnableDataConsumptionResponseType
}

func GetEnableDataConsumptionResponseTypeEnum() EnableDataConsumptionResponseTypeEnum {
	return EnableDataConsumptionResponseTypeEnum{
		INTERNET: EnableDataConsumptionResponseType{
			value: "INTERNET",
		},
		INTRANET: EnableDataConsumptionResponseType{
			value: "INTRANET",
		},
	}
}

func (c EnableDataConsumptionResponseType) Value() string {
	return c.value
}

func (c EnableDataConsumptionResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EnableDataConsumptionResponseType) UnmarshalJSON(b []byte) error {
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
