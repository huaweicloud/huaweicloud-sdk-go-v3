package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPreviewFindingsResponse Response Object
type ListPreviewFindingsResponse struct {
	Findings *[]PreviewFinding `json:"findings,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPreviewFindingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPreviewFindingsResponse struct{}"
	}

	return strings.Join([]string{"ListPreviewFindingsResponse", string(data)}, " ")
}
