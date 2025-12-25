package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListCollectorParserTemplateRequest Request Object
type ListCollectorParserTemplateRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 解析器标题
	Title *string `json:"title,omitempty"`

	// 解析器相关描述
	Description *string `json:"description,omitempty"`

	// 第几页
	Offset *int64 `json:"offset,omitempty"`

	// 每页数量
	Limit *int64 `json:"limit,omitempty"`

	// 排序字段
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**: 目录排序 - asc 升序排列 - desc 降序排列  **约束限制** 不涉及 **取值范围**: - asc - desc  **默认值** 不涉及
	SortDir *ListCollectorParserTemplateRequestSortDir `json:"sort_dir,omitempty"`
}

func (o ListCollectorParserTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollectorParserTemplateRequest struct{}"
	}

	return strings.Join([]string{"ListCollectorParserTemplateRequest", string(data)}, " ")
}

type ListCollectorParserTemplateRequestSortDir struct {
	value string
}

type ListCollectorParserTemplateRequestSortDirEnum struct {
	ASC  ListCollectorParserTemplateRequestSortDir
	DESC ListCollectorParserTemplateRequestSortDir
}

func GetListCollectorParserTemplateRequestSortDirEnum() ListCollectorParserTemplateRequestSortDirEnum {
	return ListCollectorParserTemplateRequestSortDirEnum{
		ASC: ListCollectorParserTemplateRequestSortDir{
			value: "asc",
		},
		DESC: ListCollectorParserTemplateRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListCollectorParserTemplateRequestSortDir) Value() string {
	return c.value
}

func (c ListCollectorParserTemplateRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCollectorParserTemplateRequestSortDir) UnmarshalJSON(b []byte) error {
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
