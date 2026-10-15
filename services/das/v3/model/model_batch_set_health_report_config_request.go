package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchSetHealthReportConfigRequest Request Object
type BatchSetHealthReportConfigRequest struct {

	// 请求语言类型。en-us：英文。 zh-cn：中文。
	XLanguage *BatchSetHealthReportConfigRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchSetHealthReportConfigReq `json:"body,omitempty"`
}

func (o BatchSetHealthReportConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetHealthReportConfigRequest struct{}"
	}

	return strings.Join([]string{"BatchSetHealthReportConfigRequest", string(data)}, " ")
}

type BatchSetHealthReportConfigRequestXLanguage struct {
	value string
}

type BatchSetHealthReportConfigRequestXLanguageEnum struct {
	EN_US BatchSetHealthReportConfigRequestXLanguage
	ZH_CN BatchSetHealthReportConfigRequestXLanguage
}

func GetBatchSetHealthReportConfigRequestXLanguageEnum() BatchSetHealthReportConfigRequestXLanguageEnum {
	return BatchSetHealthReportConfigRequestXLanguageEnum{
		EN_US: BatchSetHealthReportConfigRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchSetHealthReportConfigRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchSetHealthReportConfigRequestXLanguage) Value() string {
	return c.value
}

func (c BatchSetHealthReportConfigRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchSetHealthReportConfigRequestXLanguage) UnmarshalJSON(b []byte) error {
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
