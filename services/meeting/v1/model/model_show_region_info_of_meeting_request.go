package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRegionInfoOfMeetingRequest Request Object
type ShowRegionInfoOfMeetingRequest struct {

	// 会议ID。 > 创建会议时返回的conferenceID。不是vmrConferenceID。
	ConferenceID string `json:"conferenceID"`
}

func (o ShowRegionInfoOfMeetingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRegionInfoOfMeetingRequest struct{}"
	}

	return strings.Join([]string{"ShowRegionInfoOfMeetingRequest", string(data)}, " ")
}
