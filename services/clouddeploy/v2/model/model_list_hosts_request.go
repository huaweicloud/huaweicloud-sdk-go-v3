package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ListHostsRequest struct {
	// 主机组id

	GroupId string `json:"group_id"`
	// 是否为代理机

	AsProxy *bool `json:"as_proxy,omitempty"`
	// 偏移量,表示从此偏移量开始查询,offset大于等于0

	Offset *int32 `json:"offset,omitempty"`
	// 每页显示的条目数量，默认为1000

	Limit *int32 `json:"limit,omitempty"`
	// 主机名，可输入中英文，数字和符号(-_.)

	Name *string `json:"name,omitempty"`
	// 排序字段，支持：AS_PROXY|HOST_NAME|OS|OWNER_NAME|as_proxy|host_name|os|owner_name|nickName。不填默认为：as_proxy

	SortKey *string `json:"sort_key,omitempty"`
	// 排序方式,默认为：DESC

	SortDir *ListHostsRequestSortDir `json:"sort_dir,omitempty"`
	// 返回结果是否加密

	WithAuth *bool `json:"with_auth,omitempty"`
}

func (o ListHostsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostsRequest struct{}"
	}

	return strings.Join([]string{"ListHostsRequest", string(data)}, " ")
}

type ListHostsRequestSortDir struct {
	value string
}

type ListHostsRequestSortDirEnum struct {
	DESC ListHostsRequestSortDir
	ASC  ListHostsRequestSortDir
}

func GetListHostsRequestSortDirEnum() ListHostsRequestSortDirEnum {
	return ListHostsRequestSortDirEnum{
		DESC: ListHostsRequestSortDir{
			value: "DESC",
		},
		ASC: ListHostsRequestSortDir{
			value: "ASC",
		},
	}
}

func (c ListHostsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHostsRequestSortDir) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
