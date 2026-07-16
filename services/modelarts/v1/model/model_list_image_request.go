package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListImageRequest Request Object
type ListImageRequest struct {

	// **参数解释**：每一页显示的镜像实例数量。 **约束限制**：不涉及。 **取值范围**：正整数。 **默认取值**：200。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：镜像名称。 **约束限制**：不涉及。 **取值范围**：长度限制为512个字符，支持小写字母、数字、中划线、下划线和点。 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：镜像名称是否模糊匹配查询。 **约束限制**：不涉及。 **取值范围**：布尔类型： - true：支持模糊匹配查询。 - false：不支持模糊匹配查询。  **默认取值**：true。
	NameFuzzyMatch *bool `json:"name_fuzzy_match,omitempty"`

	// **参数解释**：镜像所属组织，可以在SWR控制台“组织管理”创建和查看。 **约束限制**：不涉及。 **取值范围**：长度限制为64个字符，支持大小写字母、数字、中划线、下划线和点号，且必须是小写字母开头。 **默认取值**：不涉及。
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**：分页记录的起始位置偏移量。 **约束限制**：不涉及。 **取值范围**：非负整数。 **默认取值**：0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：镜像支持服务类型。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： - COMMON：通用镜像。 - INFERENCE: 建议仅在推理部署场景使用。 - TRAIN: 建议仅在训练任务场景使用。 - DEV: 建议仅在开发调测场景使用。 - UNKNOWN: 未明确设置的镜像支持的服务类型。  **默认取值**：UNKNOWN。
	ServiceType *ListImageRequestServiceType `json:"service_type,omitempty"`

	// **参数解释**：实例排序方式。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： - ASC：升序 - DESC：降序  **默认取值**：DESC。
	SortDir *ListImageRequestSortDir `json:"sort_dir,omitempty"`

	// **参数解释**：排序的字段，多个字段使用(“,”)逗号分隔。 **约束限制**：不涉及。 **取值范围**：长度限制为128个字符，支持大小写字母、数字、中划线、下划线和逗号。 **默认取值**：不涉及。
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**：镜像类型。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： - BUILD_IN：系统内置镜像。 - DEDICATED：用户保存的镜像。  **默认取值**：BUILD_IN。
	Type *ListImageRequestType `json:"type,omitempty"`

	// **参数解释**：工作空间ID。未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。 **约束限制**：不涉及。 **取值范围**：0或32位仅包含字符0-9或小写字母a-z的字符串。 **默认取值**：0。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**：镜像展示name。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ShowName *string `json:"show_name,omitempty"`

	// **参数解释**：镜像展示Tag。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ShowTag *string `json:"show_tag,omitempty"`
}

func (o ListImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageRequest struct{}"
	}

	return strings.Join([]string{"ListImageRequest", string(data)}, " ")
}

type ListImageRequestServiceType struct {
	value string
}

type ListImageRequestServiceTypeEnum struct {
	COMMON    ListImageRequestServiceType
	DEV       ListImageRequestServiceType
	INFERENCE ListImageRequestServiceType
	TRAIN     ListImageRequestServiceType
	UNKNOWN   ListImageRequestServiceType
}

func GetListImageRequestServiceTypeEnum() ListImageRequestServiceTypeEnum {
	return ListImageRequestServiceTypeEnum{
		COMMON: ListImageRequestServiceType{
			value: "COMMON",
		},
		DEV: ListImageRequestServiceType{
			value: "DEV",
		},
		INFERENCE: ListImageRequestServiceType{
			value: "INFERENCE",
		},
		TRAIN: ListImageRequestServiceType{
			value: "TRAIN",
		},
		UNKNOWN: ListImageRequestServiceType{
			value: "UNKNOWN",
		},
	}
}

func (c ListImageRequestServiceType) Value() string {
	return c.value
}

func (c ListImageRequestServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImageRequestServiceType) UnmarshalJSON(b []byte) error {
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

type ListImageRequestSortDir struct {
	value string
}

type ListImageRequestSortDirEnum struct {
	ASC  ListImageRequestSortDir
	DESC ListImageRequestSortDir
}

func GetListImageRequestSortDirEnum() ListImageRequestSortDirEnum {
	return ListImageRequestSortDirEnum{
		ASC: ListImageRequestSortDir{
			value: "ASC",
		},
		DESC: ListImageRequestSortDir{
			value: "DESC",
		},
	}
}

func (c ListImageRequestSortDir) Value() string {
	return c.value
}

func (c ListImageRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImageRequestSortDir) UnmarshalJSON(b []byte) error {
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

type ListImageRequestType struct {
	value string
}

type ListImageRequestTypeEnum struct {
	BUILD_IN  ListImageRequestType
	DEDICATED ListImageRequestType
}

func GetListImageRequestTypeEnum() ListImageRequestTypeEnum {
	return ListImageRequestTypeEnum{
		BUILD_IN: ListImageRequestType{
			value: "BUILD_IN",
		},
		DEDICATED: ListImageRequestType{
			value: "DEDICATED",
		},
	}
}

func (c ListImageRequestType) Value() string {
	return c.value
}

func (c ListImageRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImageRequestType) UnmarshalJSON(b []byte) error {
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
