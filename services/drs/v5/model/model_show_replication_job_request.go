package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowReplicationJobRequest Request Object
type ShowReplicationJobRequest struct {

	// 备份迁移任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。 en-us：英文 zh-cn：中文
	XLanguage *ShowReplicationJobRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowReplicationJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReplicationJobRequest struct{}"
	}

	return strings.Join([]string{"ShowReplicationJobRequest", string(data)}, " ")
}

type ShowReplicationJobRequestXLanguage struct {
	value string
}

type ShowReplicationJobRequestXLanguageEnum struct {
	EN_US ShowReplicationJobRequestXLanguage
	ZH_CN ShowReplicationJobRequestXLanguage
}

func GetShowReplicationJobRequestXLanguageEnum() ShowReplicationJobRequestXLanguageEnum {
	return ShowReplicationJobRequestXLanguageEnum{
		EN_US: ShowReplicationJobRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowReplicationJobRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowReplicationJobRequestXLanguage) Value() string {
	return c.value
}

func (c ShowReplicationJobRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowReplicationJobRequestXLanguage) UnmarshalJSON(b []byte) error {
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
