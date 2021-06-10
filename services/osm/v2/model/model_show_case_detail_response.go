package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowCaseDetailResponse struct {
	IncidentDetailInfo *IncidentDetailInfoV2 `json:"incident_detail_info,omitempty"`
	HttpStatusCode     int                   `json:"-"`
}

func (o ShowCaseDetailResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowCaseDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowCaseDetailResponse", string(data)}, " ")
}
