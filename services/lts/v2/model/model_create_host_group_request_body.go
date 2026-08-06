package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateHostGroupRequestBody 创建主机组请求体
type CreateHostGroupRequestBody struct {

	// 主机组名称
	HostGroupName string `json:"host_group_name"`

	// 主机组类型。windows：windows类型，linux：linux类型
	HostGroupType CreateHostGroupRequestBodyHostGroupType `json:"host_group_type"`

	// 主机组ID列表。主机类型必须与主机组类型一致
	HostIdList *[]string `json:"host_id_list,omitempty"`

	// 标签信息。KEY不能重复
	HostGroupTag *[]HostGroupTag `json:"host_group_tag,omitempty"`

	// **参数解释：** 主机组类型。支持两种主机组类型，分别为IP类型和LABEL类型。 **约束限制：** 不涉及 **取值范围：** - IP - LABEL **默认取值：** IP。
	AgentAccessType *string `json:"agent_access_type,omitempty"`

	// **参数解释：** 自定义标识。主机组类型为LABEL类型，该字段必填。 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	Labels *[]string `json:"labels,omitempty"`
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

func (c CreateHostGroupRequestBodyHostGroupType) Value() string {
	return c.value
}

func (c CreateHostGroupRequestBodyHostGroupType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateHostGroupRequestBodyHostGroupType) UnmarshalJSON(b []byte) error {
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
