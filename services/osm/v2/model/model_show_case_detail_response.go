package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowCaseDetailResponse struct {
	IncidentDetailInfo *IncidentDetailInfoV2 `json:"incident_detail_info,omitempty"`
	HttpStatusCode     int                   `json:"-"`
}

func (o ShowCaseDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCaseDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowCaseDetailResponse", string(data)}, " ")
}
