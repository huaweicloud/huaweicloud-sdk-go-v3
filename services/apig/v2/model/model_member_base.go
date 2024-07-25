package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type MemberBase struct {

	// 后端服务器地址  后端实例类型为ip时必填
	Host *string `json:"host,omitempty"`

	// 权重值。  允许您对后端服务进行评级，权重值越大，转发到该云服务的请求数量越多。
	Weight *int32 `json:"weight,omitempty"`

	// 是否备用节点。  开启后对应后端服务为备用节点，仅当非备用节点全部故障时工作。  实例需要升级到对应版本才支持此功能，如果不支持请联系技术支持。
	IsBackup *bool `json:"is_backup,omitempty"`

	// 后端服务器组名称。为后端服务地址选择服务器组，便于统一修改对应服务器组的后端地址。
	MemberGroupName *string `json:"member_group_name,omitempty"`

	// 后端服务器状态   - 1：可用   - 2：不可用
	Status *MemberBaseStatus `json:"status,omitempty"`

	// 后端服务器端口
	Port *int32 `json:"port,omitempty"`
}

func (o MemberBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberBase struct{}"
	}

	return strings.Join([]string{"MemberBase", string(data)}, " ")
}

type MemberBaseStatus struct {
	value int32
}

type MemberBaseStatusEnum struct {
	E_1 MemberBaseStatus
	E_2 MemberBaseStatus
}

func GetMemberBaseStatusEnum() MemberBaseStatusEnum {
	return MemberBaseStatusEnum{
		E_1: MemberBaseStatus{
			value: 1,
		}, E_2: MemberBaseStatus{
			value: 2,
		},
	}
}

func (c MemberBaseStatus) Value() int32 {
	return c.value
}

func (c MemberBaseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MemberBaseStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
