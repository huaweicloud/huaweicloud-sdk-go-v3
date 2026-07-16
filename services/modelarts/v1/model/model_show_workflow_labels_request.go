package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowWorkflowLabelsRequest Request Object
type ShowWorkflowLabelsRequest struct {

	// 返回的数据条目数。
	Limit *int32 `json:"limit,omitempty"`

	// 数据条目偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// instance order
	Order *ShowWorkflowLabelsRequestOrder `json:"order,omitempty"`

	// 指定排序字段。  可选值： - user_name：IAM用户名称 - create_time：创建时间
	SortBy *ShowWorkflowLabelsRequestSortBy `json:"sort_by,omitempty"`

	// **参数解释**：工作流模板ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	TemplateId *string `json:"template_id,omitempty"`
}

func (o ShowWorkflowLabelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkflowLabelsRequest struct{}"
	}

	return strings.Join([]string{"ShowWorkflowLabelsRequest", string(data)}, " ")
}

type ShowWorkflowLabelsRequestOrder struct {
	value string
}

type ShowWorkflowLabelsRequestOrderEnum struct {
	ASC  ShowWorkflowLabelsRequestOrder
	DESC ShowWorkflowLabelsRequestOrder
}

func GetShowWorkflowLabelsRequestOrderEnum() ShowWorkflowLabelsRequestOrderEnum {
	return ShowWorkflowLabelsRequestOrderEnum{
		ASC: ShowWorkflowLabelsRequestOrder{
			value: "asc",
		},
		DESC: ShowWorkflowLabelsRequestOrder{
			value: "desc",
		},
	}
}

func (c ShowWorkflowLabelsRequestOrder) Value() string {
	return c.value
}

func (c ShowWorkflowLabelsRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowWorkflowLabelsRequestOrder) UnmarshalJSON(b []byte) error {
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

type ShowWorkflowLabelsRequestSortBy struct {
	value string
}

type ShowWorkflowLabelsRequestSortByEnum struct {
	USER_NAME   ShowWorkflowLabelsRequestSortBy
	CREATE_TIME ShowWorkflowLabelsRequestSortBy
}

func GetShowWorkflowLabelsRequestSortByEnum() ShowWorkflowLabelsRequestSortByEnum {
	return ShowWorkflowLabelsRequestSortByEnum{
		USER_NAME: ShowWorkflowLabelsRequestSortBy{
			value: "user_name",
		},
		CREATE_TIME: ShowWorkflowLabelsRequestSortBy{
			value: "create_time",
		},
	}
}

func (c ShowWorkflowLabelsRequestSortBy) Value() string {
	return c.value
}

func (c ShowWorkflowLabelsRequestSortBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowWorkflowLabelsRequestSortBy) UnmarshalJSON(b []byte) error {
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
