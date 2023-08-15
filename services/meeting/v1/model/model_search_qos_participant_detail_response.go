package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchQosParticipantDetailResponse Response Object
type SearchQosParticipantDetailResponse struct {
	User *QosParticipantInfo `json:"user,omitempty"`

	Qos            *QosInfo `json:"qos,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o SearchQosParticipantDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchQosParticipantDetailResponse struct{}"
	}

	return strings.Join([]string{"SearchQosParticipantDetailResponse", string(data)}, " ")
}
