package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeConfigResponse **参数解释：** 在线服务升级配置。
type UpgradeConfigResponse struct {

	// **参数解释：** 升级类型。 **取值范围：** - ROLLING：滚动升级，默认值。
	Type *string `json:"type,omitempty"`

	RollingUpdate *RollingUpdateResponse `json:"rolling_update,omitempty"`
}

func (o UpgradeConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeConfigResponse struct{}"
	}

	return strings.Join([]string{"UpgradeConfigResponse", string(data)}, " ")
}
