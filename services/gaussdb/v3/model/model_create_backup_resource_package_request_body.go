package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBackupResourcePackageRequestBody **参数解释**:  创建备份资源包请求体。
type CreateBackupResourcePackageRequestBody struct {

	// **参数解释**：  备份资源包规格码。  **约束限制**：  不涉及。  **取值范围**：  备份资源包规格码可根据[查询备份资源包规格](https://support.huaweicloud.com/api-taurusdb/ShowBackupResourcePackageFlavors.html)接口获取。  **默认取值**：  不涉及。
	SpecCode string `json:"spec_code"`

	// **参数解释**：  备份资源包数量。  **约束限制**：  不涉及。  **取值范围**：  1-10。  **默认取值**：  不涉及。
	Num int32 `json:"num"`

	ChargeInfo *TaurusDbChargeInfo `json:"charge_info"`
}

func (o CreateBackupResourcePackageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBackupResourcePackageRequestBody struct{}"
	}

	return strings.Join([]string{"CreateBackupResourcePackageRequestBody", string(data)}, " ")
}
