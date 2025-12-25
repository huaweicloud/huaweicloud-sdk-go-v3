package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListCollectorChannelsRequest Request Object
type ListCollectorChannelsRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 相关标题
	Title *string `json:"title,omitempty"`

	// 链接模块ID
	ConnectionModuleId *string `json:"connection_module_id,omitempty"`

	// 解析器ID
	ParserId *string `json:"parser_id,omitempty"`

	// 组ID
	GroupId *string `json:"group_id,omitempty"`

	// 第几页
	Offset *int32 `json:"offset,omitempty"`

	// 每页数量
	Limit *int32 `json:"limit,omitempty"`

	// 排序字段
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**: 目录排序 - asc 升序排列 - desc 降序排列  **约束限制** 不涉及 **取值范围**: - asc - desc  **默认值** 不涉及
	SortDir *ListCollectorChannelsRequestSortDir `json:"sort_dir,omitempty"`
}

func (o ListCollectorChannelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollectorChannelsRequest struct{}"
	}

	return strings.Join([]string{"ListCollectorChannelsRequest", string(data)}, " ")
}

type ListCollectorChannelsRequestSortDir struct {
	value string
}

type ListCollectorChannelsRequestSortDirEnum struct {
	ASC  ListCollectorChannelsRequestSortDir
	DESC ListCollectorChannelsRequestSortDir
}

func GetListCollectorChannelsRequestSortDirEnum() ListCollectorChannelsRequestSortDirEnum {
	return ListCollectorChannelsRequestSortDirEnum{
		ASC: ListCollectorChannelsRequestSortDir{
			value: "asc",
		},
		DESC: ListCollectorChannelsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListCollectorChannelsRequestSortDir) Value() string {
	return c.value
}

func (c ListCollectorChannelsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCollectorChannelsRequestSortDir) UnmarshalJSON(b []byte) error {
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
