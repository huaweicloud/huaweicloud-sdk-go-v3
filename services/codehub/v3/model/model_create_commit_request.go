package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateCommitRequest struct {

	// 仓库短id
	RepoId int32 `json:"repo_id" xml:"repo_id"`

	Body *CreateCommitRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateCommitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCommitRequest struct{}"
	}

	return strings.Join([]string{"CreateCommitRequest", string(data)}, " ")
}
