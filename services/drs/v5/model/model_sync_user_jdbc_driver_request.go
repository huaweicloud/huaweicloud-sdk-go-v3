package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SyncUserJdbcDriverRequest Request Object
type SyncUserJdbcDriverRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *SyncUserJdbcDriverRequestXLanguage `json:"X-Language,omitempty"`

	Body *UpdateUserDriverReq `json:"body,omitempty"`
}

func (o SyncUserJdbcDriverRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncUserJdbcDriverRequest struct{}"
	}

	return strings.Join([]string{"SyncUserJdbcDriverRequest", string(data)}, " ")
}

type SyncUserJdbcDriverRequestXLanguage struct {
	value string
}

type SyncUserJdbcDriverRequestXLanguageEnum struct {
	EN_US SyncUserJdbcDriverRequestXLanguage
	ZH_CN SyncUserJdbcDriverRequestXLanguage
}

func GetSyncUserJdbcDriverRequestXLanguageEnum() SyncUserJdbcDriverRequestXLanguageEnum {
	return SyncUserJdbcDriverRequestXLanguageEnum{
		EN_US: SyncUserJdbcDriverRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: SyncUserJdbcDriverRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c SyncUserJdbcDriverRequestXLanguage) Value() string {
	return c.value
}

func (c SyncUserJdbcDriverRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SyncUserJdbcDriverRequestXLanguage) UnmarshalJSON(b []byte) error {
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
