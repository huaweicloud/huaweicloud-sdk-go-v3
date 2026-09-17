package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBranchRequest Request Object
type CreateBranchRequest struct {
	Body *BranchVersionInfo `json:"body,omitempty"`
}

func (o CreateBranchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBranchRequest struct{}"
	}

	return strings.Join([]string{"CreateBranchRequest", string(data)}, " ")
}
