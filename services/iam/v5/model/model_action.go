package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Action 授权项。
type Action struct {

	// 三段式的授权项名称，例如\"iam:policies:createV5\"。
	Name string `json:"name"`

	// 在策略中使用此授权项时授予的访问级别。
	AccessLevel ActionAccessLevel `json:"access_level"`

	// 该授权项是否仅作为权限点，不对应任何操作。
	PermissionOnly bool `json:"permission_only"`

	Description *Description `json:"description"`

	// 授权项别名列表，用以兼容授权项改名或者拆分新授权项的场景。
	Aliases *[]string `json:"aliases,omitempty"`

	// 与该授权项关联的资源列表，用于定义授权项的资源级权限。
	Resources *[]ActionAssociatedResource `json:"resources,omitempty"`

	// 该授权项支持的，且与资源无关的服务自定义条件属性以及部分全局属性。
	ConditionKeys *[]string `json:"condition_keys,omitempty"`
}

func (o Action) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Action struct{}"
	}

	return strings.Join([]string{"Action", string(data)}, " ")
}

type ActionAccessLevel struct {
	value string
}

type ActionAccessLevelEnum struct {
	LIST                  ActionAccessLevel
	READ                  ActionAccessLevel
	WRITE                 ActionAccessLevel
	PERMISSION_MANAGEMENT ActionAccessLevel
	TAGGING               ActionAccessLevel
}

func GetActionAccessLevelEnum() ActionAccessLevelEnum {
	return ActionAccessLevelEnum{
		LIST: ActionAccessLevel{
			value: "list",
		},
		READ: ActionAccessLevel{
			value: "read",
		},
		WRITE: ActionAccessLevel{
			value: "write",
		},
		PERMISSION_MANAGEMENT: ActionAccessLevel{
			value: "permission_management",
		},
		TAGGING: ActionAccessLevel{
			value: "tagging",
		},
	}
}

func (c ActionAccessLevel) Value() string {
	return c.value
}

func (c ActionAccessLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ActionAccessLevel) UnmarshalJSON(b []byte) error {
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
