package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMergeRequestVoteResponse Response Object
type UpdateMergeRequestVoteResponse struct {

	// **参数解释：** 打分记录的id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 合并请求id。
	MergeRequestId *int32 `json:"merge_request_id,omitempty"`

	// **参数解释：** 分数。
	Score *int32 `json:"score,omitempty"`

	Author         *UserSafeDto `json:"author,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateMergeRequestVoteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMergeRequestVoteResponse struct{}"
	}

	return strings.Join([]string{"UpdateMergeRequestVoteResponse", string(data)}, " ")
}
