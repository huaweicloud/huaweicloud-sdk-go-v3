package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKillProcessTaskPreviewRequest Request Object
type ShowKillProcessTaskPreviewRequest struct {

	// **参数解释**：              请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us：英文。 - zh-cn：中文。  **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  获取方法请参见[查询实例列表](https://support.huaweicloud.com/api-taurusdb/ListGaussMySqlInstancesUnifyStatus.html)。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：    持续时间，单位为秒。持续时间大于等于该值的会话满足Kill条件。  **约束限制**：    整数。  **取值范围**：    由数字组成，其中数字表示的范围应为1-2147483647。  **默认取值**：    不涉及。
	Time string `json:"time"`

	// **参数解释**：  用户名。  **约束限制**：  不涉及。  **取值范围**：  名称在1到32个字符之间，由字母、数字、中划线或下划线组成，不能包含其他特殊字符。  **默认取值**：  不涉及。
	User *string `json:"user,omitempty"`

	// **参数解释**：    主机。    **约束限制**：    不涉及。    **取值范围**：    不涉及。  **默认取值**：    不涉及。
	Host *string `json:"host,omitempty"`

	// **参数解释**：    数据库名。    **约束限制**：    不涉及。   **取值范围**：    1到64个字符之间，由字母、数字、中划线或下划线组成，不能包含其他特殊字符。   **默认取值**：    不涉及。
	Db *string `json:"db,omitempty"`

	// **参数解释**：    命令。    **约束限制**：    不涉及。    **取值范围**：    不涉及。  **默认取值**：    不涉及。
	Command *string `json:"command,omitempty"`

	// **参数解释**：    SQL语句。    **约束限制**：    不涉及。  **取值范围**：    不涉及。  **默认取值**：    不涉及。
	Info *string `json:"info,omitempty"`
}

func (o ShowKillProcessTaskPreviewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKillProcessTaskPreviewRequest struct{}"
	}

	return strings.Join([]string{"ShowKillProcessTaskPreviewRequest", string(data)}, " ")
}
