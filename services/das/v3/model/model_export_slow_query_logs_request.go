package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ExportSlowQueryLogsRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
	// 数据库类型。支持MySQL、GaussDB(for MySQL)、PostgreSQL。

	DatastoreType string `json:"datastore_type"`
	// 开始时间（Unix timestamp），单位：毫秒。

	StartAt int64 `json:"start_at"`
	// 结束时间（Unix timestamp），单位：毫秒。

	EndAt int64 `json:"end_at"`
	// 每页记录数。最大为2000。

	Limit int32 `json:"limit"`
	// 指定一个标识符。获取第一页时不用赋值，获取下一页时取上页查询结果的返回值。

	Marker *string `json:"marker,omitempty"`
	// 请求语言类型。

	XLanguage *ExportSlowQueryLogsRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ExportSlowQueryLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportSlowQueryLogsRequest struct{}"
	}

	return strings.Join([]string{"ExportSlowQueryLogsRequest", string(data)}, " ")
}

type ExportSlowQueryLogsRequestXLanguage struct {
	value string
}

type ExportSlowQueryLogsRequestXLanguageEnum struct {
	EN_USZH_CN ExportSlowQueryLogsRequestXLanguage
}

func GetExportSlowQueryLogsRequestXLanguageEnum() ExportSlowQueryLogsRequestXLanguageEnum {
	return ExportSlowQueryLogsRequestXLanguageEnum{
		EN_USZH_CN: ExportSlowQueryLogsRequestXLanguage{
			value: "en-us、zh-cn",
		},
	}
}

func (c ExportSlowQueryLogsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportSlowQueryLogsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
