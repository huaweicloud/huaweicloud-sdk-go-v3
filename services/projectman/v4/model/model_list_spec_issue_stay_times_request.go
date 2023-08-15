package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSpecIssueStayTimesRequest Request Object
type ListSpecIssueStayTimesRequest struct {
	Body *ListSpecIssueStayTimesRequestBody `json:"body,omitempty"`
}

func (o ListSpecIssueStayTimesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSpecIssueStayTimesRequest struct{}"
	}

	return strings.Join([]string{"ListSpecIssueStayTimesRequest", string(data)}, " ")
}
