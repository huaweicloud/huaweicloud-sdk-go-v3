package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type GroupInfo struct {

	// 组ID。
	Id string `json:"id"`

	// 组名称。
	Name string `json:"name"`

	// 组角色类型。
	Role GroupInfoRole `json:"role"`

	// 组连接地址，如未开启负载均衡，则返回的是组内节点连接地址串。
	Endpoint string `json:"endpoint"`

	// 组ipv6连接地址。
	Ipv6Endpoint *string `json:"ipv6_endpoint,omitempty"`

	// 是否开启负载均衡。
	IsLoadBalance bool `json:"is_load_balance"`

	// 是否默认组。
	IsDefaultGroup bool `json:"is_default_group"`

	// 单节点CPU核数。
	CpuNumPerNode int32 `json:"cpu_num_per_node"`

	// 单节点内存大小,单位G。
	MemNumPerNode int32 `json:"mem_num_per_node"`

	// CPU架构。
	Architecture GroupInfoArchitecture `json:"architecture"`

	// 节点信息列表。
	NodeList []GroupNodeInfo `json:"node_list"`
}

func (o GroupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupInfo struct{}"
	}

	return strings.Join([]string{"GroupInfo", string(data)}, " ")
}

type GroupInfoRole struct {
	value string
}

type GroupInfoRoleEnum struct {
	RW GroupInfoRole
	R  GroupInfoRole
}

func GetGroupInfoRoleEnum() GroupInfoRoleEnum {
	return GroupInfoRoleEnum{
		RW: GroupInfoRole{
			value: "rw",
		},
		R: GroupInfoRole{
			value: "r",
		},
	}
}

func (c GroupInfoRole) Value() string {
	return c.value
}

func (c GroupInfoRole) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GroupInfoRole) UnmarshalJSON(b []byte) error {
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

type GroupInfoArchitecture struct {
	value string
}

type GroupInfoArchitectureEnum struct {
	X86 GroupInfoArchitecture
	ARM GroupInfoArchitecture
}

func GetGroupInfoArchitectureEnum() GroupInfoArchitectureEnum {
	return GroupInfoArchitectureEnum{
		X86: GroupInfoArchitecture{
			value: "X86",
		},
		ARM: GroupInfoArchitecture{
			value: "ARM",
		},
	}
}

func (c GroupInfoArchitecture) Value() string {
	return c.value
}

func (c GroupInfoArchitecture) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GroupInfoArchitecture) UnmarshalJSON(b []byte) error {
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
