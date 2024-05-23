package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TagResourceRequest Request Object
type TagResourceRequest struct {

	// 资源类型
	ResourceType TagResourceRequestResourceType `json:"resource_type"`

	// 资源ID
	ResourceId string `json:"resource_id"`

	Body *TagsReq `json:"body,omitempty"`
}

func (o TagResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagResourceRequest struct{}"
	}

	return strings.Join([]string{"TagResourceRequest", string(data)}, " ")
}

type TagResourceRequestResourceType struct {
	value string
}

type TagResourceRequestResourceTypeEnum struct {
	CONFIGPOLICY_ASSIGNMENTS TagResourceRequestResourceType
}

func GetTagResourceRequestResourceTypeEnum() TagResourceRequestResourceTypeEnum {
	return TagResourceRequestResourceTypeEnum{
		CONFIGPOLICY_ASSIGNMENTS: TagResourceRequestResourceType{
			value: "config:policyAssignments",
		},
	}
}

func (c TagResourceRequestResourceType) Value() string {
	return c.value
}

func (c TagResourceRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TagResourceRequestResourceType) UnmarshalJSON(b []byte) error {
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
