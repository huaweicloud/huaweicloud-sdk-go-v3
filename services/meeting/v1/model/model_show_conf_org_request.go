package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowConfOrgRequest struct {

	// 会议ID。
	ConferenceID string `json:"conferenceID" xml:"conferenceID"`
}

func (o ShowConfOrgRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfOrgRequest struct{}"
	}

	return strings.Join([]string{"ShowConfOrgRequest", string(data)}, " ")
}
