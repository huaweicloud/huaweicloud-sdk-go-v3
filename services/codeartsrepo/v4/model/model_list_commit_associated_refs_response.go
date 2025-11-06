package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCommitAssociatedRefsResponse Response Object
type ListCommitAssociatedRefsResponse struct {

	// 分支/tag列表
	Body           *[]string `json:"body,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListCommitAssociatedRefsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCommitAssociatedRefsResponse struct{}"
	}

	return strings.Join([]string{"ListCommitAssociatedRefsResponse", string(data)}, " ")
}
