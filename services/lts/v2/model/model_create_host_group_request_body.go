package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 创建主机组请求体
type CreateHostGroupRequestBody struct {
	// 主机组名称

	HostGroupName string `json:"host_group_name"`
	// 主机组类型。windows：windows类型，linux：linux类型

	HostGroupType CreateHostGroupRequestBodyHostGroupType `json:"host_group_type"`
	// 主机组ID列表。主机类型必须与主机组类型一致

	HostIdList *[]string `json:"host_id_list,omitempty"`
	// 标签信息。KEY不能重复

	HostGroupTag *[]HostGroupTag `json:"host_group_tag,omitempty"`
}

func (o CreateHostGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHostGroupRequestBody struct{}"
	}

	return strings.Join([]string{"CreateHostGroupRequestBody", string(data)}, " ")
}

type CreateHostGroupRequestBodyHostGroupType struct {
	value string
}

type CreateHostGroupRequestBodyHostGroupTypeEnum struct {
	LINUX   CreateHostGroupRequestBodyHostGroupType
	WINDOWS CreateHostGroupRequestBodyHostGroupType
}

func GetCreateHostGroupRequestBodyHostGroupTypeEnum() CreateHostGroupRequestBodyHostGroupTypeEnum {
	return CreateHostGroupRequestBodyHostGroupTypeEnum{
		LINUX: CreateHostGroupRequestBodyHostGroupType{
			value: "linux",
		},
		WINDOWS: CreateHostGroupRequestBodyHostGroupType{
			value: "windows",
		},
	}
}

func (c CreateHostGroupRequestBodyHostGroupType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateHostGroupRequestBodyHostGroupType) UnmarshalJSON(b []byte) error {
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
