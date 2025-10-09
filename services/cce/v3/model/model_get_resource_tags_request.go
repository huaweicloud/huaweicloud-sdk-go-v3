package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// GetResourceTagsRequest Request Object
type GetResourceTagsRequest struct {

	// **参数解释**： 资源类型 **约束限制：** 不涉及 **取值范围：** - cce-cluster：集群  **默认取值：** 不涉及
	ResourceType GetResourceTagsRequestResourceType `json:"resource_type"`

	// **参数解释**： 资源id。例：集群id，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ResourceId string `json:"resource_id"`
}

func (o GetResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"GetResourceTagsRequest", string(data)}, " ")
}

type GetResourceTagsRequestResourceType struct {
	value string
}

type GetResourceTagsRequestResourceTypeEnum struct {
	CCE_CLUSTER GetResourceTagsRequestResourceType
}

func GetGetResourceTagsRequestResourceTypeEnum() GetResourceTagsRequestResourceTypeEnum {
	return GetResourceTagsRequestResourceTypeEnum{
		CCE_CLUSTER: GetResourceTagsRequestResourceType{
			value: "cce-cluster",
		},
	}
}

func (c GetResourceTagsRequestResourceType) Value() string {
	return c.value
}

func (c GetResourceTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetResourceTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
