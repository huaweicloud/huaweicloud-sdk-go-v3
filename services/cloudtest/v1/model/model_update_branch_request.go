package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBranchRequest Request Object
type UpdateBranchRequest struct {

	// 分支URI
	BranchUri string `json:"branch_uri"`

	Body *BranchVersionInfo `json:"body,omitempty"`
}

func (o UpdateBranchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBranchRequest struct{}"
	}

	return strings.Join([]string{"UpdateBranchRequest", string(data)}, " ")
}
