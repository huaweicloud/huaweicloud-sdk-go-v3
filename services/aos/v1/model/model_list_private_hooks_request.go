package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPrivateHooksRequest Request Object
type ListPrivateHooksRequest struct {

	// 用户指定的，对于此请求的唯一Id，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	// 排序字段，仅支持给予create_time
	SortKey *[]ListPrivateHooksRequestSortKey `json:"sort_key,omitempty"`

	// 指定升序还是降序   * `asc` - 升序   * `desc` - 降序
	SortDir *[]ListPrivateHooksRequestSortDir `json:"sort_dir,omitempty"`

	// 分页标记。当一页无法返回所有结果，上一次的请求将返回next_marker以指引还有更多页数，用户可以将next_marker中的值放到此处以查询下一页的信息。此marker只能用于与上一请求指定的相同参数的请求。不指定时默认从第一页开始查询。
	Marker *string `json:"marker,omitempty"`

	// 每页返回的最多结果数量
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListPrivateHooksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateHooksRequest struct{}"
	}

	return strings.Join([]string{"ListPrivateHooksRequest", string(data)}, " ")
}

type ListPrivateHooksRequestSortKey struct {
	value string
}

type ListPrivateHooksRequestSortKeyEnum struct {
	CREATE_TIME ListPrivateHooksRequestSortKey
}

func GetListPrivateHooksRequestSortKeyEnum() ListPrivateHooksRequestSortKeyEnum {
	return ListPrivateHooksRequestSortKeyEnum{
		CREATE_TIME: ListPrivateHooksRequestSortKey{
			value: "create_time",
		},
	}
}

func (c ListPrivateHooksRequestSortKey) Value() string {
	return c.value
}

func (c ListPrivateHooksRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPrivateHooksRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListPrivateHooksRequestSortDir struct {
	value string
}

type ListPrivateHooksRequestSortDirEnum struct {
	ASC  ListPrivateHooksRequestSortDir
	DESC ListPrivateHooksRequestSortDir
}

func GetListPrivateHooksRequestSortDirEnum() ListPrivateHooksRequestSortDirEnum {
	return ListPrivateHooksRequestSortDirEnum{
		ASC: ListPrivateHooksRequestSortDir{
			value: "asc",
		},
		DESC: ListPrivateHooksRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListPrivateHooksRequestSortDir) Value() string {
	return c.value
}

func (c ListPrivateHooksRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPrivateHooksRequestSortDir) UnmarshalJSON(b []byte) error {
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
