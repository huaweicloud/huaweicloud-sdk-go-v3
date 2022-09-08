package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateSystemIssueV4Request struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	Body *CreateSystemIssueRequestV4 `json:"body,omitempty"`
}

func (o CreateSystemIssueV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSystemIssueV4Request struct{}"
	}

	return strings.Join([]string{"CreateSystemIssueV4Request", string(data)}, " ")
}
