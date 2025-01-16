package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPrrTemplateRequest Request Object
type ListPrrTemplateRequest struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// 应用类型 core: 核心应用 non-core: 非核心应用
	ApplicationType *ListPrrTemplateRequestApplicationType `json:"application_type,omitempty"`

	// 分页参数
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListPrrTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrrTemplateRequest struct{}"
	}

	return strings.Join([]string{"ListPrrTemplateRequest", string(data)}, " ")
}

type ListPrrTemplateRequestApplicationType struct {
	value string
}

type ListPrrTemplateRequestApplicationTypeEnum struct {
	CORE     ListPrrTemplateRequestApplicationType
	NON_CORE ListPrrTemplateRequestApplicationType
}

func GetListPrrTemplateRequestApplicationTypeEnum() ListPrrTemplateRequestApplicationTypeEnum {
	return ListPrrTemplateRequestApplicationTypeEnum{
		CORE: ListPrrTemplateRequestApplicationType{
			value: "core",
		},
		NON_CORE: ListPrrTemplateRequestApplicationType{
			value: "non-core",
		},
	}
}

func (c ListPrrTemplateRequestApplicationType) Value() string {
	return c.value
}

func (c ListPrrTemplateRequestApplicationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPrrTemplateRequestApplicationType) UnmarshalJSON(b []byte) error {
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
