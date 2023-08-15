package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIssueV4Request Request Object
type CreateIssueV4Request struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	Body *CreateIssueRequestV4 `json:"body,omitempty"`
}

func (o CreateIssueV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIssueV4Request struct{}"
	}

	return strings.Join([]string{"CreateIssueV4Request", string(data)}, " ")
}
