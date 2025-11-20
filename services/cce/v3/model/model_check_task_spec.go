package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckTaskSpec struct {

	// **参数解释：** 插件升级目标版本 **取值范围：** 不涉及
	AddonTargetVersion *string `json:"addonTargetVersion,omitempty"`
}

func (o CheckTaskSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckTaskSpec struct{}"
	}

	return strings.Join([]string{"CheckTaskSpec", string(data)}, " ")
}
