package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSlowlogRequest Request Object
type ListSlowlogRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 返回结果按该关键字排序，支持start_time，duration，默认为“start_time”
	SortKey *ListSlowlogRequestSortKey `json:"sort_key,omitempty"`

	// 降序或升序（分别对应desc和asc，默认为“desc”）
	SortDir *ListSlowlogRequestSortDir `json:"sort_dir,omitempty"`

	// 查询开始时间，时间为UTC时间的Unix时间戳。如：1598803200000。
	StartTime string `json:"start_time"`

	// 查询结束时间，时间为UTC时间的Unix时间戳。如：1599494399000。
	EndTime string `json:"end_time"`

	// 查询节点，分为proxy和server。
	Role *string `json:"role,omitempty"`
}

func (o ListSlowlogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowlogRequest struct{}"
	}

	return strings.Join([]string{"ListSlowlogRequest", string(data)}, " ")
}

type ListSlowlogRequestSortKey struct {
	value string
}

type ListSlowlogRequestSortKeyEnum struct {
	START_TIME ListSlowlogRequestSortKey
	DURATION   ListSlowlogRequestSortKey
}

func GetListSlowlogRequestSortKeyEnum() ListSlowlogRequestSortKeyEnum {
	return ListSlowlogRequestSortKeyEnum{
		START_TIME: ListSlowlogRequestSortKey{
			value: "start_time",
		},
		DURATION: ListSlowlogRequestSortKey{
			value: "duration",
		},
	}
}

func (c ListSlowlogRequestSortKey) Value() string {
	return c.value
}

func (c ListSlowlogRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSlowlogRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListSlowlogRequestSortDir struct {
	value string
}

type ListSlowlogRequestSortDirEnum struct {
	DESC ListSlowlogRequestSortDir
	ASC  ListSlowlogRequestSortDir
}

func GetListSlowlogRequestSortDirEnum() ListSlowlogRequestSortDirEnum {
	return ListSlowlogRequestSortDirEnum{
		DESC: ListSlowlogRequestSortDir{
			value: "desc",
		},
		ASC: ListSlowlogRequestSortDir{
			value: "asc",
		},
	}
}

func (c ListSlowlogRequestSortDir) Value() string {
	return c.value
}

func (c ListSlowlogRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSlowlogRequestSortDir) UnmarshalJSON(b []byte) error {
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
