package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProjectBranchRequest Request Object
type CreateProjectBranchRequest struct {
	Body *BranchVersionInfo `json:"body,omitempty"`
}

func (o CreateProjectBranchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectBranchRequest struct{}"
	}

	return strings.Join([]string{"CreateProjectBranchRequest", string(data)}, " ")
}
