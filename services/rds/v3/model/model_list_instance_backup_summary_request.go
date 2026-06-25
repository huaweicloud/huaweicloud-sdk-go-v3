package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceBackupSummaryRequest Request Object
type ListInstanceBackupSummaryRequest struct {

	// **参数解释**：  引擎类型。只支持筛选RDS for MySQL,RDS for MariaDB, TaurusDB标准版引擎  **约束限制**：  不涉及。  **取值范围**：  - mysql - mariadb - taurusdb  **默认取值**：  不涉及。
	Engine *string `json:"engine,omitempty"`

	// **参数解释**：  排序字段。 根据指定字段排序查询结果。  **约束限制**：  不涉及。  **取值范围**： - backup_used_space:根据备份数据总量排序 - obs ：根据日志备份用量排序  **默认取值**：  不涉及。
	OrderField *string `json:"order_field,omitempty"`

	// **参数解释**：  排序规则。  **约束限制**：  不涉及。  **取值范围**：  - ASC ： 根据排序字段升序 - DESC ：根据排序字段降序  **默认取值**：  DESC
	OrderRule *string `json:"order_rule,omitempty"`

	// **参数解释**：  实例ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**：  实例名称。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	InstanceName *string `json:"instance_name,omitempty"`

	// **参数解释**：  索引位置，偏移量。从第一条数据偏移offset条数据后开始查询。  **约束限制**：  必须为数字，不能为负数。  **取值范围**：  大于等于0的整数。  **默认取值**：  0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：  查询记录数。  **约束限制**：  不涉及。  **取值范围**：  默认为10，不能为负数，最小值为1，最大值为50。  **默认取值**：  10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ListInstanceBackupSummaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceBackupSummaryRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceBackupSummaryRequest", string(data)}, " ")
}
