package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProgressDetailResponse Response Object
type ShowProgressDetailResponse struct {

	// **参数解释:** 检查任务的执行状态。 **取值范围:** - 0：检查中。 - 1：检查失败。 - 2：检查成功。 - 3：任务中止。 - null：未检查。
	TaskStatus *int32 `json:"task_status,omitempty"`

	Progress       *ProgressDetailV2 `json:"progress,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowProgressDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProgressDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowProgressDetailResponse", string(data)}, " ")
}
