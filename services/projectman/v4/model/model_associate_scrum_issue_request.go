package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateScrumIssueRequest Request Object
type AssociateScrumIssueRequest struct {
	Body *AssociateIssueRequest `json:"body,omitempty"`
}

func (o AssociateScrumIssueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateScrumIssueRequest struct{}"
	}

	return strings.Join([]string{"AssociateScrumIssueRequest", string(data)}, " ")
}
