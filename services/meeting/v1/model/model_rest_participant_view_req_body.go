package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestParticipantViewReqBody 会场选看请求。
type RestParticipantViewReqBody struct {

	// 选看类型。 - 2: 选看会场
	ViewType int32 `json:"viewType"`

	// 被选看的与会者标识。
	ParticipantID string `json:"participantID"`
}

func (o RestParticipantViewReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestParticipantViewReqBody struct{}"
	}

	return strings.Join([]string{"RestParticipantViewReqBody", string(data)}, " ")
}
