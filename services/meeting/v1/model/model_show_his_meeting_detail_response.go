package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowHisMeetingDetailResponse struct {
	ConferenceData *ConferenceInfo `json:"conferenceData,omitempty" xml:"conferenceData"`

	Data           *PageParticipant `json:"data,omitempty" xml:"data"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowHisMeetingDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHisMeetingDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowHisMeetingDetailResponse", string(data)}, " ")
}
