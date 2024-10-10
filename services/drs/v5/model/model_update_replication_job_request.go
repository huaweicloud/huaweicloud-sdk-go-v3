package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateReplicationJobRequest Request Object
type UpdateReplicationJobRequest struct {

	// 备份迁移任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。 en-us：英文 zh-cn：中文
	XLanguage *UpdateReplicationJobRequestXLanguage `json:"X-Language,omitempty"`

	Body *ModifyOfflineTaskReq `json:"body,omitempty"`
}

func (o UpdateReplicationJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateReplicationJobRequest struct{}"
	}

	return strings.Join([]string{"UpdateReplicationJobRequest", string(data)}, " ")
}

type UpdateReplicationJobRequestXLanguage struct {
	value string
}

type UpdateReplicationJobRequestXLanguageEnum struct {
	EN_US UpdateReplicationJobRequestXLanguage
	ZH_CN UpdateReplicationJobRequestXLanguage
}

func GetUpdateReplicationJobRequestXLanguageEnum() UpdateReplicationJobRequestXLanguageEnum {
	return UpdateReplicationJobRequestXLanguageEnum{
		EN_US: UpdateReplicationJobRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: UpdateReplicationJobRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c UpdateReplicationJobRequestXLanguage) Value() string {
	return c.value
}

func (c UpdateReplicationJobRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateReplicationJobRequestXLanguage) UnmarshalJSON(b []byte) error {
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
