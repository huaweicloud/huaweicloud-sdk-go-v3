package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDiffCommitRequest Request Object
type ShowDiffCommitRequest struct {

	// 仓库短id
	RepoId int32 `json:"repo_id"`

	// commit id，仓库的branch名或tag名
	Sha string `json:"sha"`
}

func (o ShowDiffCommitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDiffCommitRequest struct{}"
	}

	return strings.Join([]string{"ShowDiffCommitRequest", string(data)}, " ")
}
