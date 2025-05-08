package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListScriptResourceTagsRequest Request Object
type ListScriptResourceTagsRequest struct {

	// 资源类型： - coc:script
	ResourceType ListScriptResourceTagsRequestResourceType `json:"resource_type"`

	// 查询的限制数量。
	Limit *string `json:"limit,omitempty"`

	// 查询的起始位置。
	Offset *string `json:"offset,omitempty"`
}

func (o ListScriptResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScriptResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"ListScriptResourceTagsRequest", string(data)}, " ")
}

type ListScriptResourceTagsRequestResourceType struct {
	value string
}

type ListScriptResourceTagsRequestResourceTypeEnum struct {
	COCSCRIPT ListScriptResourceTagsRequestResourceType
}

func GetListScriptResourceTagsRequestResourceTypeEnum() ListScriptResourceTagsRequestResourceTypeEnum {
	return ListScriptResourceTagsRequestResourceTypeEnum{
		COCSCRIPT: ListScriptResourceTagsRequestResourceType{
			value: "coc:script",
		},
	}
}

func (c ListScriptResourceTagsRequestResourceType) Value() string {
	return c.value
}

func (c ListScriptResourceTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListScriptResourceTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
