package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 物理专线更新参数
type UpdateDirectConnect struct {

	// 物理专线的名字
	Name *string `json:"name,omitempty"`

	// 物理专线的描述信息
	Description *string `json:"description,omitempty"`

	// 指定托管专线接入带宽,单位Mbps
	Bandwidth *int32 `json:"bandwidth,omitempty"`

	// 物理专线对端所在的物理位置，省/市/街道或IDC名字
	PeerLocation *string `json:"peer_location,omitempty"`

	// 更新资源状态，合法值是：PENDING_PAY
	Status *UpdateDirectConnectStatus `json:"status,omitempty"`

	// 更新运营商状态，合法值是：ACTIVE,DOWN
	ProviderStatus *UpdateDirectConnectProviderStatus `json:"provider_status,omitempty"`
}

func (o UpdateDirectConnect) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDirectConnect struct{}"
	}

	return strings.Join([]string{"UpdateDirectConnect", string(data)}, " ")
}

type UpdateDirectConnectStatus struct {
	value string
}

type UpdateDirectConnectStatusEnum struct {
	PENDING_PAY UpdateDirectConnectStatus
}

func GetUpdateDirectConnectStatusEnum() UpdateDirectConnectStatusEnum {
	return UpdateDirectConnectStatusEnum{
		PENDING_PAY: UpdateDirectConnectStatus{
			value: "PENDING_PAY",
		},
	}
}

func (c UpdateDirectConnectStatus) Value() string {
	return c.value
}

func (c UpdateDirectConnectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDirectConnectStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type UpdateDirectConnectProviderStatus struct {
	value string
}

type UpdateDirectConnectProviderStatusEnum struct {
	ACTIVE UpdateDirectConnectProviderStatus
	DOWN   UpdateDirectConnectProviderStatus
}

func GetUpdateDirectConnectProviderStatusEnum() UpdateDirectConnectProviderStatusEnum {
	return UpdateDirectConnectProviderStatusEnum{
		ACTIVE: UpdateDirectConnectProviderStatus{
			value: "ACTIVE",
		},
		DOWN: UpdateDirectConnectProviderStatus{
			value: "DOWN",
		},
	}
}

func (c UpdateDirectConnectProviderStatus) Value() string {
	return c.value
}

func (c UpdateDirectConnectProviderStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDirectConnectProviderStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
