package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateIssuesResponse Response Object
type AssociateIssuesResponse struct {
	Body           *[]AssociateIpdIssuesResp `json:"body,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o AssociateIssuesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateIssuesResponse struct{}"
	}

	return strings.Join([]string{"AssociateIssuesResponse", string(data)}, " ")
}
