package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListCollectorParsersRequest Request Object
type ListCollectorParsersRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**: 查询类型 - QUICK 快速 - GENERAL 常规  **约束限制** 不涉及 **取值范围**: - QUICK - GENERAL  **默认值** 不涉及
	QueryType *ListCollectorParsersRequestQueryType `json:"query_type,omitempty"`

	// 标题
	Title *string `json:"title,omitempty"`

	// 相关描述
	Description *string `json:"description,omitempty"`

	// 第几页
	Offset *int64 `json:"offset,omitempty"`

	// 每页数量
	Limit *int64 `json:"limit,omitempty"`

	// 排序字段
	SortKey *int64 `json:"sort_key,omitempty"`

	// **参数解释**: 目录排序 - asc 升序排列 - desc 降序排列  **约束限制** 不涉及 **取值范围**: - asc - desc  **默认值** 不涉及
	SortDir *ListCollectorParsersRequestSortDir `json:"sort_dir,omitempty"`
}

func (o ListCollectorParsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollectorParsersRequest struct{}"
	}

	return strings.Join([]string{"ListCollectorParsersRequest", string(data)}, " ")
}

type ListCollectorParsersRequestQueryType struct {
	value string
}

type ListCollectorParsersRequestQueryTypeEnum struct {
	QUICK   ListCollectorParsersRequestQueryType
	GENERAL ListCollectorParsersRequestQueryType
}

func GetListCollectorParsersRequestQueryTypeEnum() ListCollectorParsersRequestQueryTypeEnum {
	return ListCollectorParsersRequestQueryTypeEnum{
		QUICK: ListCollectorParsersRequestQueryType{
			value: "QUICK",
		},
		GENERAL: ListCollectorParsersRequestQueryType{
			value: "GENERAL",
		},
	}
}

func (c ListCollectorParsersRequestQueryType) Value() string {
	return c.value
}

func (c ListCollectorParsersRequestQueryType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCollectorParsersRequestQueryType) UnmarshalJSON(b []byte) error {
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

type ListCollectorParsersRequestSortDir struct {
	value string
}

type ListCollectorParsersRequestSortDirEnum struct {
	ASC  ListCollectorParsersRequestSortDir
	DESC ListCollectorParsersRequestSortDir
}

func GetListCollectorParsersRequestSortDirEnum() ListCollectorParsersRequestSortDirEnum {
	return ListCollectorParsersRequestSortDirEnum{
		ASC: ListCollectorParsersRequestSortDir{
			value: "asc",
		},
		DESC: ListCollectorParsersRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListCollectorParsersRequestSortDir) Value() string {
	return c.value
}

func (c ListCollectorParsersRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCollectorParsersRequestSortDir) UnmarshalJSON(b []byte) error {
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
