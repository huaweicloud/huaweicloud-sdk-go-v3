package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListCesTargetProjectTagsRequest Request Object
type ListCesTargetProjectTagsRequest struct {

	// **参数解释**： 资源类型。     **约束限制**： 不涉及。 **取值范围**： 只能为CES-alarm，代表告警规则。长度[1,32]个字符。      **默认取值**： 不涉及。
	ResourceType ListCesTargetProjectTagsRequestResourceType `json:"resource_type"`
}

func (o ListCesTargetProjectTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCesTargetProjectTagsRequest struct{}"
	}

	return strings.Join([]string{"ListCesTargetProjectTagsRequest", string(data)}, " ")
}

type ListCesTargetProjectTagsRequestResourceType struct {
	value string
}

type ListCesTargetProjectTagsRequestResourceTypeEnum struct {
	CES_ALARM ListCesTargetProjectTagsRequestResourceType
}

func GetListCesTargetProjectTagsRequestResourceTypeEnum() ListCesTargetProjectTagsRequestResourceTypeEnum {
	return ListCesTargetProjectTagsRequestResourceTypeEnum{
		CES_ALARM: ListCesTargetProjectTagsRequestResourceType{
			value: "CES-alarm",
		},
	}
}

func (c ListCesTargetProjectTagsRequestResourceType) Value() string {
	return c.value
}

func (c ListCesTargetProjectTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCesTargetProjectTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
