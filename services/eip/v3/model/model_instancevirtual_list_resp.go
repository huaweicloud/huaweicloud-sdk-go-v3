package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// InstancevirtualListResp GEIP实例的vn信息
type InstancevirtualListResp struct {

	// virtualnexthop的uuid
	Id *string `json:"id,omitempty"`

	// virtualnexthop的所有者
	Owner *string `json:"owner,omitempty"`

	// 标识网关所在位置POD、AZ、REGION、GLOBAL
	Location *string `json:"location,omitempty"`

	// 功能说明：nexthops的转发模式  取值范围：'ACTIVE-ACTIVE'多活模式、'ACTIVE-STANDBY'主备模式
	ForwardMode *InstancevirtualListRespForwardMode `json:"forward_mode,omitempty"`

	// 功能说明：网关所在集群信息，可为空  取值范围：0-36长度的字符串
	ClusterId *string `json:"cluster_id,omitempty"`

	// 功能说明：nexthops在底层的负载均衡策略  取值范围：'2_TUPLE'二元组、'3_TUPLE'三元组、'5_TUPLE'五元组
	HashMode *InstancevirtualListRespHashMode `json:"hash_mode,omitempty"`

	// 功能说明：下一跳所属网络类型  取值范围：'VXLAN'、'VLAN'
	Type *InstancevirtualListRespType `json:"type,omitempty"`

	// 功能说明：网络id标识，与type组合使用  取值范围：type=VXLAN时取值0-16777215,type=VLAN时取值0-4095
	Vni *int32 `json:"vni,omitempty"`

	// 下一跳信息列表
	Nexthops *[]NexthopDict `json:"nexthops,omitempty"`

	// 功能说明：VirtualNexthop对象创建时间，UTC格式
	CreatedAt *string `json:"created_at,omitempty"`

	// 功能说明：VirtualNexthop对象更新时间，UTC格式
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o InstancevirtualListResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstancevirtualListResp struct{}"
	}

	return strings.Join([]string{"InstancevirtualListResp", string(data)}, " ")
}

type InstancevirtualListRespForwardMode struct {
	value string
}

type InstancevirtualListRespForwardModeEnum struct {
	ACTIVE_ACTIVE  InstancevirtualListRespForwardMode
	ACTIVE_STANDBY InstancevirtualListRespForwardMode
}

func GetInstancevirtualListRespForwardModeEnum() InstancevirtualListRespForwardModeEnum {
	return InstancevirtualListRespForwardModeEnum{
		ACTIVE_ACTIVE: InstancevirtualListRespForwardMode{
			value: "ACTIVE-ACTIVE",
		},
		ACTIVE_STANDBY: InstancevirtualListRespForwardMode{
			value: "ACTIVE-STANDBY",
		},
	}
}

func (c InstancevirtualListRespForwardMode) Value() string {
	return c.value
}

func (c InstancevirtualListRespForwardMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstancevirtualListRespForwardMode) UnmarshalJSON(b []byte) error {
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

type InstancevirtualListRespHashMode struct {
	value string
}

type InstancevirtualListRespHashModeEnum struct {
	E_2_TUPLE InstancevirtualListRespHashMode
	E_3_TUPLE InstancevirtualListRespHashMode
	E_5_TUPLE InstancevirtualListRespHashMode
}

func GetInstancevirtualListRespHashModeEnum() InstancevirtualListRespHashModeEnum {
	return InstancevirtualListRespHashModeEnum{
		E_2_TUPLE: InstancevirtualListRespHashMode{
			value: "2_TUPLE",
		},
		E_3_TUPLE: InstancevirtualListRespHashMode{
			value: "3_TUPLE",
		},
		E_5_TUPLE: InstancevirtualListRespHashMode{
			value: "5_TUPLE",
		},
	}
}

func (c InstancevirtualListRespHashMode) Value() string {
	return c.value
}

func (c InstancevirtualListRespHashMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstancevirtualListRespHashMode) UnmarshalJSON(b []byte) error {
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

type InstancevirtualListRespType struct {
	value string
}

type InstancevirtualListRespTypeEnum struct {
	VLAN  InstancevirtualListRespType
	VXLAN InstancevirtualListRespType
}

func GetInstancevirtualListRespTypeEnum() InstancevirtualListRespTypeEnum {
	return InstancevirtualListRespTypeEnum{
		VLAN: InstancevirtualListRespType{
			value: "VLAN",
		},
		VXLAN: InstancevirtualListRespType{
			value: "VXLAN",
		},
	}
}

func (c InstancevirtualListRespType) Value() string {
	return c.value
}

func (c InstancevirtualListRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstancevirtualListRespType) UnmarshalJSON(b []byte) error {
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
