package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListHostClustersRequest Request Object
type ListHostClustersRequest struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 主机集群模糊查询信息
	Name *string `json:"name,omitempty"`

	// 操作系统：windows|linux
	Os *ListHostClustersRequestOs `json:"os,omitempty"`

	// 页码数
	PageIndex *int32 `json:"page_index,omitempty"`

	// 每页显示的条目数量，默认为10
	PageSize *int32 `json:"page_size,omitempty"`

	// 排序字段：nick_name|name|owner_name|create_time，不传使用默认排序
	SortField *string `json:"sort_field,omitempty"`

	// 排序方式：DESC、ASC，默认为DESC
	SortType *ListHostClustersRequestSortType `json:"sort_type,omitempty"`

	// 是否为代理机
	IsProxyMode *int32 `json:"is_proxy_mode,omitempty"`

	// 自定义资源池id
	SlaveClusterId *string `json:"slave_cluster_id,omitempty"`
}

func (o ListHostClustersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostClustersRequest struct{}"
	}

	return strings.Join([]string{"ListHostClustersRequest", string(data)}, " ")
}

type ListHostClustersRequestOs struct {
	value string
}

type ListHostClustersRequestOsEnum struct {
	WINDOWS ListHostClustersRequestOs
	LINUX   ListHostClustersRequestOs
}

func GetListHostClustersRequestOsEnum() ListHostClustersRequestOsEnum {
	return ListHostClustersRequestOsEnum{
		WINDOWS: ListHostClustersRequestOs{
			value: "windows",
		},
		LINUX: ListHostClustersRequestOs{
			value: "linux",
		},
	}
}

func (c ListHostClustersRequestOs) Value() string {
	return c.value
}

func (c ListHostClustersRequestOs) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHostClustersRequestOs) UnmarshalJSON(b []byte) error {
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

type ListHostClustersRequestSortType struct {
	value string
}

type ListHostClustersRequestSortTypeEnum struct {
	DESC ListHostClustersRequestSortType
	ASC  ListHostClustersRequestSortType
}

func GetListHostClustersRequestSortTypeEnum() ListHostClustersRequestSortTypeEnum {
	return ListHostClustersRequestSortTypeEnum{
		DESC: ListHostClustersRequestSortType{
			value: "DESC",
		},
		ASC: ListHostClustersRequestSortType{
			value: "ASC",
		},
	}
}

func (c ListHostClustersRequestSortType) Value() string {
	return c.value
}

func (c ListHostClustersRequestSortType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHostClustersRequestSortType) UnmarshalJSON(b []byte) error {
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
