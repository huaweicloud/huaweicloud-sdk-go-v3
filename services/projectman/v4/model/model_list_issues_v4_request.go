package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListIssuesV4Request struct {
	// devcloud的项目id

	ProjectId string `json:"project_id"`

	Body *ListIssueRequestV4 `json:"body,omitempty"`
}

func (o ListIssuesV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIssuesV4Request struct{}"
	}

	return strings.Join([]string{"ListIssuesV4Request", string(data)}, " ")
}
