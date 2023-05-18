package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SearchIssuesRequest struct {
	Body *ListWorkTableIssueRequestV4RequestBody `json:"body,omitempty"`
}

func (o SearchIssuesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchIssuesRequest struct{}"
	}

	return strings.Join([]string{"SearchIssuesRequest", string(data)}, " ")
}
