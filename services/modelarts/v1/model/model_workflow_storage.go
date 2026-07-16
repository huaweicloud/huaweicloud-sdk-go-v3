package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// WorkflowStorage 工作流出存储信息。
type WorkflowStorage struct {

	// 工作流存储的名称。填写1-64位，只包含英文、数字、下划线（_）和中划线（-），并且以英文开头的名称。
	Name *string `json:"name,omitempty"`

	// 工作流存储的类型，当前只支持obs。
	Type *WorkflowStorageType `json:"type,omitempty"`

	// 统一存储的根路径，当前只支持OBS路径。
	Path *string `json:"path,omitempty"`
}

func (o WorkflowStorage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowStorage struct{}"
	}

	return strings.Join([]string{"WorkflowStorage", string(data)}, " ")
}

type WorkflowStorageType struct {
	value string
}

type WorkflowStorageTypeEnum struct {
	OBS WorkflowStorageType
}

func GetWorkflowStorageTypeEnum() WorkflowStorageTypeEnum {
	return WorkflowStorageTypeEnum{
		OBS: WorkflowStorageType{
			value: "obs",
		},
	}
}

func (c WorkflowStorageType) Value() string {
	return c.value
}

func (c WorkflowStorageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WorkflowStorageType) UnmarshalJSON(b []byte) error {
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
