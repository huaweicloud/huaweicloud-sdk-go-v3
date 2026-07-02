package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowLockWaitSessionRequest Request Object
type ShowLockWaitSessionRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认值**：  en-us。
	XLanguage *ShowLockWaitSessionRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  获取方法请参见[查询实例列表](https://support.huaweicloud.com/api-taurusdb/ListGaussMySqlInstancesUnifyStatus.html)。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 点ID。  获取方法请参见[查询实例详情信息](https://support.huaweicloud.com/api-taurusdb/ShowGaussMySqlInstanceInfoUnifyStatus.html)。 **约束限制**： 不涉及。 **取值范围**： 只能由英文字母、数字组成，且长度为36个字符。 **默认值**： 不涉及。
	NodeId string `json:"node_id"`

	// **参数解释**： 会话进程ID。  获取方法请参见[查询节点用户会话线程](https://support.huaweicloud.com/api-taurusdb/ListTaurusDbNodeProcesses.html)。  **约束限制**： 不能为空。 **取值范围**： 正整数。 **默认取值**： 不涉及。
	Pid int64 `json:"pid"`
}

func (o ShowLockWaitSessionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLockWaitSessionRequest struct{}"
	}

	return strings.Join([]string{"ShowLockWaitSessionRequest", string(data)}, " ")
}

type ShowLockWaitSessionRequestXLanguage struct {
	value string
}

type ShowLockWaitSessionRequestXLanguageEnum struct {
	ZH_CN ShowLockWaitSessionRequestXLanguage
	EN_US ShowLockWaitSessionRequestXLanguage
}

func GetShowLockWaitSessionRequestXLanguageEnum() ShowLockWaitSessionRequestXLanguageEnum {
	return ShowLockWaitSessionRequestXLanguageEnum{
		ZH_CN: ShowLockWaitSessionRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowLockWaitSessionRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowLockWaitSessionRequestXLanguage) Value() string {
	return c.value
}

func (c ShowLockWaitSessionRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowLockWaitSessionRequestXLanguage) UnmarshalJSON(b []byte) error {
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
