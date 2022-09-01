package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepoListInfoV2 struct {

	// 仓库列表
	Repositories *[]RepoInfoV2 `json:"repositories,omitempty" xml:"repositories"`

	// 仓库总数
	Total *int32 `json:"total,omitempty" xml:"total"`
}

func (o RepoListInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoListInfoV2 struct{}"
	}

	return strings.Join([]string{"RepoListInfoV2", string(data)}, " ")
}
