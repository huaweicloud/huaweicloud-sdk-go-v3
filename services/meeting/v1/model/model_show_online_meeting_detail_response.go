package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowOnlineMeetingDetailResponse struct {
	ConferenceData *ConferenceInfo `json:"conferenceData,omitempty"`

	Data           *PageParticipant `json:"data,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowOnlineMeetingDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOnlineMeetingDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowOnlineMeetingDetailResponse", string(data)}, " ")
}
