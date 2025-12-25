package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListComponentTemplateRequest Request Object
type ListComponentTemplateRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 组件ID
	ComponentId string `json:"component_id"`

	// **参数解释**: 文件类型 - JVM JVM配置文件类型 - LOG4J2 Log4j2日志配置文件类型 - YML YAML配置文件类型  **约束限制** 不涉及 **取值范围**: - JVM - LOG4J2 - YML   **默认值** 不涉及
	FileType *ListComponentTemplateRequestFileType `json:"file_type,omitempty"`

	// 每页数量
	Limit *int64 `json:"limit,omitempty"`

	// **参数解释：** 偏移量 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Offset *int32 `json:"offset,omitempty"`

	// 按照属性排序。
	SortKey *string `json:"sort_key,omitempty"`

	// 排序顺序，支持 `ASC` 或 `DESC`。
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListComponentTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComponentTemplateRequest struct{}"
	}

	return strings.Join([]string{"ListComponentTemplateRequest", string(data)}, " ")
}

type ListComponentTemplateRequestFileType struct {
	value string
}

type ListComponentTemplateRequestFileTypeEnum struct {
	JVM     ListComponentTemplateRequestFileType
	LOG4_J2 ListComponentTemplateRequestFileType
	YML     ListComponentTemplateRequestFileType
}

func GetListComponentTemplateRequestFileTypeEnum() ListComponentTemplateRequestFileTypeEnum {
	return ListComponentTemplateRequestFileTypeEnum{
		JVM: ListComponentTemplateRequestFileType{
			value: "JVM",
		},
		LOG4_J2: ListComponentTemplateRequestFileType{
			value: "LOG4J2",
		},
		YML: ListComponentTemplateRequestFileType{
			value: "YML",
		},
	}
}

func (c ListComponentTemplateRequestFileType) Value() string {
	return c.value
}

func (c ListComponentTemplateRequestFileType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListComponentTemplateRequestFileType) UnmarshalJSON(b []byte) error {
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
