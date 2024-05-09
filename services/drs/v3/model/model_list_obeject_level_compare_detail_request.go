package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListObejectLevelCompareDetailRequest Request Object
type ListObejectLevelCompareDetailRequest struct {

	// 请求语言类型。
	XLanguage *ListObejectLevelCompareDetailRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`

	// 对象类型： - DB：数据库。 - TABLE：表。 - VIEW：视图。 - EVENT：事件。 - ROUTINE：存储过程和函数。 - INDEX：索引。 - TRIGGER：触发器。 - SYNONYM：同义词。 - FUNCTION：函数。 - PROCEDURE：存储过程。 - TYPE：自定义类型。 - RULE：规则。 - DEFAULT_TYPE：缺省值。 - PLAN_GUIDE：执行计划。 - CONSTRAINT：约束。 - FILE_GROUP：文件组。 - PARTITION_FUNCTION：分区函数。 - PARTITION_SCHEME：分区方案。 - TABLE_COLLATION：表的排序规则。
	CompareType string `json:"compare_type"`

	// 对比任务ID，不填写时默认返回最新的对比任务信息。
	CompareJobId *string `json:"compare_job_id,omitempty"`

	// 每页显示的条目数量，最大值1000。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListObejectLevelCompareDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListObejectLevelCompareDetailRequest struct{}"
	}

	return strings.Join([]string{"ListObejectLevelCompareDetailRequest", string(data)}, " ")
}

type ListObejectLevelCompareDetailRequestXLanguage struct {
	value string
}

type ListObejectLevelCompareDetailRequestXLanguageEnum struct {
	EN_US ListObejectLevelCompareDetailRequestXLanguage
	ZH_CN ListObejectLevelCompareDetailRequestXLanguage
}

func GetListObejectLevelCompareDetailRequestXLanguageEnum() ListObejectLevelCompareDetailRequestXLanguageEnum {
	return ListObejectLevelCompareDetailRequestXLanguageEnum{
		EN_US: ListObejectLevelCompareDetailRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListObejectLevelCompareDetailRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListObejectLevelCompareDetailRequestXLanguage) Value() string {
	return c.value
}

func (c ListObejectLevelCompareDetailRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListObejectLevelCompareDetailRequestXLanguage) UnmarshalJSON(b []byte) error {
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
