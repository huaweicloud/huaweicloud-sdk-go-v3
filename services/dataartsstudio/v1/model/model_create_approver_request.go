package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateApproverRequest Request Object
type CreateApproverRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *ApproverParam `json:"body,omitempty"`
}

func (o CreateApproverRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApproverRequest struct{}"
	}

	return strings.Join([]string{"CreateApproverRequest", string(data)}, " ")
}
