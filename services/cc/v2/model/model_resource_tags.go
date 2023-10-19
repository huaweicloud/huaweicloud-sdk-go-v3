package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ResourceTags 资源标签Tag汇总复数
type ResourceTags struct {

	// 动作。|- create：创建。 delete：删除。
	Action *ResourceTagsAction `json:"action,omitempty"`

	// 批量添加/删除资源标签
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o ResourceTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTags struct{}"
	}

	return strings.Join([]string{"ResourceTags", string(data)}, " ")
}

type ResourceTagsAction struct {
	value string
}

type ResourceTagsActionEnum struct {
	CREATE ResourceTagsAction
	DELETE ResourceTagsAction
}

func GetResourceTagsActionEnum() ResourceTagsActionEnum {
	return ResourceTagsActionEnum{
		CREATE: ResourceTagsAction{
			value: "create",
		},
		DELETE: ResourceTagsAction{
			value: "delete",
		},
	}
}

func (c ResourceTagsAction) Value() string {
	return c.value
}

func (c ResourceTagsAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceTagsAction) UnmarshalJSON(b []byte) error {
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
