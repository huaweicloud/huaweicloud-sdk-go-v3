package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeDbVersionFinalConfirmRequestBody **参数解释**  小版本升级成功后，确认是否完成升级请求体。  **约束限制**  不涉及。
type UpgradeDbVersionFinalConfirmRequestBody struct {

	// **参数解释**  要进行的操作。confirm表示确认完成升级，rollback表示要回滚至原始版本  **取值范围**  confirm,rollback  **约束限制**  不涉及  **默认取值**  不涉及
	Action string `json:"action"`

	// **参数解释**  回滚时是否强制回滚  **取值范围**  true,false  **约束限制**  不涉及  **默认取值**  false
	ForceRollback *bool `json:"force_rollback,omitempty"`
}

func (o UpgradeDbVersionFinalConfirmRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeDbVersionFinalConfirmRequestBody struct{}"
	}

	return strings.Join([]string{"UpgradeDbVersionFinalConfirmRequestBody", string(data)}, " ")
}
