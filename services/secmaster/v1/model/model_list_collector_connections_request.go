package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListCollectorConnectionsRequest Request Object
type ListCollectorConnectionsRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**: 连接类型 - FILTER 过滤 - INPUT 输入 - OUTPUT 输出  **约束限制** 不涉及 **取值范围**: - FILTER - INPUT - OUTPUT  **默认值** 不涉及
	ConnectionType *ListCollectorConnectionsRequestConnectionType `json:"connection_type,omitempty"`

	// 标题
	Title *string `json:"title,omitempty"`

	// 相关描述
	Description *string `json:"description,omitempty"`

	// 第几页
	Offset *int32 `json:"offset,omitempty"`

	// 每页数量
	Limit *int32 `json:"limit,omitempty"`

	// 排序字段
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**: 目录排序 - asc 升序排列 - desc 降序排列  **约束限制** 不涉及 **取值范围**: - asc - desc  **默认值** 不涉及
	SortDir *ListCollectorConnectionsRequestSortDir `json:"sort_dir,omitempty"`
}

func (o ListCollectorConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollectorConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ListCollectorConnectionsRequest", string(data)}, " ")
}

type ListCollectorConnectionsRequestConnectionType struct {
	value string
}

type ListCollectorConnectionsRequestConnectionTypeEnum struct {
	FILTER ListCollectorConnectionsRequestConnectionType
	INPUT  ListCollectorConnectionsRequestConnectionType
	OUTPUT ListCollectorConnectionsRequestConnectionType
}

func GetListCollectorConnectionsRequestConnectionTypeEnum() ListCollectorConnectionsRequestConnectionTypeEnum {
	return ListCollectorConnectionsRequestConnectionTypeEnum{
		FILTER: ListCollectorConnectionsRequestConnectionType{
			value: "FILTER",
		},
		INPUT: ListCollectorConnectionsRequestConnectionType{
			value: "INPUT",
		},
		OUTPUT: ListCollectorConnectionsRequestConnectionType{
			value: "OUTPUT",
		},
	}
}

func (c ListCollectorConnectionsRequestConnectionType) Value() string {
	return c.value
}

func (c ListCollectorConnectionsRequestConnectionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCollectorConnectionsRequestConnectionType) UnmarshalJSON(b []byte) error {
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

type ListCollectorConnectionsRequestSortDir struct {
	value string
}

type ListCollectorConnectionsRequestSortDirEnum struct {
	ASC  ListCollectorConnectionsRequestSortDir
	DESC ListCollectorConnectionsRequestSortDir
}

func GetListCollectorConnectionsRequestSortDirEnum() ListCollectorConnectionsRequestSortDirEnum {
	return ListCollectorConnectionsRequestSortDirEnum{
		ASC: ListCollectorConnectionsRequestSortDir{
			value: "asc",
		},
		DESC: ListCollectorConnectionsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListCollectorConnectionsRequestSortDir) Value() string {
	return c.value
}

func (c ListCollectorConnectionsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCollectorConnectionsRequestSortDir) UnmarshalJSON(b []byte) error {
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
