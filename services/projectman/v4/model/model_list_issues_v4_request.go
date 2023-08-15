package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIssuesV4Request Request Object
type ListIssuesV4Request struct {

	// devcloud项目的32位id
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
