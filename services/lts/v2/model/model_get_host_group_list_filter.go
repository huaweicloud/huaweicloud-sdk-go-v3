package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 主机组过滤参数
type GetHostGroupListFilter struct {
	// 主机组类型。windows：windows类型，linux：linux类型

	HostGroupType *GetHostGroupListFilterHostGroupType `json:"host_group_type,omitempty"`
	// 主机组名称列表。

	HostGroupNameList *[]string `json:"host_group_name_list,omitempty"`
	// 主机名称列表。

	HostNameList *[]string `json:"host_name_list,omitempty"`

	HostGroupTag *GetHostGroupListTag `json:"host_group_tag,omitempty"`
}

func (o GetHostGroupListFilter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetHostGroupListFilter struct{}"
	}

	return strings.Join([]string{"GetHostGroupListFilter", string(data)}, " ")
}

type GetHostGroupListFilterHostGroupType struct {
	value string
}

type GetHostGroupListFilterHostGroupTypeEnum struct {
	WINDOWS GetHostGroupListFilterHostGroupType
	LINUX   GetHostGroupListFilterHostGroupType
}

func GetGetHostGroupListFilterHostGroupTypeEnum() GetHostGroupListFilterHostGroupTypeEnum {
	return GetHostGroupListFilterHostGroupTypeEnum{
		WINDOWS: GetHostGroupListFilterHostGroupType{
			value: "windows",
		},
		LINUX: GetHostGroupListFilterHostGroupType{
			value: "linux",
		},
	}
}

func (c GetHostGroupListFilterHostGroupType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetHostGroupListFilterHostGroupType) UnmarshalJSON(b []byte) error {
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
