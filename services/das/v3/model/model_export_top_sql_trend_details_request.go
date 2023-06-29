package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExportTopSqlTrendDetailsRequest Request Object
type ExportTopSqlTrendDetailsRequest struct {

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
	XLanguage *ExportTopSqlTrendDetailsRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ExportTopSqlTrendDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportTopSqlTrendDetailsRequest struct{}"
	}

	return strings.Join([]string{"ExportTopSqlTrendDetailsRequest", string(data)}, " ")
}

type ExportTopSqlTrendDetailsRequestXLanguage struct {
	value string
}

type ExportTopSqlTrendDetailsRequestXLanguageEnum struct {
	EN_US ExportTopSqlTrendDetailsRequestXLanguage
	ZH_CN ExportTopSqlTrendDetailsRequestXLanguage
}

func GetExportTopSqlTrendDetailsRequestXLanguageEnum() ExportTopSqlTrendDetailsRequestXLanguageEnum {
	return ExportTopSqlTrendDetailsRequestXLanguageEnum{
		EN_US: ExportTopSqlTrendDetailsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ExportTopSqlTrendDetailsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ExportTopSqlTrendDetailsRequestXLanguage) Value() string {
	return c.value
}

func (c ExportTopSqlTrendDetailsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportTopSqlTrendDetailsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
