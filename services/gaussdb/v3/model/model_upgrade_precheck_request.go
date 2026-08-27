package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradePrecheckRequest **参数解释**：  实例升级预检查的请求体。  **约束限制**：  不涉及。
type UpgradePrecheckRequest struct {

	// **参数解释**：  升级预检查实例信息。  **约束限制**：  不涉及。
	DatabasesInstanceInfos []PreCheckForUpgradeDatabasesSingleInstance `json:"databases_instance_infos"`
}

func (o UpgradePrecheckRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradePrecheckRequest struct{}"
	}

	return strings.Join([]string{"UpgradePrecheckRequest", string(data)}, " ")
}
