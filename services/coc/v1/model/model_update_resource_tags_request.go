package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateResourceTagsRequest Request Object
type UpdateResourceTagsRequest struct {

	// 资源id。
	ResourceId string `json:"resource_id"`

	// 资源类型： - coc:script
	ResourceType UpdateResourceTagsRequestResourceType `json:"resource_type"`

	Body *UpdateResourceTagsRequestBody `json:"body,omitempty"`
}

func (o UpdateResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"UpdateResourceTagsRequest", string(data)}, " ")
}

type UpdateResourceTagsRequestResourceType struct {
	value string
}

type UpdateResourceTagsRequestResourceTypeEnum struct {
	COCSCRIPT UpdateResourceTagsRequestResourceType
}

func GetUpdateResourceTagsRequestResourceTypeEnum() UpdateResourceTagsRequestResourceTypeEnum {
	return UpdateResourceTagsRequestResourceTypeEnum{
		COCSCRIPT: UpdateResourceTagsRequestResourceType{
			value: "coc:script",
		},
	}
}

func (c UpdateResourceTagsRequestResourceType) Value() string {
	return c.value
}

func (c UpdateResourceTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateResourceTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
