package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListNodesRequest Request Object
type ListNodesRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 节点名称
	NodeName *string `json:"node_name,omitempty"`

	// **参数解释：** 偏移量 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Offset *int64 `json:"offset,omitempty"`

	// **参数解释：** 数据量 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Limit *int64 `json:"limit,omitempty"`

	// 限制
	SortKey *int64 `json:"sort_key,omitempty"`

	// **参数解释**: 目录排序 - asc 升序排列 - desc 降序排列  **约束限制** 不涉及 **取值范围**: - asc - desc  **默认值** 不涉及
	SortDir *ListNodesRequestSortDir `json:"sort_dir,omitempty"`
}

func (o ListNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNodesRequest struct{}"
	}

	return strings.Join([]string{"ListNodesRequest", string(data)}, " ")
}

type ListNodesRequestSortDir struct {
	value string
}

type ListNodesRequestSortDirEnum struct {
	ASC  ListNodesRequestSortDir
	DESC ListNodesRequestSortDir
}

func GetListNodesRequestSortDirEnum() ListNodesRequestSortDirEnum {
	return ListNodesRequestSortDirEnum{
		ASC: ListNodesRequestSortDir{
			value: "asc",
		},
		DESC: ListNodesRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListNodesRequestSortDir) Value() string {
	return c.value
}

func (c ListNodesRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNodesRequestSortDir) UnmarshalJSON(b []byte) error {
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
