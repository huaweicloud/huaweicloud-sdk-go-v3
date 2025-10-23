package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCherryPickMergeRequestResponse Response Object
type CreateCherryPickMergeRequestResponse struct {

	// CherryPick结果
	State *string `json:"state,omitempty"`

	// CherryPick标题
	Title *string `json:"title,omitempty"`

	// CherryPick临时分支名名称
	CherryPickBranchName *string `json:"cherry_pick_branch_name,omitempty"`
	HttpStatusCode       int     `json:"-"`
}

func (o CreateCherryPickMergeRequestResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCherryPickMergeRequestResponse struct{}"
	}

	return strings.Join([]string{"CreateCherryPickMergeRequestResponse", string(data)}, " ")
}
