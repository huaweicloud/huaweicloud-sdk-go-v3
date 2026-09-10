package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MigrationBindWorkspace 资源关联的工作空间信息。
type MigrationBindWorkspace struct {

	// 工作空间ID。
	Id string `json:"id"`

	// 操作类型，band表示关联资源到工作空间，remove表示取消关联。
	Action MigrationBindWorkspaceAction `json:"action"`
}

func (o MigrationBindWorkspace) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrationBindWorkspace struct{}"
	}

	return strings.Join([]string{"MigrationBindWorkspace", string(data)}, " ")
}

type MigrationBindWorkspaceAction struct {
	value string
}

type MigrationBindWorkspaceActionEnum struct {
	BAND   MigrationBindWorkspaceAction
	REMOVE MigrationBindWorkspaceAction
}

func GetMigrationBindWorkspaceActionEnum() MigrationBindWorkspaceActionEnum {
	return MigrationBindWorkspaceActionEnum{
		BAND: MigrationBindWorkspaceAction{
			value: "band",
		},
		REMOVE: MigrationBindWorkspaceAction{
			value: "remove",
		},
	}
}

func (c MigrationBindWorkspaceAction) Value() string {
	return c.value
}

func (c MigrationBindWorkspaceAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MigrationBindWorkspaceAction) UnmarshalJSON(b []byte) error {
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
