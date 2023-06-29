package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAnonymousAuthRandomRequest Request Object
type CreateAnonymousAuthRandomRequest struct {

	// 会议ID。
	ConferenceID string `json:"conferenceID"`

	// 会议密码。
	XPassword string `json:"X-Password"`
}

func (o CreateAnonymousAuthRandomRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAnonymousAuthRandomRequest struct{}"
	}

	return strings.Join([]string{"CreateAnonymousAuthRandomRequest", string(data)}, " ")
}
