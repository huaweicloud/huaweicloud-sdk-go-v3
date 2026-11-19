package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteRetainedBackupRequest Request Object
type DeleteRetainedBackupRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us：英文。 - zh-cn：中文。  **默认取值**：  en-us。
	XLanguage *DeleteRetainedBackupRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  通过调用[查询回收站实例信息](https://support.huaweicloud.com/api-taurusdb/ListRecycleInstances.html)接口获取。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  待删除的回收站保留备份ID。  获取方法：调用[查询实例回收站中保留的备份列表](https://support.huaweicloud.com/api-taurusdb/ListRetainedBackups.html)接口获取。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为br07，长度为36个字符。  **默认取值**：  不涉及。
	BackupId string `json:"backup_id"`
}

func (o DeleteRetainedBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRetainedBackupRequest struct{}"
	}

	return strings.Join([]string{"DeleteRetainedBackupRequest", string(data)}, " ")
}

type DeleteRetainedBackupRequestXLanguage struct {
	value string
}

type DeleteRetainedBackupRequestXLanguageEnum struct {
	ZH_CN DeleteRetainedBackupRequestXLanguage
	EN_US DeleteRetainedBackupRequestXLanguage
}

func GetDeleteRetainedBackupRequestXLanguageEnum() DeleteRetainedBackupRequestXLanguageEnum {
	return DeleteRetainedBackupRequestXLanguageEnum{
		ZH_CN: DeleteRetainedBackupRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: DeleteRetainedBackupRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c DeleteRetainedBackupRequestXLanguage) Value() string {
	return c.value
}

func (c DeleteRetainedBackupRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteRetainedBackupRequestXLanguage) UnmarshalJSON(b []byte) error {
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
