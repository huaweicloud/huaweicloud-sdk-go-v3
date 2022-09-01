package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateIssueV4Request struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id" xml:"project_id"`

	Body *CreateIssueRequestV4 `json:"body,omitempty" xml:"body"`
}

func (o CreateIssueV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIssueV4Request struct{}"
	}

	return strings.Join([]string{"CreateIssueV4Request", string(data)}, " ")
}
