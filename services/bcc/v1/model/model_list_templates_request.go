package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTemplatesRequest Request Object
type ListTemplatesRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 模板类型，取值为系统预置或者自定义，取值范围：system,custom
	Type *ListTemplatesRequestType `json:"type,omitempty"`

	// 偏移量，默认取值0，最大取值100
	Offset *int64 `json:"offset,omitempty"`

	// 单次请求的条数，取值范围1到100，默认取值10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListTemplatesRequest", string(data)}, " ")
}

type ListTemplatesRequestType struct {
	value string
}

type ListTemplatesRequestTypeEnum struct {
	SYSTEM ListTemplatesRequestType
	CUSTOM ListTemplatesRequestType
}

func GetListTemplatesRequestTypeEnum() ListTemplatesRequestTypeEnum {
	return ListTemplatesRequestTypeEnum{
		SYSTEM: ListTemplatesRequestType{
			value: "system",
		},
		CUSTOM: ListTemplatesRequestType{
			value: "custom",
		},
	}
}

func (c ListTemplatesRequestType) Value() string {
	return c.value
}

func (c ListTemplatesRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTemplatesRequestType) UnmarshalJSON(b []byte) error {
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
