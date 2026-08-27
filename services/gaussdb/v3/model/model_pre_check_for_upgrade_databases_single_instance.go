package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PreCheckForUpgradeDatabasesSingleInstance **参数解释**：  升级预检查实例信息。  **取值范围**：  不涉及。
type PreCheckForUpgradeDatabasesSingleInstance struct {

	// **参数解释**：  实例当前的内核版本。可通过调用[查询内核版本信息](https://support.huaweicloud.com/api-taurusdb/ShowInstanceDatabaseVersion.html)接口获取。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	CurrentVersion string `json:"current_version"`

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  获取方法请参见[查询实例列表](https://support.huaweicloud.com/api-taurusdb/ListGaussMySqlInstancesUnifyStatus.html)。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`
}

func (o PreCheckForUpgradeDatabasesSingleInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreCheckForUpgradeDatabasesSingleInstance struct{}"
	}

	return strings.Join([]string{"PreCheckForUpgradeDatabasesSingleInstance", string(data)}, " ")
}
