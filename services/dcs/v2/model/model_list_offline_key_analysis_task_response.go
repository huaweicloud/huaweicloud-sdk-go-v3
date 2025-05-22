package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOfflineKeyAnalysisTaskResponse Response Object
type ListOfflineKeyAnalysisTaskResponse struct {

	// **参数解释**： 实例ID。 **取值范围**： 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**： 离线全量key分析的总数。 **取值范围**： 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 离线全量key分析记录列表。
	Records        *[]OfflineKeyAnalysisRecord `json:"records,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListOfflineKeyAnalysisTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOfflineKeyAnalysisTaskResponse struct{}"
	}

	return strings.Join([]string{"ListOfflineKeyAnalysisTaskResponse", string(data)}, " ")
}
