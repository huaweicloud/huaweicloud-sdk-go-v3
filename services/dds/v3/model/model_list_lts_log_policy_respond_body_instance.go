package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListLtsLogPolicyRespondBodyInstance 实例的简要信息。
type ListLtsLogPolicyRespondBodyInstance struct {

	// 实例ID。
	Id *string `json:"id,omitempty"`

	// 实例名字。
	Name *string `json:"name,omitempty"`

	// 实例类型，对应单节点，副本集，集群三种。
	Mode *ListLtsLogPolicyRespondBodyInstanceMode `json:"mode,omitempty"`

	Datastore *ListLtsLogPolicyRespondBodyInstanceDatastore `json:"datastore,omitempty"`

	// 实例状态。
	Status *string `json:"status,omitempty"`

	// 实例归属的企业项目ID，如果是default企业项目，值为0，如果是其他企业项目，请参考企业项目相关内容。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 实例所有正在执行的动作。
	Actions *[]string `json:"actions,omitempty"`
}

func (o ListLtsLogPolicyRespondBodyInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLtsLogPolicyRespondBodyInstance struct{}"
	}

	return strings.Join([]string{"ListLtsLogPolicyRespondBodyInstance", string(data)}, " ")
}

type ListLtsLogPolicyRespondBodyInstanceMode struct {
	value string
}

type ListLtsLogPolicyRespondBodyInstanceModeEnum struct {
	SINGLE   ListLtsLogPolicyRespondBodyInstanceMode
	REPLICA  ListLtsLogPolicyRespondBodyInstanceMode
	SHARDING ListLtsLogPolicyRespondBodyInstanceMode
}

func GetListLtsLogPolicyRespondBodyInstanceModeEnum() ListLtsLogPolicyRespondBodyInstanceModeEnum {
	return ListLtsLogPolicyRespondBodyInstanceModeEnum{
		SINGLE: ListLtsLogPolicyRespondBodyInstanceMode{
			value: "Single",
		},
		REPLICA: ListLtsLogPolicyRespondBodyInstanceMode{
			value: "Replica",
		},
		SHARDING: ListLtsLogPolicyRespondBodyInstanceMode{
			value: "Sharding",
		},
	}
}

func (c ListLtsLogPolicyRespondBodyInstanceMode) Value() string {
	return c.value
}

func (c ListLtsLogPolicyRespondBodyInstanceMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListLtsLogPolicyRespondBodyInstanceMode) UnmarshalJSON(b []byte) error {
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
