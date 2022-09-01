package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepoListInfo struct {

	// 仓库列表
	Repositorys *[]RepoInfo `json:"repositorys,omitempty" xml:"repositorys"`

	// 仓库总数
	Total *int32 `json:"total,omitempty" xml:"total"`
}

func (o RepoListInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoListInfo struct{}"
	}

	return strings.Join([]string{"RepoListInfo", string(data)}, " ")
}
