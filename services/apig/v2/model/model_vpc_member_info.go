package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/sdktime"
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"
	"errors"
	"strings"
)

type VpcMemberInfo struct {
	// 后端服务器地址  后端实例类型为ip时必填

	Host *string `json:"host,omitempty"`
	// 权重值。  允许您对后端服务进行评级，权重值越大，转发到该云服务的请求数量越多。

	Weight *int32 `json:"weight,omitempty"`
	// 是否备用节点。  开启后对应后端服务为备用节点，仅当非备用节点全部故障时工作。  实例需要升级到对应版本才支持此功能，若不支持请联系技术支持。

	IsBackup *bool `json:"is_backup,omitempty"`
	// 后端服务器组名称。为后端服务地址选择服务器组，便于统一修改对应服务器组的后端地址。  暂不支持

	MemberGroupName *string `json:"member_group_name,omitempty"`
	// 后端服务器状态   - 1：可用   - 2：不可用

	Status *VpcMemberInfoStatus `json:"status,omitempty"`
	// 后端服务器端口

	Port *int32 `json:"port,omitempty"`
	// 后端云服务器的编号。  后端实例类型为instance时生效，支持英文，数字，“-”,“_”，1 ~ 64字符。

	EcsId *string `json:"ecs_id,omitempty"`
	// 后端云服务器的名称。  后端实例类型为instance时生效，支持汉字，英文，数字，“-”,“_”,“.”，1 ~ 64字符。

	EcsName *string `json:"ecs_name,omitempty"`
	// 后端实例对象的编号

	Id *string `json:"id,omitempty"`
	// VPC通道的编号

	VpcChannelId *string `json:"vpc_channel_id,omitempty"`
	// 后端实例增加到VPC通道的时间

	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`
	// 后端服务器组编号  暂不支持

	MemberGroupId *string `json:"member_group_id,omitempty"`
}

func (o VpcMemberInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcMemberInfo struct{}"
	}

	return strings.Join([]string{"VpcMemberInfo", string(data)}, " ")
}

type VpcMemberInfoStatus struct {
	value int32
}

type VpcMemberInfoStatusEnum struct {
	E_1 VpcMemberInfoStatus
	E_2 VpcMemberInfoStatus
}

func GetVpcMemberInfoStatusEnum() VpcMemberInfoStatusEnum {
	return VpcMemberInfoStatusEnum{
		E_1: VpcMemberInfoStatus{
			value: 1,
		}, E_2: VpcMemberInfoStatus{
			value: 2,
		},
	}
}

func (c VpcMemberInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VpcMemberInfoStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
