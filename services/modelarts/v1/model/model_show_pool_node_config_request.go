package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPoolNodeConfigRequest Request Object
type ShowPoolNodeConfigRequest struct {

	// **参数解释**：池名称。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PoolName string `json:"pool_name"`
}

func (o ShowPoolNodeConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPoolNodeConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowPoolNodeConfigRequest", string(data)}, " ")
}
