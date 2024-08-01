package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateGroupRequestBody struct {

	// 组名称
	Name string `json:"name"`

	// 组类型，type：rw读写、r只读
	Type CreateGroupRequestBodyType `json:"type"`

	// 节点规格ID。
	FlavorId string `json:"flavor_id"`

	// 节点信息列表
	Nodes []NodeInfo `json:"nodes"`
}

func (o CreateGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGroupRequestBody struct{}"
	}

	return strings.Join([]string{"CreateGroupRequestBody", string(data)}, " ")
}

type CreateGroupRequestBodyType struct {
	value string
}

type CreateGroupRequestBodyTypeEnum struct {
	RW CreateGroupRequestBodyType
	R  CreateGroupRequestBodyType
}

func GetCreateGroupRequestBodyTypeEnum() CreateGroupRequestBodyTypeEnum {
	return CreateGroupRequestBodyTypeEnum{
		RW: CreateGroupRequestBodyType{
			value: "rw",
		},
		R: CreateGroupRequestBodyType{
			value: "r",
		},
	}
}

func (c CreateGroupRequestBodyType) Value() string {
	return c.value
}

func (c CreateGroupRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateGroupRequestBodyType) UnmarshalJSON(b []byte) error {
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
