package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ChatResourceConfigInfo 资源配置。
type ChatResourceConfigInfo struct {

	// 资源id
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源数量
	CountResource *int32 `json:"count_resource,omitempty"`

	// 资源来源。 * PURCHASED: 购买 * SP_ALLOCATED：SP分配 * ADMIN_ALLOCATED：系统管理员分配 > * 开通按需；购买按需套餐包、一次性资源、包周期等都属于PURCHASED。
	ResourceSource *ChatResourceConfigInfoResourceSource `json:"resource_source,omitempty"`

	// 资源计费类型。 * PERIODIC: 包周期 * ONE_TIME：一次性计费 * ON_DEMAND: 按需计费 > * 一次性计费包括：租户订购的一次性资源，SP管理员分配给租户的一次性资源，SP管理员分配给租户的按需套餐包资源，系统管理员分配的资源（分钟数等）。
	ChargeMode *ChatResourceConfigInfoChargeMode `json:"charge_mode,omitempty"`

	// 资源数量。智能交互并发路数。
	ResourceNums *int32 `json:"resource_nums,omitempty"`

	// 可用资源数量。智能交互并发路数。
	ResourceAvailableNums *int32 `json:"resource_available_nums,omitempty"`

	// 资源状态。 * ACTIVE: 正常 * FREEZE：冻结 * DELETED：退订或过期
	Status *ChatResourceConfigInfoStatus `json:"status,omitempty"`

	// 资源过期时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"
	ExpireTime *string `json:"expire_time,omitempty"`
}

func (o ChatResourceConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChatResourceConfigInfo struct{}"
	}

	return strings.Join([]string{"ChatResourceConfigInfo", string(data)}, " ")
}

type ChatResourceConfigInfoResourceSource struct {
	value string
}

type ChatResourceConfigInfoResourceSourceEnum struct {
	PURCHASED       ChatResourceConfigInfoResourceSource
	SP_ALLOCATED    ChatResourceConfigInfoResourceSource
	ADMIN_ALLOCATED ChatResourceConfigInfoResourceSource
}

func GetChatResourceConfigInfoResourceSourceEnum() ChatResourceConfigInfoResourceSourceEnum {
	return ChatResourceConfigInfoResourceSourceEnum{
		PURCHASED: ChatResourceConfigInfoResourceSource{
			value: "PURCHASED",
		},
		SP_ALLOCATED: ChatResourceConfigInfoResourceSource{
			value: "SP_ALLOCATED",
		},
		ADMIN_ALLOCATED: ChatResourceConfigInfoResourceSource{
			value: "ADMIN_ALLOCATED",
		},
	}
}

func (c ChatResourceConfigInfoResourceSource) Value() string {
	return c.value
}

func (c ChatResourceConfigInfoResourceSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChatResourceConfigInfoResourceSource) UnmarshalJSON(b []byte) error {
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

type ChatResourceConfigInfoChargeMode struct {
	value string
}

type ChatResourceConfigInfoChargeModeEnum struct {
	PERIODIC  ChatResourceConfigInfoChargeMode
	ONE_TIME  ChatResourceConfigInfoChargeMode
	ON_DEMAND ChatResourceConfigInfoChargeMode
}

func GetChatResourceConfigInfoChargeModeEnum() ChatResourceConfigInfoChargeModeEnum {
	return ChatResourceConfigInfoChargeModeEnum{
		PERIODIC: ChatResourceConfigInfoChargeMode{
			value: "PERIODIC",
		},
		ONE_TIME: ChatResourceConfigInfoChargeMode{
			value: "ONE_TIME",
		},
		ON_DEMAND: ChatResourceConfigInfoChargeMode{
			value: "ON_DEMAND",
		},
	}
}

func (c ChatResourceConfigInfoChargeMode) Value() string {
	return c.value
}

func (c ChatResourceConfigInfoChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChatResourceConfigInfoChargeMode) UnmarshalJSON(b []byte) error {
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

type ChatResourceConfigInfoStatus struct {
	value string
}

type ChatResourceConfigInfoStatusEnum struct {
	ACTIVE  ChatResourceConfigInfoStatus
	FREEZE  ChatResourceConfigInfoStatus
	DELETED ChatResourceConfigInfoStatus
}

func GetChatResourceConfigInfoStatusEnum() ChatResourceConfigInfoStatusEnum {
	return ChatResourceConfigInfoStatusEnum{
		ACTIVE: ChatResourceConfigInfoStatus{
			value: "ACTIVE",
		},
		FREEZE: ChatResourceConfigInfoStatus{
			value: "FREEZE",
		},
		DELETED: ChatResourceConfigInfoStatus{
			value: "DELETED",
		},
	}
}

func (c ChatResourceConfigInfoStatus) Value() string {
	return c.value
}

func (c ChatResourceConfigInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChatResourceConfigInfoStatus) UnmarshalJSON(b []byte) error {
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
