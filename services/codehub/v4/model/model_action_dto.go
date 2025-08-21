package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ActionDto **参数解释：** 提交动作设置参数。
type ActionDto struct {

	// **参数解释：** 要执行的操作。 **取值范围：** - create，创建文件。 - create_dir，创建目录。 - update，更新。 - move，移动。 - delete，删除。 - chmod，更改权限。
	Action ActionDtoAction `json:"action"`

	// **参数解释：** 文件路径。 **取值范围：** 不涉及。
	FilePath string `json:"file_path"`

	// **参数解释：** 上一个路径。 **取值范围：** 不涉及。
	PreviousPath *string `json:"previous_path,omitempty"`

	// **参数解释：** 文件内容。 **取值范围：** 不涉及。
	Content *string `json:"content,omitempty"`

	// **参数解释：** 编码方式。 **取值范围：** - text。 - base64.
	Encoding *string `json:"encoding,omitempty"`

	// **参数解释：** 上次已知的文件提交ID。 **取值范围：** 不涉及。
	LastCommitId *bool `json:"last_commit_id,omitempty"`

	// **参数解释：** 执行文件模式。 **取值范围：** - true，启用文件上的执行标志。 - false，禁用文件上的执行标志
	ExecuteFilemode *bool `json:"execute_filemode,omitempty"`
}

func (o ActionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionDto struct{}"
	}

	return strings.Join([]string{"ActionDto", string(data)}, " ")
}

type ActionDtoAction struct {
	value string
}

type ActionDtoActionEnum struct {
	CREATE     ActionDtoAction
	CREATE_DIR ActionDtoAction
	UPDATE     ActionDtoAction
	MOVE       ActionDtoAction
	DELETE     ActionDtoAction
	CHMOD      ActionDtoAction
}

func GetActionDtoActionEnum() ActionDtoActionEnum {
	return ActionDtoActionEnum{
		CREATE: ActionDtoAction{
			value: "create",
		},
		CREATE_DIR: ActionDtoAction{
			value: "create_dir",
		},
		UPDATE: ActionDtoAction{
			value: "update",
		},
		MOVE: ActionDtoAction{
			value: "move",
		},
		DELETE: ActionDtoAction{
			value: "delete",
		},
		CHMOD: ActionDtoAction{
			value: "chmod",
		},
	}
}

func (c ActionDtoAction) Value() string {
	return c.value
}

func (c ActionDtoAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ActionDtoAction) UnmarshalJSON(b []byte) error {
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
