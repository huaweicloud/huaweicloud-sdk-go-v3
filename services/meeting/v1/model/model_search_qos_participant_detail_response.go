package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SearchQosParticipantDetailResponse struct {
	User *QosParticipantInfo `json:"user,omitempty" xml:"user"`

	Qos            *QosInfo `json:"qos,omitempty" xml:"qos"`
	HttpStatusCode int      `json:"-"`
}

func (o SearchQosParticipantDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchQosParticipantDetailResponse struct{}"
	}

	return strings.Join([]string{"SearchQosParticipantDetailResponse", string(data)}, " ")
}
