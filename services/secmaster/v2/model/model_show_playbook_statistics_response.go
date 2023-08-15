package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPlaybookStatisticsResponse Response Object
type ShowPlaybookStatisticsResponse struct {

	// Error code
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	Data *PlaybookStatisticDetail `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPlaybookStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlaybookStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowPlaybookStatisticsResponse", string(data)}, " ")
}
