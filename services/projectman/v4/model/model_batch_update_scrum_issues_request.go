package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateScrumIssuesRequest Request Object
type BatchUpdateScrumIssuesRequest struct {
	Body *BatchUpdateRequest `json:"body,omitempty"`
}

func (o BatchUpdateScrumIssuesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateScrumIssuesRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateScrumIssuesRequest", string(data)}, " ")
}
