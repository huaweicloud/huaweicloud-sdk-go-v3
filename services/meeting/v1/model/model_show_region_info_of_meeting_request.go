package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowRegionInfoOfMeetingRequest struct {
	// 会议ID。

	ConferenceID string `json:"conferenceID"`
}

func (o ShowRegionInfoOfMeetingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRegionInfoOfMeetingRequest struct{}"
	}

	return strings.Join([]string{"ShowRegionInfoOfMeetingRequest", string(data)}, " ")
}
