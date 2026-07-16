package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListImageGroupRequest Request Object
type ListImageGroupRequest struct {

	// **参数解释**：镜像名称。 **约束限制**：不涉及。 **取值范围**：长度限制为512个字符，支持小写字母、数字、中划线、下划线和点。 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：镜像名称是否模糊匹配查询。 **约束限制**：不涉及。 **取值范围**：布尔类型： - true：支持模糊匹配查询。 - false：不支持模糊匹配查询。  **默认取值**：true。
	NameFuzzyMatch *bool `json:"name_fuzzy_match,omitempty"`

	// **参数解释**：镜像所属组织，可以在SWR控制台“组织管理”创建和查看。 **约束限制**：不涉及。 **取值范围**：长度限制为64个字符，支持大小写字母、数字、中划线、下划线和点号，且必须是小写字母开头。 **默认取值**：不涉及。
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**：镜像类型。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： - BUILD_IN：系统内置镜像。 - DEDICATED：用户保存的镜像。  **默认取值**：不涉及。
	Type *ListImageGroupRequestType `json:"type,omitempty"`

	// **参数解释**：工作空间ID。[获取方法请参见[查询工作空间列表](ListWorkspace.xml)。](tag:hc)未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。 **约束限制**：不涉及。 **取值范围**：0或32位仅包含字符0-9或小写字母a-z的字符串。 **默认取值**：0。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**：每一页显示的镜像实例数量。 **约束限制**：不涉及。 **取值范围**：正整数。 **默认取值**：200。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：分页记录的起始位置偏移量。 **约束限制**：不涉及。 **取值范围**：非负整数。 **默认取值**：0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：企业版SWR仓库ID。 **参数约束**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	SwrInstanceId *string `json:"swr_instance_id,omitempty"`
}

func (o ListImageGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageGroupRequest struct{}"
	}

	return strings.Join([]string{"ListImageGroupRequest", string(data)}, " ")
}

type ListImageGroupRequestType struct {
	value string
}

type ListImageGroupRequestTypeEnum struct {
	BUILD_IN  ListImageGroupRequestType
	DEDICATED ListImageGroupRequestType
}

func GetListImageGroupRequestTypeEnum() ListImageGroupRequestTypeEnum {
	return ListImageGroupRequestTypeEnum{
		BUILD_IN: ListImageGroupRequestType{
			value: "BUILD_IN",
		},
		DEDICATED: ListImageGroupRequestType{
			value: "DEDICATED",
		},
	}
}

func (c ListImageGroupRequestType) Value() string {
	return c.value
}

func (c ListImageGroupRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImageGroupRequestType) UnmarshalJSON(b []byte) error {
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
