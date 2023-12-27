package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateTagsReqbody struct {

	// 添加标签方式
	Action CreateTagsReqbodyAction `json:"action"`

	// 是否对外接口调用
	IsOpen bool `json:"is_open"`

	// 标签字段信息
	Tags []TagsBody `json:"tags"`
}

func (o CreateTagsReqbody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTagsReqbody struct{}"
	}

	return strings.Join([]string{"CreateTagsReqbody", string(data)}, " ")
}

type CreateTagsReqbodyAction struct {
	value string
}

type CreateTagsReqbodyActionEnum struct {
	CREATE CreateTagsReqbodyAction
	DELETE CreateTagsReqbodyAction
}

func GetCreateTagsReqbodyActionEnum() CreateTagsReqbodyActionEnum {
	return CreateTagsReqbodyActionEnum{
		CREATE: CreateTagsReqbodyAction{
			value: "create",
		},
		DELETE: CreateTagsReqbodyAction{
			value: "delete",
		},
	}
}

func (c CreateTagsReqbodyAction) Value() string {
	return c.value
}

func (c CreateTagsReqbodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTagsReqbodyAction) UnmarshalJSON(b []byte) error {
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
