package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTopSqlsRequest Request Object
type ListTopSqlsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 语言。
	XLanguage *ListTopSqlsRequestXLanguage `json:"X-Language,omitempty"`

	// 排序字段: avg_cpu_time:平均CPU耗时 total_cpu_time：总CPU耗时 total_duration_time：总执行时间 avg_duration_time：平均执行时间 total_rows：总行数 avg_rows：平均行数 total_logical_reads：总逻辑读 avg_logical_reads：平均逻辑读
	SortKey *string `json:"sort_key,omitempty"`

	// TOP 数量，最大支持15个。
	Limit *int32 `json:"limit,omitempty"`

	// 搜索文本内容。
	Statement *string `json:"statement,omitempty"`

	// 排序规则： desc： 降序 asc: 升序
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListTopSqlsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopSqlsRequest struct{}"
	}

	return strings.Join([]string{"ListTopSqlsRequest", string(data)}, " ")
}

type ListTopSqlsRequestXLanguage struct {
	value string
}

type ListTopSqlsRequestXLanguageEnum struct {
	ZH_CN ListTopSqlsRequestXLanguage
	EN_US ListTopSqlsRequestXLanguage
}

func GetListTopSqlsRequestXLanguageEnum() ListTopSqlsRequestXLanguageEnum {
	return ListTopSqlsRequestXLanguageEnum{
		ZH_CN: ListTopSqlsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListTopSqlsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListTopSqlsRequestXLanguage) Value() string {
	return c.value
}

func (c ListTopSqlsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTopSqlsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
