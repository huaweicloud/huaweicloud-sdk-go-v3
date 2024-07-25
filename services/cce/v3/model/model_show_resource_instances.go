package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowResourceInstances
type ShowResourceInstances struct {

	// action类型，\"filter\"或者\"count\"。
	Action ShowResourceInstancesAction `json:"action"`

	Tags *TagFilter `json:"tags,omitempty"`

	TagsAny *TagFilter `json:"tags_any,omitempty"`

	NotTags *TagFilter `json:"not_tags,omitempty"`

	NotTagsAny *TagFilter `json:"not_tags_any,omitempty"`

	SysTags *TagFilter `json:"sys_tags,omitempty"`

	// 忽略其他标签字段，返回不带任何标签的资源。
	WithoutAnyTag *bool `json:"without_any_tag,omitempty"`

	Limit *string `json:"limit,omitempty"`

	Offset *string `json:"offset,omitempty"`

	Matches *[]interface{} `json:"matches,omitempty"`
}

func (o ShowResourceInstances) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceInstances struct{}"
	}

	return strings.Join([]string{"ShowResourceInstances", string(data)}, " ")
}

type ShowResourceInstancesAction struct {
	value string
}

type ShowResourceInstancesActionEnum struct {
	FILTER ShowResourceInstancesAction
	COUNT  ShowResourceInstancesAction
}

func GetShowResourceInstancesActionEnum() ShowResourceInstancesActionEnum {
	return ShowResourceInstancesActionEnum{
		FILTER: ShowResourceInstancesAction{
			value: "filter",
		},
		COUNT: ShowResourceInstancesAction{
			value: "count",
		},
	}
}

func (c ShowResourceInstancesAction) Value() string {
	return c.value
}

func (c ShowResourceInstancesAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowResourceInstancesAction) UnmarshalJSON(b []byte) error {
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
