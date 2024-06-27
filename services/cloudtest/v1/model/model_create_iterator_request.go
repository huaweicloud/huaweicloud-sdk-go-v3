package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIteratorRequest Request Object
type CreateIteratorRequest struct {

	// 分支URI
	BranchUri string `json:"branch_uri"`

	Body *IteratorVersionInfo `json:"body,omitempty"`
}

func (o CreateIteratorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIteratorRequest struct{}"
	}

	return strings.Join([]string{"CreateIteratorRequest", string(data)}, " ")
}
