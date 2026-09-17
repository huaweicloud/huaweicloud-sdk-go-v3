package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelScrumAssociateRequest Request Object
type CancelScrumAssociateRequest struct {
	Body *CancelAssociateIssueRequest `json:"body,omitempty"`
}

func (o CancelScrumAssociateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelScrumAssociateRequest struct{}"
	}

	return strings.Join([]string{"CancelScrumAssociateRequest", string(data)}, " ")
}
