package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckScheduleTaskExistRequestBody 查询任务存在性请求体
type CheckScheduleTaskExistRequestBody struct {

	// **参数解释**： 定时任务类型。  **约束限制**： 不涉及。  **取值范围**：   - PROXY_VERSION_UPGRADE：表示升级数据库代理的内核小版本。   - VERSION_UPGRADE：表示升级实例的内核小版本。   - RESIZE_FLAVOR：表示实例规格变更。   - REBOOT_NODE：表示重启节点。   - REBOOT_INSTANCE：表示重启实例。  **默认取值**：   不涉及。
	ScheduleType string `json:"schedule_type"`

	// **参数解释**： 数据库代理ID。 获取方法请参见[查询数据库代理信息列表](https://support.huaweicloud.com/api-taurusdb/ShowGaussMySqlProxyList.html)。  **约束限制**： 不涉及。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	ProxyId *string `json:"proxy_id,omitempty"`
}

func (o CheckScheduleTaskExistRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckScheduleTaskExistRequestBody struct{}"
	}

	return strings.Join([]string{"CheckScheduleTaskExistRequestBody", string(data)}, " ")
}
