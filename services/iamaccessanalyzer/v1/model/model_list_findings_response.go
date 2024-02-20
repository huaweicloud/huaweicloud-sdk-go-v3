package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFindingsResponse Response Object
type ListFindingsResponse struct {
	Findings *[]Finding `json:"findings,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListFindingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFindingsResponse struct{}"
	}

	return strings.Join([]string{"ListFindingsResponse", string(data)}, " ")
}
