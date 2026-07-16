package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListFlavorsRequest Request Object
type ListFlavorsRequest struct {

	// **参数解释**：规格处理器类型。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： - CPU - GPU - [ASCEND](tag:hc,hk,fcs_super)  **默认取值**：不涉及。
	Category *ListFlavorsRequestCategory `json:"category,omitempty"`

	// **参数解释**：每一页显示的有效规格数量，默认不限制。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：分页记录的起始位置偏移量。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：集群类型。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： - MANAGED：公共集群 - DEDICATED：专属集群  **默认取值**：不涉及。
	Type *ListFlavorsRequestType `json:"type,omitempty"`

	// **参数解释**：排序方式。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： - ASC：升序 - DESC：降序  **默认取值**：DESC。
	SortDir *ListFlavorsRequestSortDir `json:"sort_dir,omitempty"`

	// **参数解释**：排序的字段，多个字段使用(“,”)逗号分隔。 **约束限制**：不涉及。 **取值范围**：长度限制为128个字符，支持大小写字母、数字、中划线、下划线和逗号。 **默认取值**：不涉及。
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**：资源类型 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： -ASCEND_SNT9：昇腾910芯片。 -ASCEND_SNT9B：昇腾910B芯片。 -ASCEND_SNT3：昇腾310芯片。  **默认取值**：不涉及。
	FlavorType *string `json:"flavor_type,omitempty"`

	// **参数解释**：特性名称。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： - NOTEBOOK：用户显式创建的Notebook实例。  **默认取值**：NOTEBOOK。
	Feature *string `json:"feature,omitempty"`
}

func (o ListFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ListFlavorsRequest", string(data)}, " ")
}

type ListFlavorsRequestCategory struct {
	value string
}

type ListFlavorsRequestCategoryEnum struct {
	ASCEND ListFlavorsRequestCategory
	CPU    ListFlavorsRequestCategory
	GPU    ListFlavorsRequestCategory
}

func GetListFlavorsRequestCategoryEnum() ListFlavorsRequestCategoryEnum {
	return ListFlavorsRequestCategoryEnum{
		ASCEND: ListFlavorsRequestCategory{
			value: "ASCEND",
		},
		CPU: ListFlavorsRequestCategory{
			value: "CPU",
		},
		GPU: ListFlavorsRequestCategory{
			value: "GPU",
		},
	}
}

func (c ListFlavorsRequestCategory) Value() string {
	return c.value
}

func (c ListFlavorsRequestCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlavorsRequestCategory) UnmarshalJSON(b []byte) error {
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

type ListFlavorsRequestType struct {
	value string
}

type ListFlavorsRequestTypeEnum struct {
	DEDICATED      ListFlavorsRequestType
	DEDICATED_ROMA ListFlavorsRequestType
	MANAGED        ListFlavorsRequestType
	MANAGED_ROMA   ListFlavorsRequestType
}

func GetListFlavorsRequestTypeEnum() ListFlavorsRequestTypeEnum {
	return ListFlavorsRequestTypeEnum{
		DEDICATED: ListFlavorsRequestType{
			value: "DEDICATED",
		},
		DEDICATED_ROMA: ListFlavorsRequestType{
			value: "DEDICATED_ROMA",
		},
		MANAGED: ListFlavorsRequestType{
			value: "MANAGED",
		},
		MANAGED_ROMA: ListFlavorsRequestType{
			value: "MANAGED_ROMA",
		},
	}
}

func (c ListFlavorsRequestType) Value() string {
	return c.value
}

func (c ListFlavorsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlavorsRequestType) UnmarshalJSON(b []byte) error {
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

type ListFlavorsRequestSortDir struct {
	value string
}

type ListFlavorsRequestSortDirEnum struct {
	ASC  ListFlavorsRequestSortDir
	DESC ListFlavorsRequestSortDir
}

func GetListFlavorsRequestSortDirEnum() ListFlavorsRequestSortDirEnum {
	return ListFlavorsRequestSortDirEnum{
		ASC: ListFlavorsRequestSortDir{
			value: "ASC",
		},
		DESC: ListFlavorsRequestSortDir{
			value: "DESC",
		},
	}
}

func (c ListFlavorsRequestSortDir) Value() string {
	return c.value
}

func (c ListFlavorsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlavorsRequestSortDir) UnmarshalJSON(b []byte) error {
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
