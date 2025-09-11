package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowScrumIssueSeveritiesResponse Response Object
type ShowScrumIssueSeveritiesResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowScrumIssueSeveritiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScrumIssueSeveritiesResponse struct{}"
	}

	return strings.Join([]string{"ShowScrumIssueSeveritiesResponse", string(data)}, " ")
}
