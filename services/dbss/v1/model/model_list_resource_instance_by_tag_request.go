package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListResourceInstanceByTagRequest Request Object
type ListResourceInstanceByTagRequest struct {

	// **参数解释**：  资源类型。 **约束限制**： 不涉及 **取值范围**：  - auditInstance: 审计  **默认取值**： 不涉及
	ResourceType ListResourceInstanceByTagRequestResourceType `json:"resource_type"`

	// **参数解释**： 每页查询记录数。 **约束限制**： 仅支持大于0的整数 **取值范围**： 大于0小于等于10000 **默认取值**： 默认值为1000
	Limit *string `json:"limit,omitempty"`

	// **参数解释**： 分页偏移量，从第一条数据偏移offset条数据后开始查询 **约束限制**： 仅支持大于等于0的整数 **取值范围**： 大于等于0 **默认取值**： 默认值为0
	Offset *string `json:"offset,omitempty"`

	Body *ResourceInstanceTagRequest `json:"body,omitempty"`
}

func (o ListResourceInstanceByTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceInstanceByTagRequest struct{}"
	}

	return strings.Join([]string{"ListResourceInstanceByTagRequest", string(data)}, " ")
}

type ListResourceInstanceByTagRequestResourceType struct {
	value string
}

type ListResourceInstanceByTagRequestResourceTypeEnum struct {
	AUDIT_INSTANCE ListResourceInstanceByTagRequestResourceType
}

func GetListResourceInstanceByTagRequestResourceTypeEnum() ListResourceInstanceByTagRequestResourceTypeEnum {
	return ListResourceInstanceByTagRequestResourceTypeEnum{
		AUDIT_INSTANCE: ListResourceInstanceByTagRequestResourceType{
			value: "auditInstance",
		},
	}
}

func (c ListResourceInstanceByTagRequestResourceType) Value() string {
	return c.value
}

func (c ListResourceInstanceByTagRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourceInstanceByTagRequestResourceType) UnmarshalJSON(b []byte) error {
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
