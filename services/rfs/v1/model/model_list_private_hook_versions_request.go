package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPrivateHookVersionsRequest Request Object
type ListPrivateHookVersionsRequest struct {

	// 用户指定的，对于此请求的唯一Id，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	// 私有hook的名字。此名字在domain_id+region下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。
	HookName string `json:"hook_name"`

	// 私有hook（private-hook）的唯一Id。  此Id由资源编排服务在生成私有hook的时候生成，为UUID。  由于私有hook名称仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的私有hook，删除，再重新创建一个同名私有hook。  对于团队并行开发，用户可能希望确保，当前我操作的私有hook就是我认为的那个，而不是其他队友删除后创建的同名私有hook。因此，使用Id就可以做到强匹配。  资源编排服务保证每次创建的私有hook所对应的Id都不相同，更新不会影响Id。如果给予的hook_id和当前hook的Id不一致，则返回400。
	HookId *string `json:"hook_id,omitempty"`

	// 排序字段，仅支持给予create_time
	SortKey *[]ListPrivateHookVersionsRequestSortKey `json:"sort_key,omitempty"`

	// 指定升序还是降序   * `asc` - 升序   * `desc` - 降序
	SortDir *[]ListPrivateHookVersionsRequestSortDir `json:"sort_dir,omitempty"`

	// 分页标记。当一页无法返回所有结果，上一次的请求将返回next_marker以指引还有更多页数，用户可以将next_marker中的值放到此处以查询下一页的信息。此marker只能用于与上一请求指定的相同参数的请求。不指定时默认从第一页开始查询。
	Marker *string `json:"marker,omitempty"`

	// 每页返回的最多结果数量
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListPrivateHookVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateHookVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListPrivateHookVersionsRequest", string(data)}, " ")
}

type ListPrivateHookVersionsRequestSortKey struct {
	value string
}

type ListPrivateHookVersionsRequestSortKeyEnum struct {
	CREATE_TIME ListPrivateHookVersionsRequestSortKey
}

func GetListPrivateHookVersionsRequestSortKeyEnum() ListPrivateHookVersionsRequestSortKeyEnum {
	return ListPrivateHookVersionsRequestSortKeyEnum{
		CREATE_TIME: ListPrivateHookVersionsRequestSortKey{
			value: "create_time",
		},
	}
}

func (c ListPrivateHookVersionsRequestSortKey) Value() string {
	return c.value
}

func (c ListPrivateHookVersionsRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPrivateHookVersionsRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListPrivateHookVersionsRequestSortDir struct {
	value string
}

type ListPrivateHookVersionsRequestSortDirEnum struct {
	ASC  ListPrivateHookVersionsRequestSortDir
	DESC ListPrivateHookVersionsRequestSortDir
}

func GetListPrivateHookVersionsRequestSortDirEnum() ListPrivateHookVersionsRequestSortDirEnum {
	return ListPrivateHookVersionsRequestSortDirEnum{
		ASC: ListPrivateHookVersionsRequestSortDir{
			value: "asc",
		},
		DESC: ListPrivateHookVersionsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListPrivateHookVersionsRequestSortDir) Value() string {
	return c.value
}

func (c ListPrivateHookVersionsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPrivateHookVersionsRequestSortDir) UnmarshalJSON(b []byte) error {
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
