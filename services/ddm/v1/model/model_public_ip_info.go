package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PublicIpInfo DDM绑定的弹性公网IP信息
type PublicIpInfo struct {

	// 弹性公网IP唯一标识。
	Id *string `json:"id,omitempty"`

	// 弹性公网IP的类型。
	Type *PublicIpInfoType `json:"type,omitempty"`

	// 弹性公网IP的地址。
	PublicIp *string `json:"public_ip,omitempty"`

	// 绑定弹性公网IP的私有IP地址。
	PrivateIp *string `json:"private_ip,omitempty"`

	// 绑定弹性公网IP的实体ID。
	BindingEntityId *string `json:"binding_entity_id,omitempty"`

	// 绑定弹性公网IP的实体名称。
	BindingEntityName *string `json:"binding_entity_name,omitempty"`

	// 绑定弹性公网IP的实体类型（目前只支持绑定node）。
	BindingEntityType *PublicIpInfoBindingEntityType `json:"binding_entity_type,omitempty"`
}

func (o PublicIpInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicIpInfo struct{}"
	}

	return strings.Join([]string{"PublicIpInfo", string(data)}, " ")
}

type PublicIpInfoType struct {
	value string
}

type PublicIpInfoTypeEnum struct {
	E_5_BGPBGP  PublicIpInfoType
	E_5_SBGPBGP PublicIpInfoType
}

func GetPublicIpInfoTypeEnum() PublicIpInfoTypeEnum {
	return PublicIpInfoTypeEnum{
		E_5_BGPBGP: PublicIpInfoType{
			value: "5_bgp（全动态BGP）",
		},
		E_5_SBGPBGP: PublicIpInfoType{
			value: "5_sbgp（静态BGP）",
		},
	}
}

func (c PublicIpInfoType) Value() string {
	return c.value
}

func (c PublicIpInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PublicIpInfoType) UnmarshalJSON(b []byte) error {
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

type PublicIpInfoBindingEntityType struct {
	value string
}

type PublicIpInfoBindingEntityTypeEnum struct {
	NODE       PublicIpInfoBindingEntityType
	NODE_GROUP PublicIpInfoBindingEntityType
}

func GetPublicIpInfoBindingEntityTypeEnum() PublicIpInfoBindingEntityTypeEnum {
	return PublicIpInfoBindingEntityTypeEnum{
		NODE: PublicIpInfoBindingEntityType{
			value: "node（节点）",
		},
		NODE_GROUP: PublicIpInfoBindingEntityType{
			value: "node_group（节点组）",
		},
	}
}

func (c PublicIpInfoBindingEntityType) Value() string {
	return c.value
}

func (c PublicIpInfoBindingEntityType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PublicIpInfoBindingEntityType) UnmarshalJSON(b []byte) error {
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
