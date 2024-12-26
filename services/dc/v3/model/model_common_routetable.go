package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CommonRoutetable 路由明细
type CommonRoutetable struct {

	// 路由id
	Id string `json:"id"`

	// 租户id
	TenantId string `json:"tenant_id"`

	// 网关id
	GatewayId *string `json:"gateway_id,omitempty"`

	// 路由子网
	Destination string `json:"destination"`

	// 下一跳id
	Nexthop string `json:"nexthop"`

	// 路由类型: - customized: 默认路由 - specific: 自定义路由 - bgp: 动态路由
	ObtainMode CommonRoutetableObtainMode `json:"obtain_mode"`

	// 路由状态: - ACTIVE: 下发正常 - ERROR: 下发失败 - PENDING_CREATE: 待下发
	Status CommonRoutetableStatus `json:"status"`

	AddressFamily *AddressFamily `json:"address_family"`

	// 路由描述
	Description *string `json:"description,omitempty"`

	// 下一跳类型: - vif_peer: 虚拟接口对等体 - gdgw: 全域接入网关
	Type CommonRoutetableType `json:"type"`
}

func (o CommonRoutetable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonRoutetable struct{}"
	}

	return strings.Join([]string{"CommonRoutetable", string(data)}, " ")
}

type CommonRoutetableObtainMode struct {
	value string
}

type CommonRoutetableObtainModeEnum struct {
	CUSTOMIZED CommonRoutetableObtainMode
	SPECIFIC   CommonRoutetableObtainMode
	BGP        CommonRoutetableObtainMode
}

func GetCommonRoutetableObtainModeEnum() CommonRoutetableObtainModeEnum {
	return CommonRoutetableObtainModeEnum{
		CUSTOMIZED: CommonRoutetableObtainMode{
			value: "customized",
		},
		SPECIFIC: CommonRoutetableObtainMode{
			value: "specific",
		},
		BGP: CommonRoutetableObtainMode{
			value: "bgp",
		},
	}
}

func (c CommonRoutetableObtainMode) Value() string {
	return c.value
}

func (c CommonRoutetableObtainMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CommonRoutetableObtainMode) UnmarshalJSON(b []byte) error {
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

type CommonRoutetableStatus struct {
	value string
}

type CommonRoutetableStatusEnum struct {
	ACTIVE         CommonRoutetableStatus
	ERROR          CommonRoutetableStatus
	PENDING_CREATE CommonRoutetableStatus
}

func GetCommonRoutetableStatusEnum() CommonRoutetableStatusEnum {
	return CommonRoutetableStatusEnum{
		ACTIVE: CommonRoutetableStatus{
			value: "ACTIVE",
		},
		ERROR: CommonRoutetableStatus{
			value: "ERROR",
		},
		PENDING_CREATE: CommonRoutetableStatus{
			value: "PENDING_CREATE",
		},
	}
}

func (c CommonRoutetableStatus) Value() string {
	return c.value
}

func (c CommonRoutetableStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CommonRoutetableStatus) UnmarshalJSON(b []byte) error {
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

type CommonRoutetableType struct {
	value string
}

type CommonRoutetableTypeEnum struct {
	VIF_PEER CommonRoutetableType
	GDGW     CommonRoutetableType
}

func GetCommonRoutetableTypeEnum() CommonRoutetableTypeEnum {
	return CommonRoutetableTypeEnum{
		VIF_PEER: CommonRoutetableType{
			value: "vif_peer",
		},
		GDGW: CommonRoutetableType{
			value: "gdgw",
		},
	}
}

func (c CommonRoutetableType) Value() string {
	return c.value
}

func (c CommonRoutetableType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CommonRoutetableType) UnmarshalJSON(b []byte) error {
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
