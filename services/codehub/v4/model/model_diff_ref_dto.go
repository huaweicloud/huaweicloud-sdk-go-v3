package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DiffRefDto 合并请求变更commit，包含base_sha、start_sha、head_sha
type DiffRefDto struct {

	// 目标分支以及源分支的共同祖先commitId
	BaseSha *string `json:"base_sha,omitempty"`

	// 源分支的最新commitId
	HeadSha *string `json:"head_sha,omitempty"`

	// 目标分支的最新commitId
	StartSha *string `json:"start_sha,omitempty"`
}

func (o DiffRefDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiffRefDto struct{}"
	}

	return strings.Join([]string{"DiffRefDto", string(data)}, " ")
}
