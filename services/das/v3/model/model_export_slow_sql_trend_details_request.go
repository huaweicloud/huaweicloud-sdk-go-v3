package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExportSlowSqlTrendDetailsRequest Request Object
type ExportSlowSqlTrendDetailsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 开始时间（Unix timestamp），单位：毫秒。
	StartAt int64 `json:"start_at"`

	// 结束时间（Unix timestamp），单位：毫秒。
	EndAt int64 `json:"end_at"`

	// 数据库类型。支持MySQL和GaussDB(for MySQL)。
	DatastoreType string `json:"datastore_type"`

	// 节点ID。
	NodeId *string `json:"node_id,omitempty"`

	// 请求语言类型。
	XLanguage *ExportSlowSqlTrendDetailsRequestXLanguage `json:"X-Language,omitempty"`

	// 偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 每页记录数，默认为20，最大取值100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ExportSlowSqlTrendDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportSlowSqlTrendDetailsRequest struct{}"
	}

	return strings.Join([]string{"ExportSlowSqlTrendDetailsRequest", string(data)}, " ")
}

type ExportSlowSqlTrendDetailsRequestXLanguage struct {
	value string
}

type ExportSlowSqlTrendDetailsRequestXLanguageEnum struct {
	EN_US ExportSlowSqlTrendDetailsRequestXLanguage
	ZH_CN ExportSlowSqlTrendDetailsRequestXLanguage
}

func GetExportSlowSqlTrendDetailsRequestXLanguageEnum() ExportSlowSqlTrendDetailsRequestXLanguageEnum {
	return ExportSlowSqlTrendDetailsRequestXLanguageEnum{
		EN_US: ExportSlowSqlTrendDetailsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ExportSlowSqlTrendDetailsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ExportSlowSqlTrendDetailsRequestXLanguage) Value() string {
	return c.value
}

func (c ExportSlowSqlTrendDetailsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportSlowSqlTrendDetailsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
