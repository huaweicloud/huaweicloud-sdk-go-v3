package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchDeleteResourceTagRequest Request Object
type BatchDeleteResourceTagRequest struct {

	// **参数解释**：  资源类型。 **约束限制**： 不涉及 **取值范围**：  - auditInstance: 审计  **默认取值**： 不涉及
	ResourceType BatchDeleteResourceTagRequestResourceType `json:"resource_type"`

	// **参数解释**：  资源ID。可在查询实例列表接口的resource_id字段获取。 **约束限制**： 不涉及 **取值范围**： 以查询实例列表接口获取值为准，字符长度32-64。 **默认取值**： 不涉及
	ResourceId string `json:"resource_id"`

	Body *ResourceTagDeleteRequest `json:"body,omitempty"`
}

func (o BatchDeleteResourceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteResourceTagRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteResourceTagRequest", string(data)}, " ")
}

type BatchDeleteResourceTagRequestResourceType struct {
	value string
}

type BatchDeleteResourceTagRequestResourceTypeEnum struct {
	AUDIT_INSTANCE BatchDeleteResourceTagRequestResourceType
}

func GetBatchDeleteResourceTagRequestResourceTypeEnum() BatchDeleteResourceTagRequestResourceTypeEnum {
	return BatchDeleteResourceTagRequestResourceTypeEnum{
		AUDIT_INSTANCE: BatchDeleteResourceTagRequestResourceType{
			value: "auditInstance",
		},
	}
}

func (c BatchDeleteResourceTagRequestResourceType) Value() string {
	return c.value
}

func (c BatchDeleteResourceTagRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteResourceTagRequestResourceType) UnmarshalJSON(b []byte) error {
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
