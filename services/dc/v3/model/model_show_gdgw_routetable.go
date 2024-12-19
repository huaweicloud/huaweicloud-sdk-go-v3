package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowGdgwRoutetable 路由明细
type ShowGdgwRoutetable struct {

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
	ObtainMode ShowGdgwRoutetableObtainMode `json:"obtain_mode"`

	// 路由状态: - ACTIVE: 下发正常 - ERROR: 下发失败 - PENDING_CREATE: 待下发
	Status ShowGdgwRoutetableStatus `json:"status"`

	AddressFamily *AddressFamily `json:"address_family"`

	// 路由描述
	Description *string `json:"description,omitempty"`

	Type *RouteTypeOfGdgw `json:"type"`
}

func (o ShowGdgwRoutetable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGdgwRoutetable struct{}"
	}

	return strings.Join([]string{"ShowGdgwRoutetable", string(data)}, " ")
}

type ShowGdgwRoutetableObtainMode struct {
	value string
}

type ShowGdgwRoutetableObtainModeEnum struct {
	CUSTOMIZED ShowGdgwRoutetableObtainMode
	SPECIFIC   ShowGdgwRoutetableObtainMode
	BGP        ShowGdgwRoutetableObtainMode
}

func GetShowGdgwRoutetableObtainModeEnum() ShowGdgwRoutetableObtainModeEnum {
	return ShowGdgwRoutetableObtainModeEnum{
		CUSTOMIZED: ShowGdgwRoutetableObtainMode{
			value: "customized",
		},
		SPECIFIC: ShowGdgwRoutetableObtainMode{
			value: "specific",
		},
		BGP: ShowGdgwRoutetableObtainMode{
			value: "bgp",
		},
	}
}

func (c ShowGdgwRoutetableObtainMode) Value() string {
	return c.value
}

func (c ShowGdgwRoutetableObtainMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowGdgwRoutetableObtainMode) UnmarshalJSON(b []byte) error {
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

type ShowGdgwRoutetableStatus struct {
	value string
}

type ShowGdgwRoutetableStatusEnum struct {
	ACTIVE         ShowGdgwRoutetableStatus
	ERROR          ShowGdgwRoutetableStatus
	PENDING_CREATE ShowGdgwRoutetableStatus
}

func GetShowGdgwRoutetableStatusEnum() ShowGdgwRoutetableStatusEnum {
	return ShowGdgwRoutetableStatusEnum{
		ACTIVE: ShowGdgwRoutetableStatus{
			value: "ACTIVE",
		},
		ERROR: ShowGdgwRoutetableStatus{
			value: "ERROR",
		},
		PENDING_CREATE: ShowGdgwRoutetableStatus{
			value: "PENDING_CREATE",
		},
	}
}

func (c ShowGdgwRoutetableStatus) Value() string {
	return c.value
}

func (c ShowGdgwRoutetableStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowGdgwRoutetableStatus) UnmarshalJSON(b []byte) error {
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
