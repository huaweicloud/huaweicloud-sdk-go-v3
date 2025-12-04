package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUpgradeInstanceVersionResponse Response Object
type ShowUpgradeInstanceVersionResponse struct {

	// **参数解释**： 当前版本。   **取值范围**： 不涉及。
	CurrentVersion *string `json:"current_version,omitempty"`

	// **参数解释**： 最新版本。   **取值范围**： 不涉及。
	LatestVersion  *string `json:"latest_version,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowUpgradeInstanceVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUpgradeInstanceVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowUpgradeInstanceVersionResponse", string(data)}, " ")
}
