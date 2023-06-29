package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHisMeetingDetailResponse Response Object
type ShowHisMeetingDetailResponse struct {
	ConferenceData *ConferenceInfo `json:"conferenceData,omitempty"`

	Data           *PageParticipant `json:"data,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowHisMeetingDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHisMeetingDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowHisMeetingDetailResponse", string(data)}, " ")
}
