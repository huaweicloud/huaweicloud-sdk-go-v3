package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTestIteratorRequest Request Object
type CreateTestIteratorRequest struct {

	// 分支URI
	BranchUri string `json:"branch_uri"`

	Body *IteratorVersionInfo `json:"body,omitempty"`
}

func (o CreateTestIteratorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTestIteratorRequest struct{}"
	}

	return strings.Join([]string{"CreateTestIteratorRequest", string(data)}, " ")
}
