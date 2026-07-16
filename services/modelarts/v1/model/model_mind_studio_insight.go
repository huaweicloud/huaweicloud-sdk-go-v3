package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MindStudioInsight MindStudio Insight连接信息。
type MindStudioInsight struct {

	// **参数解释**：训练作业的MindStudio Insight地址。 **取值范围**：不涉及。
	Url *string `json:"url,omitempty"`

	// **参数解释**：训练作业的MindStudio Insight token。 **取值范围**：不涉及。
	Token *string `json:"token,omitempty"`
}

func (o MindStudioInsight) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MindStudioInsight struct{}"
	}

	return strings.Join([]string{"MindStudioInsight", string(data)}, " ")
}
