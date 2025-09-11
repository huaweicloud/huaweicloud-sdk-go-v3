package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIpdProjectIssueRequest Request Object
type CreateIpdProjectIssueRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	Body *CreateIpdProjectIssueRequestBody `json:"body,omitempty"`
}

func (o CreateIpdProjectIssueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpdProjectIssueRequest struct{}"
	}

	return strings.Join([]string{"CreateIpdProjectIssueRequest", string(data)}, " ")
}
