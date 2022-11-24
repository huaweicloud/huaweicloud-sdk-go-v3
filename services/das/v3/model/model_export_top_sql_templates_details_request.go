package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ExportTopSqlTemplatesDetailsRequest struct {

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

	// 排序字段（executeNum:执行次数, totalCost:总耗时, avgCost:平均耗时, totalScan: 总扫描行数, avgScan:平均扫描行数）。
	Sort *ExportTopSqlTemplatesDetailsRequestSort `json:"sort,omitempty"`

	// 排序顺序（true:正序, false:逆序）。
	Asc *bool `json:"asc,omitempty"`

	// 偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 每页记录数，默认为20，最大取值100。
	Limit *int32 `json:"limit,omitempty"`

	// 请求语言类型。
	XLanguage *ExportTopSqlTemplatesDetailsRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ExportTopSqlTemplatesDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportTopSqlTemplatesDetailsRequest struct{}"
	}

	return strings.Join([]string{"ExportTopSqlTemplatesDetailsRequest", string(data)}, " ")
}

type ExportTopSqlTemplatesDetailsRequestSort struct {
	value string
}

type ExportTopSqlTemplatesDetailsRequestSortEnum struct {
	EXECUTE_NUM ExportTopSqlTemplatesDetailsRequestSort
	TOTAL_COST  ExportTopSqlTemplatesDetailsRequestSort
	AVG_COST    ExportTopSqlTemplatesDetailsRequestSort
	TOTAL_SCAN  ExportTopSqlTemplatesDetailsRequestSort
	AVG_SCAN    ExportTopSqlTemplatesDetailsRequestSort
}

func GetExportTopSqlTemplatesDetailsRequestSortEnum() ExportTopSqlTemplatesDetailsRequestSortEnum {
	return ExportTopSqlTemplatesDetailsRequestSortEnum{
		EXECUTE_NUM: ExportTopSqlTemplatesDetailsRequestSort{
			value: "executeNum",
		},
		TOTAL_COST: ExportTopSqlTemplatesDetailsRequestSort{
			value: "totalCost",
		},
		AVG_COST: ExportTopSqlTemplatesDetailsRequestSort{
			value: "avgCost",
		},
		TOTAL_SCAN: ExportTopSqlTemplatesDetailsRequestSort{
			value: "totalScan",
		},
		AVG_SCAN: ExportTopSqlTemplatesDetailsRequestSort{
			value: "avgScan",
		},
	}
}

func (c ExportTopSqlTemplatesDetailsRequestSort) Value() string {
	return c.value
}

func (c ExportTopSqlTemplatesDetailsRequestSort) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportTopSqlTemplatesDetailsRequestSort) UnmarshalJSON(b []byte) error {
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

type ExportTopSqlTemplatesDetailsRequestXLanguage struct {
	value string
}

type ExportTopSqlTemplatesDetailsRequestXLanguageEnum struct {
	EN_US ExportTopSqlTemplatesDetailsRequestXLanguage
	ZH_CN ExportTopSqlTemplatesDetailsRequestXLanguage
}

func GetExportTopSqlTemplatesDetailsRequestXLanguageEnum() ExportTopSqlTemplatesDetailsRequestXLanguageEnum {
	return ExportTopSqlTemplatesDetailsRequestXLanguageEnum{
		EN_US: ExportTopSqlTemplatesDetailsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ExportTopSqlTemplatesDetailsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ExportTopSqlTemplatesDetailsRequestXLanguage) Value() string {
	return c.value
}

func (c ExportTopSqlTemplatesDetailsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportTopSqlTemplatesDetailsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
