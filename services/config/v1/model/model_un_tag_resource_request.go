package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UnTagResourceRequest Request Object
type UnTagResourceRequest struct {

	// 资源类型
	ResourceType UnTagResourceRequestResourceType `json:"resource_type"`

	// 资源ID
	ResourceId string `json:"resource_id"`

	Body *UnTagsReq `json:"body,omitempty"`
}

func (o UnTagResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnTagResourceRequest struct{}"
	}

	return strings.Join([]string{"UnTagResourceRequest", string(data)}, " ")
}

type UnTagResourceRequestResourceType struct {
	value string
}

type UnTagResourceRequestResourceTypeEnum struct {
	CONFIGPOLICY_ASSIGNMENTS UnTagResourceRequestResourceType
}

func GetUnTagResourceRequestResourceTypeEnum() UnTagResourceRequestResourceTypeEnum {
	return UnTagResourceRequestResourceTypeEnum{
		CONFIGPOLICY_ASSIGNMENTS: UnTagResourceRequestResourceType{
			value: "config:policyAssignments",
		},
	}
}

func (c UnTagResourceRequestResourceType) Value() string {
	return c.value
}

func (c UnTagResourceRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UnTagResourceRequestResourceType) UnmarshalJSON(b []byte) error {
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
