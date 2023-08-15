package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ActionsOfUsersInGroupRequest struct {

	// 要添加或移除的用户Id列表。
	UserIds []string `json:"user_ids"`

	// 操作类型。 * ADD： 添加 * DELETE： 删除
	OpType ActionsOfUsersInGroupRequestOpType `json:"op_type"`
}

func (o ActionsOfUsersInGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionsOfUsersInGroupRequest struct{}"
	}

	return strings.Join([]string{"ActionsOfUsersInGroupRequest", string(data)}, " ")
}

type ActionsOfUsersInGroupRequestOpType struct {
	value string
}

type ActionsOfUsersInGroupRequestOpTypeEnum struct {
	ADD    ActionsOfUsersInGroupRequestOpType
	DELETE ActionsOfUsersInGroupRequestOpType
}

func GetActionsOfUsersInGroupRequestOpTypeEnum() ActionsOfUsersInGroupRequestOpTypeEnum {
	return ActionsOfUsersInGroupRequestOpTypeEnum{
		ADD: ActionsOfUsersInGroupRequestOpType{
			value: "ADD",
		},
		DELETE: ActionsOfUsersInGroupRequestOpType{
			value: "DELETE",
		},
	}
}

func (c ActionsOfUsersInGroupRequestOpType) Value() string {
	return c.value
}

func (c ActionsOfUsersInGroupRequestOpType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ActionsOfUsersInGroupRequestOpType) UnmarshalJSON(b []byte) error {
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
