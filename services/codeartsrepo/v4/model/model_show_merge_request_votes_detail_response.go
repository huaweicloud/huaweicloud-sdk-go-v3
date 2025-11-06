package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMergeRequestVotesDetailResponse Response Object
type ShowMergeRequestVotesDetailResponse struct {

	// **参数解释：** 多人合计总分数。
	Scores *int32 `json:"scores,omitempty"`

	// **参数解释：** 合并请求id。
	MergeRequestId *int32 `json:"merge_request_id,omitempty"`

	// **参数解释：** 合并请求作者名。
	MergeRequestCreator *string `json:"merge_request_creator,omitempty"`

	// **参数解释：** 个人打分详情。
	Votes          *[]MergeRequestVotesDto `json:"votes,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowMergeRequestVotesDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMergeRequestVotesDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowMergeRequestVotesDetailResponse", string(data)}, " ")
}
