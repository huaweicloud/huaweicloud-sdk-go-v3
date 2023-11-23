package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SyncJdbcDriverRequest Request Object
type SyncJdbcDriverRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *SyncJdbcDriverRequestXLanguage `json:"X-Language,omitempty"`

	Body *UpdateDriverReq `json:"body,omitempty"`
}

func (o SyncJdbcDriverRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncJdbcDriverRequest struct{}"
	}

	return strings.Join([]string{"SyncJdbcDriverRequest", string(data)}, " ")
}

type SyncJdbcDriverRequestXLanguage struct {
	value string
}

type SyncJdbcDriverRequestXLanguageEnum struct {
	EN_US SyncJdbcDriverRequestXLanguage
	ZH_CN SyncJdbcDriverRequestXLanguage
}

func GetSyncJdbcDriverRequestXLanguageEnum() SyncJdbcDriverRequestXLanguageEnum {
	return SyncJdbcDriverRequestXLanguageEnum{
		EN_US: SyncJdbcDriverRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: SyncJdbcDriverRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c SyncJdbcDriverRequestXLanguage) Value() string {
	return c.value
}

func (c SyncJdbcDriverRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SyncJdbcDriverRequestXLanguage) UnmarshalJSON(b []byte) error {
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
