package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 主机组详细信息
type GetHostGroupInfo struct {

	// 主机组ID
	HostGroupId *string `json:"host_group_id,omitempty"`

	// 主机组名称
	HostGroupName *string `json:"host_group_name,omitempty"`

	// 主机组类型。linux：linux类型，windows：windows类型
	HostGroupType *GetHostGroupInfoHostGroupType `json:"host_group_type,omitempty"`

	// 主机ID列表
	HostIdList *[]string `json:"host_id_list,omitempty"`

	// 标签信息
	HostGroupTag *[]HostGroupTag `json:"host_group_tag,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o GetHostGroupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetHostGroupInfo struct{}"
	}

	return strings.Join([]string{"GetHostGroupInfo", string(data)}, " ")
}

type GetHostGroupInfoHostGroupType struct {
	value string
}

type GetHostGroupInfoHostGroupTypeEnum struct {
	LINUX   GetHostGroupInfoHostGroupType
	WINDOWS GetHostGroupInfoHostGroupType
}

func GetGetHostGroupInfoHostGroupTypeEnum() GetHostGroupInfoHostGroupTypeEnum {
	return GetHostGroupInfoHostGroupTypeEnum{
		LINUX: GetHostGroupInfoHostGroupType{
			value: "linux",
		},
		WINDOWS: GetHostGroupInfoHostGroupType{
			value: "windows",
		},
	}
}

func (c GetHostGroupInfoHostGroupType) Value() string {
	return c.value
}

func (c GetHostGroupInfoHostGroupType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetHostGroupInfoHostGroupType) UnmarshalJSON(b []byte) error {
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
