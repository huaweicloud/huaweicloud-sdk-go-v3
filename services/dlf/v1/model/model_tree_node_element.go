package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TreeNodeElement struct {
	Type *TreeNodeElementType `json:"type,omitempty"`

	ParentDirectoryId *string `json:"parent_directory_id,omitempty"`

	Name *string `json:"name,omitempty"`

	ElementId *string `json:"element_id,omitempty"`

	Owner *string `json:"owner,omitempty"`

	ProcessType *string `json:"process_type,omitempty"`

	IsSingleNodeJob *bool `json:"is_single_node_job,omitempty"`
}

func (o TreeNodeElement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TreeNodeElement struct{}"
	}

	return strings.Join([]string{"TreeNodeElement", string(data)}, " ")
}

type TreeNodeElementType struct {
	value string
}

type TreeNodeElementTypeEnum struct {
	JOB    TreeNodeElementType
	SCRIPT TreeNodeElementType
}

func GetTreeNodeElementTypeEnum() TreeNodeElementTypeEnum {
	return TreeNodeElementTypeEnum{
		JOB: TreeNodeElementType{
			value: "job",
		},
		SCRIPT: TreeNodeElementType{
			value: "script",
		},
	}
}

func (c TreeNodeElementType) Value() string {
	return c.value
}

func (c TreeNodeElementType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TreeNodeElementType) UnmarshalJSON(b []byte) error {
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
