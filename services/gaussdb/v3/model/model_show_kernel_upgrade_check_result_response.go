package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKernelUpgradeCheckResultResponse Response Object
type ShowKernelUpgradeCheckResultResponse struct {

	// **参数解释**：  预检查结果。  **取值范围**：  - true：成功。 - false：失败。
	UpgradePrecheckResult *string `json:"upgrade_precheck_result,omitempty"`

	// **参数解释**：  预检查完成时间。  **取值范围**：  格式为UNIX时间戳，单位是毫秒，时区为UTC标准时区。
	UpdatedAt *int64 `json:"updated_at,omitempty"`

	// **参数解释**：  实例预检查详情。
	UpgradePrecheckDetail *[]UpgradeDatabasePrecheckResult `json:"upgrade_precheck_detail,omitempty"`
	HttpStatusCode        int                              `json:"-"`
}

func (o ShowKernelUpgradeCheckResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKernelUpgradeCheckResultResponse struct{}"
	}

	return strings.Join([]string{"ShowKernelUpgradeCheckResultResponse", string(data)}, " ")
}
