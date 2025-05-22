package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteOfflineKeyAnalysisTaskResponse Response Object
type DeleteOfflineKeyAnalysisTaskResponse struct {

	// **参数解释**： 离线全量key分析任务ID。 **取值范围**： 不涉及。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteOfflineKeyAnalysisTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOfflineKeyAnalysisTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteOfflineKeyAnalysisTaskResponse", string(data)}, " ")
}
