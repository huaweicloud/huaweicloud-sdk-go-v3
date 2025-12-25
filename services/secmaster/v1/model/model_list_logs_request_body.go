package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListLogsRequestBody struct {

	// 数据空间ID
	DataspaceId string `json:"dataspace_id"`

	// 查询开始时间点
	From int64 `json:"from"`

	// 查询返回的原始日志条数，最大值为500
	Limit int32 `json:"limit"`

	// 查询偏移值
	Offset int32 `json:"offset"`

	// 数据管道ID
	PipeId string `json:"pipe_id"`

	// 查询语句
	Query string `json:"query"`

	// 是否按时间排序返回；可选值：asc(升序)、desc(降序)，默认为 desc
	Sort ListLogsRequestBodySort `json:"sort"`

	// 查询结束时间点
	To int64 `json:"to"`
}

func (o ListLogsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogsRequestBody struct{}"
	}

	return strings.Join([]string{"ListLogsRequestBody", string(data)}, " ")
}

type ListLogsRequestBodySort struct {
	value string
}

type ListLogsRequestBodySortEnum struct {
	DESC ListLogsRequestBodySort
	ASC  ListLogsRequestBodySort
}

func GetListLogsRequestBodySortEnum() ListLogsRequestBodySortEnum {
	return ListLogsRequestBodySortEnum{
		DESC: ListLogsRequestBodySort{
			value: "desc",
		},
		ASC: ListLogsRequestBodySort{
			value: "asc",
		},
	}
}

func (c ListLogsRequestBodySort) Value() string {
	return c.value
}

func (c ListLogsRequestBodySort) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListLogsRequestBodySort) UnmarshalJSON(b []byte) error {
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
