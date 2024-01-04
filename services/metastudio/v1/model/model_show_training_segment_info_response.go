package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTrainingSegmentInfoResponse Response Object
type ShowTrainingSegmentInfoResponse struct {

	// 确认过的语句索引。
	ConfirmedIndex *[]int32 `json:"confirmed_index,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowTrainingSegmentInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrainingSegmentInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowTrainingSegmentInfoResponse", string(data)}, " ")
}
