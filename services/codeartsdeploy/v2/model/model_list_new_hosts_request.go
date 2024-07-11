package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListNewHostsRequest Request Object
type ListNewHostsRequest struct {

	// 项目id
	GroupId string `json:"group_id"`

	// 主机名模糊查询信息
	KeyField *string `json:"key_field,omitempty"`

	// 环境id
	EnvironmentId *string `json:"environment_id,omitempty"`

	// 页码数
	PageIndex *int32 `json:"page_index,omitempty"`

	// 每页显示的条目数量，默认为10
	PageSize *int32 `json:"page_size,omitempty"`

	// 排序字段：as_proxy|host_name|owner_name，不传使用默认排序
	SortKey *string `json:"sort_key,omitempty"`

	// 排序方式：DESC、ASC，默认为DESC
	SortDir *ListNewHostsRequestSortDir `json:"sort_dir,omitempty"`

	// 是否为代理机
	AsProxy *bool `json:"as_proxy,omitempty"`
}

func (o ListNewHostsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNewHostsRequest struct{}"
	}

	return strings.Join([]string{"ListNewHostsRequest", string(data)}, " ")
}

type ListNewHostsRequestSortDir struct {
	value string
}

type ListNewHostsRequestSortDirEnum struct {
	DESC ListNewHostsRequestSortDir
	ASC  ListNewHostsRequestSortDir
}

func GetListNewHostsRequestSortDirEnum() ListNewHostsRequestSortDirEnum {
	return ListNewHostsRequestSortDirEnum{
		DESC: ListNewHostsRequestSortDir{
			value: "DESC",
		},
		ASC: ListNewHostsRequestSortDir{
			value: "ASC",
		},
	}
}

func (c ListNewHostsRequestSortDir) Value() string {
	return c.value
}

func (c ListNewHostsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNewHostsRequestSortDir) UnmarshalJSON(b []byte) error {
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
