package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBranchConflictResponse Response Object
type ShowBranchConflictResponse struct {

	// 源仓库id
	SourceRepositoryId *int32 `json:"source_repository_id,omitempty"`

	// 目标仓库id
	TargetRepositoryId *int32 `json:"target_repository_id,omitempty"`

	// 源分支
	SourceBranch *string `json:"source_branch,omitempty"`

	// 目标分支
	TargetBranch *string `json:"target_branch,omitempty"`

	// 是否有冲突
	IsConflict     *bool `json:"is_conflict,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowBranchConflictResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBranchConflictResponse struct{}"
	}

	return strings.Join([]string{"ShowBranchConflictResponse", string(data)}, " ")
}
