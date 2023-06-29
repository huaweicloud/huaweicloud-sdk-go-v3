package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateHostGroupResponse Response Object
type CreateHostGroupResponse struct {

	// 主机组ID
	HostGroupId *string `json:"host_group_id,omitempty"`

	// 主机组名称
	HostGroupName *string `json:"host_group_name,omitempty"`

	// 主机组类型。linux：linux类型，windows：windows类型
	HostGroupType *CreateHostGroupResponseHostGroupType `json:"host_group_type,omitempty"`

	// 主机ID列表
	HostIdList *[]string `json:"host_id_list,omitempty"`

	// 标签信息
	HostGroupTag *[]HostGroupTag `json:"host_group_tag,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime     *int64 `json:"update_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateHostGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHostGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateHostGroupResponse", string(data)}, " ")
}

type CreateHostGroupResponseHostGroupType struct {
	value string
}

type CreateHostGroupResponseHostGroupTypeEnum struct {
	LINUX   CreateHostGroupResponseHostGroupType
	WINDOWS CreateHostGroupResponseHostGroupType
}

func GetCreateHostGroupResponseHostGroupTypeEnum() CreateHostGroupResponseHostGroupTypeEnum {
	return CreateHostGroupResponseHostGroupTypeEnum{
		LINUX: CreateHostGroupResponseHostGroupType{
			value: "linux",
		},
		WINDOWS: CreateHostGroupResponseHostGroupType{
			value: "windows",
		},
	}
}

func (c CreateHostGroupResponseHostGroupType) Value() string {
	return c.value
}

func (c CreateHostGroupResponseHostGroupType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateHostGroupResponseHostGroupType) UnmarshalJSON(b []byte) error {
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
