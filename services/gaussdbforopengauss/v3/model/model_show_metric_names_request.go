package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowMetricNamesRequest Request Object
type ShowMetricNamesRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ShowMetricNamesRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 指标分组名称。 **约束限制**: 不涉及。 **取值范围**: - CPUMEMORY：CPU/内存。 - IOSTORAGE：IO存储。 - NETWORK：网络。 - CONNECTION：连接。 - TRANSACTION：事务。 - LOCK：锁。 - SYNCSTAT：同步状态。 - PROCESSRESOURCE：进程资源。  **默认取值**: 不涉及。
	GroupName ShowMetricNamesRequestGroupName `json:"group_name"`
}

func (o ShowMetricNamesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetricNamesRequest struct{}"
	}

	return strings.Join([]string{"ShowMetricNamesRequest", string(data)}, " ")
}

type ShowMetricNamesRequestXLanguage struct {
	value string
}

type ShowMetricNamesRequestXLanguageEnum struct {
	ZH_CN ShowMetricNamesRequestXLanguage
	EN_US ShowMetricNamesRequestXLanguage
}

func GetShowMetricNamesRequestXLanguageEnum() ShowMetricNamesRequestXLanguageEnum {
	return ShowMetricNamesRequestXLanguageEnum{
		ZH_CN: ShowMetricNamesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowMetricNamesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowMetricNamesRequestXLanguage) Value() string {
	return c.value
}

func (c ShowMetricNamesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMetricNamesRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ShowMetricNamesRequestGroupName struct {
	value string
}

type ShowMetricNamesRequestGroupNameEnum struct {
	CPUMEMORY       ShowMetricNamesRequestGroupName
	IOSTORAGE       ShowMetricNamesRequestGroupName
	NETWORK         ShowMetricNamesRequestGroupName
	CONNECTION      ShowMetricNamesRequestGroupName
	TRANSACTION     ShowMetricNamesRequestGroupName
	LOCK            ShowMetricNamesRequestGroupName
	SYNCSTAT        ShowMetricNamesRequestGroupName
	PROCESSRESOURCE ShowMetricNamesRequestGroupName
}

func GetShowMetricNamesRequestGroupNameEnum() ShowMetricNamesRequestGroupNameEnum {
	return ShowMetricNamesRequestGroupNameEnum{
		CPUMEMORY: ShowMetricNamesRequestGroupName{
			value: "CPUMEMORY",
		},
		IOSTORAGE: ShowMetricNamesRequestGroupName{
			value: "IOSTORAGE",
		},
		NETWORK: ShowMetricNamesRequestGroupName{
			value: "NETWORK",
		},
		CONNECTION: ShowMetricNamesRequestGroupName{
			value: "CONNECTION",
		},
		TRANSACTION: ShowMetricNamesRequestGroupName{
			value: "TRANSACTION",
		},
		LOCK: ShowMetricNamesRequestGroupName{
			value: "LOCK",
		},
		SYNCSTAT: ShowMetricNamesRequestGroupName{
			value: "SYNCSTAT",
		},
		PROCESSRESOURCE: ShowMetricNamesRequestGroupName{
			value: "PROCESSRESOURCE",
		},
	}
}

func (c ShowMetricNamesRequestGroupName) Value() string {
	return c.value
}

func (c ShowMetricNamesRequestGroupName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMetricNamesRequestGroupName) UnmarshalJSON(b []byte) error {
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
