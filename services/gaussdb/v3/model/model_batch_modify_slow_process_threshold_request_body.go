package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchModifySlowProcessThresholdRequestBody **参数解释**：  批量修改实例慢会话阈值请求体。  **约束限制**：  不涉及。
type BatchModifySlowProcessThresholdRequestBody struct {

	// **参数解释**：  批量修改慢会话阈值的实例ID列表。实例ID只能由英文字母、数字组成，后缀为in07，长度为36个字符。  获取方法请参见[查询实例列表](https://support.huaweicloud.com/api-taurusdb/ListGaussMySqlInstancesUnifyStatus.html)。  **约束限制**：  列表不能为空，且列表长度小于等于100。
	InstanceIds []string `json:"instance_ids"`

	// **参数解释**：  实例慢会话阈值。单位为秒。  **约束限制**：   不涉及。  **取值范围**：  1-86400。  **默认取值**：  不涉及。
	SlowProcessThreshold int32 `json:"slow_process_threshold"`
}

func (o BatchModifySlowProcessThresholdRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchModifySlowProcessThresholdRequestBody struct{}"
	}

	return strings.Join([]string{"BatchModifySlowProcessThresholdRequestBody", string(data)}, " ")
}
