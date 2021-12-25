package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowCommitsByBranchResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *CommitList `json:"result,omitempty"`
	// 响应状态

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCommitsByBranchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCommitsByBranchResponse struct{}"
	}

	return strings.Join([]string{"ShowCommitsByBranchResponse", string(data)}, " ")
}
