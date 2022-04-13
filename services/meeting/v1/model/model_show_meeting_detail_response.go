package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowMeetingDetailResponse struct {
	ConferenceData *ConferenceInfo `json:"conferenceData,omitempty"`

	Data           *PageParticipant `json:"data,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowMeetingDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMeetingDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowMeetingDetailResponse", string(data)}, " ")
}
