package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBinlogSwitchResponse Response Object
type UpdateBinlogSwitchResponse struct {

	// **参数解释**:  开启或关闭Binlog日志开关操作结果。  **取值范围**：  - SUCCESS：成功。 - FAILED：失败。
	Result *string `json:"result,omitempty"`

	// **参数解释**:  任务ID。  **取值范围**：  不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateBinlogSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBinlogSwitchResponse struct{}"
	}

	return strings.Join([]string{"UpdateBinlogSwitchResponse", string(data)}, " ")
}
