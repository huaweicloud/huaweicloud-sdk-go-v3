package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 创建工作流的请求Body
type CreateWorkflowRequestBody struct {

	// 工作流的编排定义,必须有TYPE，TYPE值必须是3种State（DELAY，OPERATION，END）中一种。每个state的名字是1-80长度的只含数字，字母，-和_的String。
	States *[]State `json:"states,omitempty"`

	// 工作流中用户可修改的参数项
	Inputs *[]Input `json:"inputs,omitempty"`

	// 工作流的描述
	Description *string `json:"description,omitempty"`

	// 工作流执行类型：同步（EXPRESS）、异步（NORMAL）
	Mode *CreateWorkflowRequestBodyMode `json:"mode,omitempty"`

	ExpressConfig *ExpressConfig `json:"express_config,omitempty"`

	FuncVpc *FuncVpc `json:"func_vpc,omitempty"`

	// 用戶传入用于创建工作流时使用的委托的委托名
	Agency *string `json:"agency,omitempty"`
}

func (o CreateWorkflowRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkflowRequestBody struct{}"
	}

	return strings.Join([]string{"CreateWorkflowRequestBody", string(data)}, " ")
}

type CreateWorkflowRequestBodyMode struct {
	value string
}

type CreateWorkflowRequestBodyModeEnum struct {
	EXPRESS CreateWorkflowRequestBodyMode
	NORMAL  CreateWorkflowRequestBodyMode
}

func GetCreateWorkflowRequestBodyModeEnum() CreateWorkflowRequestBodyModeEnum {
	return CreateWorkflowRequestBodyModeEnum{
		EXPRESS: CreateWorkflowRequestBodyMode{
			value: "EXPRESS",
		},
		NORMAL: CreateWorkflowRequestBodyMode{
			value: "NORMAL",
		},
	}
}

func (c CreateWorkflowRequestBodyMode) Value() string {
	return c.value
}

func (c CreateWorkflowRequestBodyMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateWorkflowRequestBodyMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
