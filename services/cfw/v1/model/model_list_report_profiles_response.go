package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListReportProfilesResponse Response Object
type ListReportProfilesResponse struct {
	Data           *ListReportProfilesRespData `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListReportProfilesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReportProfilesResponse struct{}"
	}

	return strings.Join([]string{"ListReportProfilesResponse", string(data)}, " ")
}
