package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowBackupVaultLockRequest Request Object
type ShowBackupVaultLockRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认取值**：  en-us。
	XLanguage *ShowBackupVaultLockRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  获取方法请参见[查询实例列表](https://support.huaweicloud.com/api-taurusdb/ListGaussMySqlInstancesUnifyStatus.html)。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`
}

func (o ShowBackupVaultLockRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupVaultLockRequest struct{}"
	}

	return strings.Join([]string{"ShowBackupVaultLockRequest", string(data)}, " ")
}

type ShowBackupVaultLockRequestXLanguage struct {
	value string
}

type ShowBackupVaultLockRequestXLanguageEnum struct {
	ZH_CN ShowBackupVaultLockRequestXLanguage
	EN_US ShowBackupVaultLockRequestXLanguage
}

func GetShowBackupVaultLockRequestXLanguageEnum() ShowBackupVaultLockRequestXLanguageEnum {
	return ShowBackupVaultLockRequestXLanguageEnum{
		ZH_CN: ShowBackupVaultLockRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowBackupVaultLockRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowBackupVaultLockRequestXLanguage) Value() string {
	return c.value
}

func (c ShowBackupVaultLockRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowBackupVaultLockRequestXLanguage) UnmarshalJSON(b []byte) error {
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
