package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type MemberInfo struct {

	// 后端服务器地址  后端实例类型为ip时必填
	Host *string `json:"host,omitempty"`

	// 权重值。  允许您对后端服务进行评级，权重值越大，转发到该云服务的请求数量越多。
	Weight *int32 `json:"weight,omitempty"`

	// 是否备用节点。  开启后对应后端服务为备用节点，仅当非备用节点全部故障时工作。  实例需要升级到对应版本才支持此功能，如果不支持请联系技术支持。
	IsBackup *bool `json:"is_backup,omitempty"`

	// 后端服务器组名称。为后端服务地址选择服务器组，便于统一修改对应服务器组的后端地址。
	MemberGroupName *string `json:"member_group_name,omitempty"`

	// 后端服务器状态   - 1：可用   - 2：不可用
	Status *MemberInfoStatus `json:"status,omitempty"`

	// 后端服务器端口
	Port *int32 `json:"port,omitempty"`

	// 后端云服务器的编号。  后端实例类型为ecs时必填，支持英文，数字，“-”,“_”，1 ~ 64字符。
	EcsId *string `json:"ecs_id,omitempty"`

	// 后端云服务器的名称。  后端实例类型为ecs时必填，支持汉字，英文，数字，“-”,“_”,“.”，1 ~ 64字符。
	EcsName *string `json:"ecs_name,omitempty"`
}

func (o MemberInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberInfo struct{}"
	}

	return strings.Join([]string{"MemberInfo", string(data)}, " ")
}

type MemberInfoStatus struct {
	value int32
}

type MemberInfoStatusEnum struct {
	E_1 MemberInfoStatus
	E_2 MemberInfoStatus
}

func GetMemberInfoStatusEnum() MemberInfoStatusEnum {
	return MemberInfoStatusEnum{
		E_1: MemberInfoStatus{
			value: 1,
		}, E_2: MemberInfoStatus{
			value: 2,
		},
	}
}

func (c MemberInfoStatus) Value() int32 {
	return c.value
}

func (c MemberInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MemberInfoStatus) UnmarshalJSON(b []byte) error {
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
