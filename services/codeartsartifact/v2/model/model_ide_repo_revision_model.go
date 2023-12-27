package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IdeRepoRevisionModel struct {

	// 仓库id集合
	RepositoryIds *string `json:"repository_ids,omitempty"`

	// 类型
	Format *string `json:"format,omitempty"`
}

func (o IdeRepoRevisionModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdeRepoRevisionModel struct{}"
	}

	return strings.Join([]string{"IdeRepoRevisionModel", string(data)}, " ")
}
