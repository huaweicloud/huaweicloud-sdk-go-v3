package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCocIssuesRequest Request Object
type CreateCocIssuesRequest struct {
	Body *CreateExternalIssuesRequest `json:"body,omitempty"`
}

func (o CreateCocIssuesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCocIssuesRequest struct{}"
	}

	return strings.Join([]string{"CreateCocIssuesRequest", string(data)}, " ")
}
