package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmTrainingSegmentResponse Response Object
type ConfirmTrainingSegmentResponse struct {

	// 是否确认成功。
	ConfirmResult *bool `json:"confirm_result,omitempty"`

	// 讲解不匹配的文本索引。
	UnmatchedIndexHit *[]int32 `json:"unmatched_index_hit,omitempty"`
	HttpStatusCode    int      `json:"-"`
}

func (o ConfirmTrainingSegmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmTrainingSegmentResponse struct{}"
	}

	return strings.Join([]string{"ConfirmTrainingSegmentResponse", string(data)}, " ")
}
