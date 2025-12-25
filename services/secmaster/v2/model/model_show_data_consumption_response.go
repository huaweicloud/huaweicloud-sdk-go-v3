package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDataConsumptionResponse Response Object
type ShowDataConsumptionResponse struct {

	// 表名称
	TableName *string `json:"table_name,omitempty"`

	// 数据空间ID
	DataspaceId *string `json:"dataspace_id,omitempty"`

	// 管道ID
	PipeId *string `json:"pipe_id,omitempty"`

	// 管道名称
	PipeName *string `json:"pipe_name,omitempty"`

	// **参数解释**: 状态 - ENABLED 启用 - DISABLED 禁用  **约束限制** 不涉及 **取值范围**: - ENABLED - DISABLED  **默认值** 不涉及
	Status *ShowDataConsumptionResponseStatus `json:"status,omitempty"`

	// **参数解释**: 网络类型 - INTERNET 互联网 - INTRANET 内联网  **约束限制** 不涉及 **取值范围**: - INTERNET - INTRANET  **默认值** 不涉及
	Type *ShowDataConsumptionResponseType `json:"type,omitempty"`

	// 接入点域名信息(格式：{dataspace}.{endpoint})
	AccessPoint *string `json:"access_point,omitempty"`

	// 订阅名称
	Subscription *string `json:"subscription,omitempty"`

	// 表Id
	TableId        *string `json:"table_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDataConsumptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataConsumptionResponse struct{}"
	}

	return strings.Join([]string{"ShowDataConsumptionResponse", string(data)}, " ")
}

type ShowDataConsumptionResponseStatus struct {
	value string
}

type ShowDataConsumptionResponseStatusEnum struct {
	ENABLED  ShowDataConsumptionResponseStatus
	DISABLED ShowDataConsumptionResponseStatus
}

func GetShowDataConsumptionResponseStatusEnum() ShowDataConsumptionResponseStatusEnum {
	return ShowDataConsumptionResponseStatusEnum{
		ENABLED: ShowDataConsumptionResponseStatus{
			value: "ENABLED",
		},
		DISABLED: ShowDataConsumptionResponseStatus{
			value: "DISABLED",
		},
	}
}

func (c ShowDataConsumptionResponseStatus) Value() string {
	return c.value
}

func (c ShowDataConsumptionResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDataConsumptionResponseStatus) UnmarshalJSON(b []byte) error {
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

type ShowDataConsumptionResponseType struct {
	value string
}

type ShowDataConsumptionResponseTypeEnum struct {
	INTERNET ShowDataConsumptionResponseType
	INTRANET ShowDataConsumptionResponseType
}

func GetShowDataConsumptionResponseTypeEnum() ShowDataConsumptionResponseTypeEnum {
	return ShowDataConsumptionResponseTypeEnum{
		INTERNET: ShowDataConsumptionResponseType{
			value: "INTERNET",
		},
		INTRANET: ShowDataConsumptionResponseType{
			value: "INTRANET",
		},
	}
}

func (c ShowDataConsumptionResponseType) Value() string {
	return c.value
}

func (c ShowDataConsumptionResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDataConsumptionResponseType) UnmarshalJSON(b []byte) error {
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
