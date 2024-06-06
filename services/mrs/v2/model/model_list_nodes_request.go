package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListNodesRequest Request Object
type ListNodesRequest struct {

	// 集群ID。
	ClusterId string `json:"cluster_id"`

	// 要查询的节点组名称。
	NodeGroup *string `json:"node_group,omitempty"`

	// 返回结果中每页显示条数。缺省值：10。
	Limit *int32 `json:"limit,omitempty"`

	// 表示作业列表从该偏移量开始查询。缺省值：1。
	Offset *int32 `json:"offset,omitempty"`

	// 指定节点名称，支持模糊搜索。
	NodeName *string `json:"node_name,omitempty"`

	// 排序键，支持对节点名称排序。
	SortKey *ListNodesRequestSortKey `json:"sort_key,omitempty"`

	// 列表排序方式，desc为降序，asc为升序。
	SortDir *ListNodesRequestSortDir `json:"sort_dir,omitempty"`

	// 是否查询节点详情。该字段设为true时可能会影响接口性能。
	QueryNodeDetail *bool `json:"query_node_detail,omitempty"`

	// 是否查询ECS详情信息，会涉及对ECS接口调用。
	QueryEcsDetail *bool `json:"query_ecs_detail,omitempty"`

	// 指定内网IP。
	InternalIp *string `json:"internal_ip,omitempty"`
}

func (o ListNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNodesRequest struct{}"
	}

	return strings.Join([]string{"ListNodesRequest", string(data)}, " ")
}

type ListNodesRequestSortKey struct {
	value string
}

type ListNodesRequestSortKeyEnum struct {
	NODE_NAME ListNodesRequestSortKey
}

func GetListNodesRequestSortKeyEnum() ListNodesRequestSortKeyEnum {
	return ListNodesRequestSortKeyEnum{
		NODE_NAME: ListNodesRequestSortKey{
			value: "node_name",
		},
	}
}

func (c ListNodesRequestSortKey) Value() string {
	return c.value
}

func (c ListNodesRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNodesRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListNodesRequestSortDir struct {
	value string
}

type ListNodesRequestSortDirEnum struct {
	DESC ListNodesRequestSortDir
	ASC  ListNodesRequestSortDir
}

func GetListNodesRequestSortDirEnum() ListNodesRequestSortDirEnum {
	return ListNodesRequestSortDirEnum{
		DESC: ListNodesRequestSortDir{
			value: "desc",
		},
		ASC: ListNodesRequestSortDir{
			value: "asc",
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
