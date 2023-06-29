package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNewBranchRequest Request Object
type CreateNewBranchRequest struct {

	// 仓库的主键id
	RepositoryId string `json:"repository_id"`

	Body *CreateNewBranchRequestBody `json:"body,omitempty"`
}

func (o CreateNewBranchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNewBranchRequest struct{}"
	}

	return strings.Join([]string{"CreateNewBranchRequest", string(data)}, " ")
}
