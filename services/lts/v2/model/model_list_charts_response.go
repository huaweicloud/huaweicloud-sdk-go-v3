package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ListChartsResponse struct {

	// id
	Id *string `json:"id,omitempty"`

	// sql语句
	Sql *string `json:"sql,omitempty"`

	// 图表名称
	Title *string `json:"title,omitempty"`

	// 图表类型
	Type *ListChartsResponseType `json:"type,omitempty"`

	// 日志组id
	LogGroupId *string `json:"log_group_id,omitempty"`

	// 日志组名称
	LogGroupName *string `json:"log_group_name,omitempty"`

	// 日志组id
	LogStreamId *string `json:"log_stream_id,omitempty"`

	// 日志流名称
	LogStreamName *string `json:"log_stream_name,omitempty"`

	// 图表配置详情
	Config         *ChartConfig `json:"config,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListChartsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListChartsResponse struct{}"
	}

	return strings.Join([]string{"ListChartsResponse", string(data)}, " ")
}

type ListChartsResponseType struct {
	value string
}

type ListChartsResponseTypeEnum struct {
	TABLE  ListChartsResponseType
	BAR    ListChartsResponseType
	LINE   ListChartsResponseType
	PIE    ListChartsResponseType
	NUMBER ListChartsResponseType
}

func GetListChartsResponseTypeEnum() ListChartsResponseTypeEnum {
	return ListChartsResponseTypeEnum{
		TABLE: ListChartsResponseType{
			value: "table",
		},
		BAR: ListChartsResponseType{
			value: "bar",
		},
		LINE: ListChartsResponseType{
			value: "line",
		},
		PIE: ListChartsResponseType{
			value: "pie",
		},
		NUMBER: ListChartsResponseType{
			value: "number",
		},
	}
}

func (c ListChartsResponseType) Value() string {
	return c.value
}

func (c ListChartsResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListChartsResponseType) UnmarshalJSON(b []byte) error {
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
