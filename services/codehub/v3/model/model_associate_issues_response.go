package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateIssuesResponse Response Object
type AssociateIssuesResponse struct {
	Error *Error `json:"error,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AssociateIssuesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateIssuesResponse struct{}"
	}

	return strings.Join([]string{"AssociateIssuesResponse", string(data)}, " ")
}
