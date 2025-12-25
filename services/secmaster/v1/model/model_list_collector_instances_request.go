package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListCollectorInstancesRequest Request Object
type ListCollectorInstancesRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 采集通道ID
	ChannelId *string `json:"channel_id,omitempty"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 节点Name
	NodeName *string `json:"node_name,omitempty"`

	// 第几页
	Offset *int32 `json:"offset,omitempty"`

	// 每页数量
	Limit *int32 `json:"limit,omitempty"`

	// 排序字段
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**: 目录排序 - asc 升序排列 - desc 降序排列  **约束限制** 不涉及 **取值范围**: - asc - desc  **默认值** 不涉及
	SortDir *ListCollectorInstancesRequestSortDir `json:"sort_dir,omitempty"`
}

func (o ListCollectorInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollectorInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListCollectorInstancesRequest", string(data)}, " ")
}

type ListCollectorInstancesRequestSortDir struct {
	value string
}

type ListCollectorInstancesRequestSortDirEnum struct {
	ASC  ListCollectorInstancesRequestSortDir
	DESC ListCollectorInstancesRequestSortDir
}

func GetListCollectorInstancesRequestSortDirEnum() ListCollectorInstancesRequestSortDirEnum {
	return ListCollectorInstancesRequestSortDirEnum{
		ASC: ListCollectorInstancesRequestSortDir{
			value: "asc",
		},
		DESC: ListCollectorInstancesRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListCollectorInstancesRequestSortDir) Value() string {
	return c.value
}

func (c ListCollectorInstancesRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCollectorInstancesRequestSortDir) UnmarshalJSON(b []byte) error {
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
