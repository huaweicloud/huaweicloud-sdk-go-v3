package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchAddResourceTagRequest Request Object
type BatchAddResourceTagRequest struct {

	// 资源类型。 - auditInstance
	ResourceType BatchAddResourceTagRequestResourceType `json:"resource_type"`

	// 资源ID。可在查询实例列表接口的resource_id字段获取。
	ResourceId string `json:"resource_id"`

	Body *ResourceTagRequest `json:"body,omitempty"`
}

func (o BatchAddResourceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddResourceTagRequest struct{}"
	}

	return strings.Join([]string{"BatchAddResourceTagRequest", string(data)}, " ")
}

type BatchAddResourceTagRequestResourceType struct {
	value string
}

type BatchAddResourceTagRequestResourceTypeEnum struct {
	AUDIT_INSTANCE BatchAddResourceTagRequestResourceType
}

func GetBatchAddResourceTagRequestResourceTypeEnum() BatchAddResourceTagRequestResourceTypeEnum {
	return BatchAddResourceTagRequestResourceTypeEnum{
		AUDIT_INSTANCE: BatchAddResourceTagRequestResourceType{
			value: "auditInstance",
		},
	}
}

func (c BatchAddResourceTagRequestResourceType) Value() string {
	return c.value
}

func (c BatchAddResourceTagRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchAddResourceTagRequestResourceType) UnmarshalJSON(b []byte) error {
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
