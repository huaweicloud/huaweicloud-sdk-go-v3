package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOfflineKeyAnalysisResponse Response Object
type CreateOfflineKeyAnalysisResponse struct {

	// **参数解释**： 离线全量key分析任务ID。 **取值范围**： 不涉及。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateOfflineKeyAnalysisResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOfflineKeyAnalysisResponse struct{}"
	}

	return strings.Join([]string{"CreateOfflineKeyAnalysisResponse", string(data)}, " ")
}
