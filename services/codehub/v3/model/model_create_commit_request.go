package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCommitRequest Request Object
type CreateCommitRequest struct {

	// 仓库短id
	RepoId int32 `json:"repo_id"`

	Body *CreateCommitRequestBody `json:"body,omitempty"`
}

func (o CreateCommitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCommitRequest struct{}"
	}

	return strings.Join([]string{"CreateCommitRequest", string(data)}, " ")
}
