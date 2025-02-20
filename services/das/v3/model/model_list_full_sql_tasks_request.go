package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListFullSqlTasksRequest Request Object
type ListFullSqlTasksRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 最小任务起止时间（Unix timestamp），单位：毫秒。
	RangeLeft *int64 `json:"range_left,omitempty"`

	// 最大任务起止时间（Unix timestamp），单位：毫秒。
	RangeRight *int64 `json:"range_right,omitempty"`

	// 最小任务创建时间（Unix timestamp），单位：毫秒。
	CreateAtLeft *int64 `json:"create_at_left,omitempty"`

	// 最大任务创建时间（Unix timestamp），单位：毫秒。
	CreateAtRight *int64 `json:"create_at_right,omitempty"`

	// 用户名。
	User *string `json:"user,omitempty"`

	// 关键字。
	Keyword *string `json:"keyword,omitempty"`

	// 数据库名。
	DbName *string `json:"db_name,omitempty"`

	// 操作类型。
	Operation *string `json:"operation,omitempty"`

	// 线程ID。
	ThreadId *string `json:"thread_id,omitempty"`

	// 事务ID。
	TrxId *string `json:"trx_id,omitempty"`

	// 执行状态（0:成功，1:失败）。
	Status *string `json:"status,omitempty"`

	// SQL模板ID。
	SqlTemplateId *string `json:"sql_template_id,omitempty"`

	// 排序字段（create_at:任务创建时间, range_start_at,range_end_at:任务起止时间）。
	SortField *string `json:"sort_field,omitempty"`

	// 排序顺序（true:正序, false:逆序）。
	Asc *bool `json:"asc,omitempty"`

	// 页码。
	Page *int32 `json:"page,omitempty"`

	// 每页记录数。最大为100。
	Limit *int32 `json:"limit,omitempty"`

	// 请求语言类型。
	XLanguage *ListFullSqlTasksRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListFullSqlTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFullSqlTasksRequest struct{}"
	}

	return strings.Join([]string{"ListFullSqlTasksRequest", string(data)}, " ")
}

type ListFullSqlTasksRequestXLanguage struct {
	value string
}

type ListFullSqlTasksRequestXLanguageEnum struct {
	EN_US ListFullSqlTasksRequestXLanguage
	ZH_CN ListFullSqlTasksRequestXLanguage
}

func GetListFullSqlTasksRequestXLanguageEnum() ListFullSqlTasksRequestXLanguageEnum {
	return ListFullSqlTasksRequestXLanguageEnum{
		EN_US: ListFullSqlTasksRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListFullSqlTasksRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListFullSqlTasksRequestXLanguage) Value() string {
	return c.value
}

func (c ListFullSqlTasksRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFullSqlTasksRequestXLanguage) UnmarshalJSON(b []byte) error {
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
