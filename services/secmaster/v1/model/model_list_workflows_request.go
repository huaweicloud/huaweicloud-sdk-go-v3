package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListWorkflowsRequest Request Object
type ListWorkflowsRequest struct {

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"content-type"`

	// **参数解释**: 工作空间ID **取值范围**: 不涉及
	WorkspaceId string `json:"workspace_id"`

	// **参数解释：** 偏移量 **约束限制：** 0-10000 **取值范围：** 不涉及 **默认取值：** 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 数据量 **约束限制：** 1-100 **取值范围：** 不涉及 **默认取值：** 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 排序顺序 - ASC：升序 - DESC：降序  **约束限制：** 不涉及 **取值范围：** - ASC：升序 - DESC：降序  **默认取值：** DESC
	Order *ListWorkflowsRequestOrder `json:"order,omitempty"`

	// **参数解释：** 排序字段， - create_time：创建时间 - update_time：更新时间  **约束限制：** 不涉及 **取值范围：** - create_time - update_time  **默认取值：** create_time
	Sortby *ListWorkflowsRequestSortby `json:"sortby,omitempty"`

	// **参数解释**: 是否已启用 **约束限制**: 不涉及         **取值范围**: - true  已启用 - false 未启用  **默认值**:  不涉及
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释**: 流程名称 **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**: 流程描述 **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**: 数据类的ID **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	DataclassId *string `json:"dataclass_id,omitempty"`

	// **参数解释**: 数据类名称 **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	DataclassName *string `json:"dataclass_name,omitempty"`

	// **参数解释**: 流程的类型 - NORMAL 通用 - SURVEY 调查 - HEMOSTASIS 止血 - EASE 缓解  **约束限制**: 不涉及         **取值范围**: - NORMAL - SURVEY - HEMOSTASIS - EASE  **默认值**:  不涉及
	AopType *ListWorkflowsRequestAopType `json:"aop_type,omitempty"`
}

func (o ListWorkflowsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkflowsRequest struct{}"
	}

	return strings.Join([]string{"ListWorkflowsRequest", string(data)}, " ")
}

type ListWorkflowsRequestOrder struct {
	value string
}

type ListWorkflowsRequestOrderEnum struct {
	ASC  ListWorkflowsRequestOrder
	DESC ListWorkflowsRequestOrder
}

func GetListWorkflowsRequestOrderEnum() ListWorkflowsRequestOrderEnum {
	return ListWorkflowsRequestOrderEnum{
		ASC: ListWorkflowsRequestOrder{
			value: "ASC",
		},
		DESC: ListWorkflowsRequestOrder{
			value: "DESC",
		},
	}
}

func (c ListWorkflowsRequestOrder) Value() string {
	return c.value
}

func (c ListWorkflowsRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListWorkflowsRequestOrder) UnmarshalJSON(b []byte) error {
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

type ListWorkflowsRequestSortby struct {
	value string
}

type ListWorkflowsRequestSortbyEnum struct {
	UPDATE_TIME ListWorkflowsRequestSortby
	CREATE_TIME ListWorkflowsRequestSortby
}

func GetListWorkflowsRequestSortbyEnum() ListWorkflowsRequestSortbyEnum {
	return ListWorkflowsRequestSortbyEnum{
		UPDATE_TIME: ListWorkflowsRequestSortby{
			value: "update_time",
		},
		CREATE_TIME: ListWorkflowsRequestSortby{
			value: "create_time",
		},
	}
}

func (c ListWorkflowsRequestSortby) Value() string {
	return c.value
}

func (c ListWorkflowsRequestSortby) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListWorkflowsRequestSortby) UnmarshalJSON(b []byte) error {
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

type ListWorkflowsRequestAopType struct {
	value string
}

type ListWorkflowsRequestAopTypeEnum struct {
	NORMAL     ListWorkflowsRequestAopType
	SURVEY     ListWorkflowsRequestAopType
	HEMOSTASIS ListWorkflowsRequestAopType
	EASE       ListWorkflowsRequestAopType
}

func GetListWorkflowsRequestAopTypeEnum() ListWorkflowsRequestAopTypeEnum {
	return ListWorkflowsRequestAopTypeEnum{
		NORMAL: ListWorkflowsRequestAopType{
			value: "NORMAL",
		},
		SURVEY: ListWorkflowsRequestAopType{
			value: "SURVEY",
		},
		HEMOSTASIS: ListWorkflowsRequestAopType{
			value: "HEMOSTASIS",
		},
		EASE: ListWorkflowsRequestAopType{
			value: "EASE",
		},
	}
}

func (c ListWorkflowsRequestAopType) Value() string {
	return c.value
}

func (c ListWorkflowsRequestAopType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListWorkflowsRequestAopType) UnmarshalJSON(b []byte) error {
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
