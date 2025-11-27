package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListReposResponse Response Object
type ListReposResponse struct {

	// 仓库源信息列表
	Items *[]RepoResponse `json:"items,omitempty"`

	// 所有页的结果的总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListReposResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReposResponse struct{}"
	}

	return strings.Join([]string{"ListReposResponse", string(data)}, " ")
}
