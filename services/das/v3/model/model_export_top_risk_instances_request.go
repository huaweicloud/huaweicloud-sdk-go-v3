package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExportTopRiskInstancesRequest Request Object
type ExportTopRiskInstancesRequest struct {

	// 开始时间（Unix timestamp），单位：毫秒。
	StartAt int64 `json:"start_at"`

	// 结束时间（Unix timestamp），单位：毫秒。
	EndAt int64 `json:"end_at"`

	// 数据库类型。
	DatastoreType string `json:"datastore_type"`

	// 返回TOP风险实例数量。
	Num *ExportTopRiskInstancesRequestNum `json:"num,omitempty"`

	// 请求语言类型。
	XLanguage *ExportTopRiskInstancesRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ExportTopRiskInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportTopRiskInstancesRequest struct{}"
	}

	return strings.Join([]string{"ExportTopRiskInstancesRequest", string(data)}, " ")
}

type ExportTopRiskInstancesRequestNum struct {
	value int32
}

type ExportTopRiskInstancesRequestNumEnum struct {
	E_10 ExportTopRiskInstancesRequestNum
	E_20 ExportTopRiskInstancesRequestNum
	E_30 ExportTopRiskInstancesRequestNum
}

func GetExportTopRiskInstancesRequestNumEnum() ExportTopRiskInstancesRequestNumEnum {
	return ExportTopRiskInstancesRequestNumEnum{
		E_10: ExportTopRiskInstancesRequestNum{
			value: 10,
		}, E_20: ExportTopRiskInstancesRequestNum{
			value: 20,
		}, E_30: ExportTopRiskInstancesRequestNum{
			value: 30,
		},
	}
}

func (c ExportTopRiskInstancesRequestNum) Value() int32 {
	return c.value
}

func (c ExportTopRiskInstancesRequestNum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportTopRiskInstancesRequestNum) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type ExportTopRiskInstancesRequestXLanguage struct {
	value string
}

type ExportTopRiskInstancesRequestXLanguageEnum struct {
	EN_US ExportTopRiskInstancesRequestXLanguage
	ZH_CN ExportTopRiskInstancesRequestXLanguage
}

func GetExportTopRiskInstancesRequestXLanguageEnum() ExportTopRiskInstancesRequestXLanguageEnum {
	return ExportTopRiskInstancesRequestXLanguageEnum{
		EN_US: ExportTopRiskInstancesRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ExportTopRiskInstancesRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ExportTopRiskInstancesRequestXLanguage) Value() string {
	return c.value
}

func (c ExportTopRiskInstancesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportTopRiskInstancesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
