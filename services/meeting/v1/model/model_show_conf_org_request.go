package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConfOrgRequest Request Object
type ShowConfOrgRequest struct {

	// 会议ID。
	ConferenceID string `json:"conferenceID"`
}

func (o ShowConfOrgRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfOrgRequest struct{}"
	}

	return strings.Join([]string{"ShowConfOrgRequest", string(data)}, " ")
}
