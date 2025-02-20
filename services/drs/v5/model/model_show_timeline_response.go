package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTimelineResponse Response Object
type ShowTimelineResponse struct {

	// 总数。
	Count *int64 `json:"count,omitempty"`

	// 时间轴列表。
	Timelines      *[]TimelineInfo `json:"timelines,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowTimelineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTimelineResponse struct{}"
	}

	return strings.Join([]string{"ShowTimelineResponse", string(data)}, " ")
}
