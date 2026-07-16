package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPoolStatisticsRequest Request Object
type ShowPoolStatisticsRequest struct {

	// **参数解释**：工作空间，默认值为0。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	WorkspaceId *string `json:"workspaceId,omitempty"`
}

func (o ShowPoolStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPoolStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ShowPoolStatisticsRequest", string(data)}, " ")
}
