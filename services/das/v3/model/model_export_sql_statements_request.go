package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ExportSqlStatementsRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
	// 开始时间（Unix timestamp），单位：毫秒。

	StartAt int64 `json:"start_at"`
	// 结束时间（Unix timestamp），单位：毫秒。

	EndAt int64 `json:"end_at"`
	// 每页记录数。最大为2000。

	Limit int32 `json:"limit"`
	// 指定一个标识符。获取第一页时不用赋值，获取下一页时取上页查询结果的返回值。

	Marker *string `json:"marker,omitempty"`
	// 数据库类型。支持MySQL和GaussDB(for MySQL)。

	DatastoreType string `json:"datastore_type"`
	// 请求语言类型。

	XLanguage *ExportSqlStatementsRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ExportSqlStatementsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportSqlStatementsRequest struct{}"
	}

	return strings.Join([]string{"ExportSqlStatementsRequest", string(data)}, " ")
}

type ExportSqlStatementsRequestXLanguage struct {
	value string
}

type ExportSqlStatementsRequestXLanguageEnum struct {
	EN_US ExportSqlStatementsRequestXLanguage
	ZH_CN ExportSqlStatementsRequestXLanguage
}

func GetExportSqlStatementsRequestXLanguageEnum() ExportSqlStatementsRequestXLanguageEnum {
	return ExportSqlStatementsRequestXLanguageEnum{
		EN_US: ExportSqlStatementsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ExportSqlStatementsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ExportSqlStatementsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportSqlStatementsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
