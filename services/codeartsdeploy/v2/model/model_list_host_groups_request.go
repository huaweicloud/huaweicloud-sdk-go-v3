package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListHostGroupsRequest Request Object
type ListHostGroupsRequest struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 局点信息
	RegionName string `json:"region_name"`

	// 操作系统：windows|linux
	Os *ListHostGroupsRequestOs `json:"os,omitempty"`

	// 偏移量,表示从此偏移量开始查询,offset大于等于0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量，默认为1000
	Limit *int32 `json:"limit,omitempty"`

	// 主机集群名
	Name *string `json:"name,omitempty"`

	// 排序字段：nickName|NAME|OWNER_NAME|CREATE_TIME|name|owner_name|create_time，不传使用默认排序
	SortKey *string `json:"sort_key,omitempty"`

	// 排序方式：DESC、ASC，默认为DESC
	SortDir *ListHostGroupsRequestSortDir `json:"sort_dir,omitempty"`
}

func (o ListHostGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListHostGroupsRequest", string(data)}, " ")
}

type ListHostGroupsRequestOs struct {
	value string
}

type ListHostGroupsRequestOsEnum struct {
	WINDOWS ListHostGroupsRequestOs
	LINUX   ListHostGroupsRequestOs
}

func GetListHostGroupsRequestOsEnum() ListHostGroupsRequestOsEnum {
	return ListHostGroupsRequestOsEnum{
		WINDOWS: ListHostGroupsRequestOs{
			value: "windows",
		},
		LINUX: ListHostGroupsRequestOs{
			value: "linux",
		},
	}
}

func (c ListHostGroupsRequestOs) Value() string {
	return c.value
}

func (c ListHostGroupsRequestOs) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHostGroupsRequestOs) UnmarshalJSON(b []byte) error {
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

type ListHostGroupsRequestSortDir struct {
	value string
}

type ListHostGroupsRequestSortDirEnum struct {
	DESC ListHostGroupsRequestSortDir
	ASC  ListHostGroupsRequestSortDir
}

func GetListHostGroupsRequestSortDirEnum() ListHostGroupsRequestSortDirEnum {
	return ListHostGroupsRequestSortDirEnum{
		DESC: ListHostGroupsRequestSortDir{
			value: "DESC",
		},
		ASC: ListHostGroupsRequestSortDir{
			value: "ASC",
		},
	}
}

func (c ListHostGroupsRequestSortDir) Value() string {
	return c.value
}

func (c ListHostGroupsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHostGroupsRequestSortDir) UnmarshalJSON(b []byte) error {
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
