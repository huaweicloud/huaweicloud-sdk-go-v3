package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// QueryLtsLogParams 此参数在请求实体中，采用json字符串格式
type QueryLtsLogParams struct {

	// 搜索起始时间（UTC时间，毫秒级）。
	StartTime string `json:"start_time"`

	// 搜索结束时间（UTC时间，毫秒级）。
	EndTime string `json:"end_time"`

	// 日志过滤条件集合，不同日志来源所需字段不同。
	Labels map[string]string `json:"labels,omitempty"`

	// 日志条数统计。默认为false(不统计)，true为统计日志条数。
	IsCount *bool `json:"is_count,omitempty"`

	// 支持关键词精确搜索。关键词指相邻两个分词符之间的单词,例：error
	Keywords *string `json:"keywords,omitempty"`

	// 日志单行序列号，第一次查询时不需要此参数，后续分页查询时需要使用，可从上次查询的返回信息中获取。line_num应在start_time 和 end_time 之间。
	LineNum *string `json:"line_num,omitempty"`

	// 若已开启自定义时间功能，需要使用该字段进行分页查询。
	Time *string `json:"__time__,omitempty"`

	// 顺序或者倒序查询, 默认为false(顺序查询)
	IsDesc *bool `json:"is_desc,omitempty"`

	// 首次查询为 “init”, 分页查询时为 “forwards”或者“backwards”, 默认为首次查询“init”, 与 is_desc 参数配合进行分页查询。
	SearchType *QueryLtsLogParamsSearchType `json:"search_type,omitempty"`

	// 表示每次查询的日志条数，不填时默认为50，建议您设置为100。
	Limit *int32 `json:"limit,omitempty"`

	// 日志关键词高亮显示，默认为true（高亮显示），false为取消高亮显示。
	Highlight *bool `json:"highlight,omitempty"`

	// 日志迭代查询，默认为false（不开启迭代），true为开启迭代。
	IsIterative *bool `json:"is_iterative,omitempty"`

	// 使用带管道符的sql分析语句进行查询，需要query参数is_analysis_query为true时生效。
	Query *string `json:"query,omitempty"`

	// 是否为带管道符的sql分析语句。当该参数为true时，将依照body体中的query参数内容进行查询，且body体中除start_time与end_time以外的参数失效，分页、排序、查询结果条数等功能请依照sql语法规则实现。查询结果的响应体不同于未启用时的查询方式，将以默认列存的形式返回查询结果。当前仅对内测用户开放。响应示例：{\"analysisLogs\":[{\"field1\":\"1\",\"field2\":\"2\",\"field3\":\"3\"},{\"field1\":\"1\",\"field2\":\"2\",\"field3\":\"3\"},{\"field1\":\"1\",\"field2\":\"2\",\"field3\":\"3\"}]}
	IsAnalysisQuery *bool `json:"is_analysis_query,omitempty"`
}

func (o QueryLtsLogParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryLtsLogParams struct{}"
	}

	return strings.Join([]string{"QueryLtsLogParams", string(data)}, " ")
}

type QueryLtsLogParamsSearchType struct {
	value string
}

type QueryLtsLogParamsSearchTypeEnum struct {
	FORWARDS  QueryLtsLogParamsSearchType
	BACKWARDS QueryLtsLogParamsSearchType
}

func GetQueryLtsLogParamsSearchTypeEnum() QueryLtsLogParamsSearchTypeEnum {
	return QueryLtsLogParamsSearchTypeEnum{
		FORWARDS: QueryLtsLogParamsSearchType{
			value: "forwards",
		},
		BACKWARDS: QueryLtsLogParamsSearchType{
			value: "backwards",
		},
	}
}

func (c QueryLtsLogParamsSearchType) Value() string {
	return c.value
}

func (c QueryLtsLogParamsSearchType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryLtsLogParamsSearchType) UnmarshalJSON(b []byte) error {
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
