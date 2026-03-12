package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateDryRunPolicyReqBody 创建试运行策略操作的请求体。
type CreateDryRunPolicyReqBody struct {

	// 要添加到新策略的策略文本内容。
	Content string `json:"content"`

	// 要分配给策略的可选说明。
	Description string `json:"description"`

	// 要分配给策略的名称。
	Name string `json:"name"`

	// 要创建的策略类型,service_control_policy服务控制策略。
	Type CreateDryRunPolicyReqBodyType `json:"type"`

	// 要附加到新创建的策略的标签列表。
	Tags *[]TagDto `json:"tags,omitempty"`
}

func (o CreateDryRunPolicyReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDryRunPolicyReqBody struct{}"
	}

	return strings.Join([]string{"CreateDryRunPolicyReqBody", string(data)}, " ")
}

type CreateDryRunPolicyReqBodyType struct {
	value string
}

type CreateDryRunPolicyReqBodyTypeEnum struct {
	SERVICE_CONTROL_POLICY CreateDryRunPolicyReqBodyType
}

func GetCreateDryRunPolicyReqBodyTypeEnum() CreateDryRunPolicyReqBodyTypeEnum {
	return CreateDryRunPolicyReqBodyTypeEnum{
		SERVICE_CONTROL_POLICY: CreateDryRunPolicyReqBodyType{
			value: "service_control_policy",
		},
	}
}

func (c CreateDryRunPolicyReqBodyType) Value() string {
	return c.value
}

func (c CreateDryRunPolicyReqBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDryRunPolicyReqBodyType) UnmarshalJSON(b []byte) error {
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
