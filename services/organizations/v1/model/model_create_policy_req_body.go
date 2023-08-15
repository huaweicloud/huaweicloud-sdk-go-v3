package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreatePolicyReqBody CreatePolicy 操作的请求体。
type CreatePolicyReqBody struct {

	// 要添加到新策略的策略文本内容。
	Content string `json:"content"`

	// 要分配给策略的可选说明。
	Description string `json:"description"`

	// 要分配给策略的名称。
	Name string `json:"name"`

	// 要创建的策略类型,service_control_policy服务控制策略；tag_policy：标签策略。
	Type CreatePolicyReqBodyType `json:"type"`

	// 要附加到新创建的策略的标签列表。
	Tags *[]TagDto `json:"tags,omitempty"`
}

func (o CreatePolicyReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePolicyReqBody struct{}"
	}

	return strings.Join([]string{"CreatePolicyReqBody", string(data)}, " ")
}

type CreatePolicyReqBodyType struct {
	value string
}

type CreatePolicyReqBodyTypeEnum struct {
	SERVICE_CONTROL_POLICY CreatePolicyReqBodyType
	TAG_POLICY             CreatePolicyReqBodyType
}

func GetCreatePolicyReqBodyTypeEnum() CreatePolicyReqBodyTypeEnum {
	return CreatePolicyReqBodyTypeEnum{
		SERVICE_CONTROL_POLICY: CreatePolicyReqBodyType{
			value: "service_control_policy",
		},
		TAG_POLICY: CreatePolicyReqBodyType{
			value: "tag_policy",
		},
	}
}

func (c CreatePolicyReqBodyType) Value() string {
	return c.value
}

func (c CreatePolicyReqBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePolicyReqBodyType) UnmarshalJSON(b []byte) error {
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
