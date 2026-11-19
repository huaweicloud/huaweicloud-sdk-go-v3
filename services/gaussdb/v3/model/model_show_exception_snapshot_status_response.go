package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExceptionSnapshotStatusResponse Response Object
type ShowExceptionSnapshotStatusResponse struct {

	// **参数解释**：  异常快照开关状态。  **取值范围**：  - true：开启。 - false：关闭。
	Enabled        *bool `json:"enabled,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowExceptionSnapshotStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExceptionSnapshotStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowExceptionSnapshotStatusResponse", string(data)}, " ")
}
