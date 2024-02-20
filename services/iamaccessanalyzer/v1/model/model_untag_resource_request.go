package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UntagResourceRequest Request Object
type UntagResourceRequest struct {

	// 资源类型。
	ResourceType UntagResourceRequestResourceType `json:"resource_type"`

	// 资源的唯一标识符。
	ResourceId string `json:"resource_id"`

	Body *UntagResourceReqBody `json:"body,omitempty"`
}

func (o UntagResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UntagResourceRequest struct{}"
	}

	return strings.Join([]string{"UntagResourceRequest", string(data)}, " ")
}

type UntagResourceRequestResourceType struct {
	value string
}

type UntagResourceRequestResourceTypeEnum struct {
	ANALYZERS UntagResourceRequestResourceType
}

func GetUntagResourceRequestResourceTypeEnum() UntagResourceRequestResourceTypeEnum {
	return UntagResourceRequestResourceTypeEnum{
		ANALYZERS: UntagResourceRequestResourceType{
			value: "analyzers",
		},
	}
}

func (c UntagResourceRequestResourceType) Value() string {
	return c.value
}

func (c UntagResourceRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UntagResourceRequestResourceType) UnmarshalJSON(b []byte) error {
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
