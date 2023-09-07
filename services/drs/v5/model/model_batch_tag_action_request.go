package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchTagActionRequest Request Object
type BatchTagActionRequest struct {

	// 资源类型。 - migration：实时迁移 - sync：实时同步 - cloudDataGuard：实时灾备 - subscription：数据订阅 - backupMigration：备份迁移 - replay：仿真回放
	ResourceType string `json:"resource_type"`

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *BatchTagActionRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchDealResourceTagReq `json:"body,omitempty"`
}

func (o BatchTagActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTagActionRequest struct{}"
	}

	return strings.Join([]string{"BatchTagActionRequest", string(data)}, " ")
}

type BatchTagActionRequestXLanguage struct {
	value string
}

type BatchTagActionRequestXLanguageEnum struct {
	EN_US BatchTagActionRequestXLanguage
	ZH_CN BatchTagActionRequestXLanguage
}

func GetBatchTagActionRequestXLanguageEnum() BatchTagActionRequestXLanguageEnum {
	return BatchTagActionRequestXLanguageEnum{
		EN_US: BatchTagActionRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchTagActionRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchTagActionRequestXLanguage) Value() string {
	return c.value
}

func (c BatchTagActionRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchTagActionRequestXLanguage) UnmarshalJSON(b []byte) error {
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
