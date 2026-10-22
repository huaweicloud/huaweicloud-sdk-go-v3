package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAiAppsRequest Request Object
type BatchDeleteAiAppsRequest struct {

	// **参数解释：** 是否保留关联的策略 **约束限制：** 不涉及 **取值范围：** true, false **默认取值：** 不涉及
	KeepPolicy bool `json:"keep_policy"`

	Body *BatchDeleteRequest `json:"body,omitempty"`
}

func (o BatchDeleteAiAppsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAiAppsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteAiAppsRequest", string(data)}, " ")
}
