package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExportSlowSqlTemplatesDetailsRequest Request Object
type ExportSlowSqlTemplatesDetailsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 开始时间（Unix timestamp），单位：毫秒。
	StartAt int64 `json:"start_at"`

	// 结束时间（Unix timestamp），单位：毫秒。
	EndAt int64 `json:"end_at"`

	// 数据库类型。支持MySQL和GaussDB(for MySQL)。
	DatastoreType string `json:"datastore_type"`

	// 数据库名称。
	DbName *string `json:"db_name,omitempty"`

	// 请求语言类型。
	XLanguage *ExportSlowSqlTemplatesDetailsRequestXLanguage `json:"X-Language,omitempty"`

	// 偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 每页记录数，默认为20，最大取值100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ExportSlowSqlTemplatesDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportSlowSqlTemplatesDetailsRequest struct{}"
	}

	return strings.Join([]string{"ExportSlowSqlTemplatesDetailsRequest", string(data)}, " ")
}

type ExportSlowSqlTemplatesDetailsRequestXLanguage struct {
	value string
}

type ExportSlowSqlTemplatesDetailsRequestXLanguageEnum struct {
	EN_US ExportSlowSqlTemplatesDetailsRequestXLanguage
	ZH_CN ExportSlowSqlTemplatesDetailsRequestXLanguage
}

func GetExportSlowSqlTemplatesDetailsRequestXLanguageEnum() ExportSlowSqlTemplatesDetailsRequestXLanguageEnum {
	return ExportSlowSqlTemplatesDetailsRequestXLanguageEnum{
		EN_US: ExportSlowSqlTemplatesDetailsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ExportSlowSqlTemplatesDetailsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ExportSlowSqlTemplatesDetailsRequestXLanguage) Value() string {
	return c.value
}

func (c ExportSlowSqlTemplatesDetailsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportSlowSqlTemplatesDetailsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
