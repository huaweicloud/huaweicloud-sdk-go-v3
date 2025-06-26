package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNamespaceRepositoriesResponse Response Object
type ListNamespaceRepositoriesResponse struct {

	// 仓库总数
	Total *int32 `json:"total,omitempty"`

	// 仓库列表详情
	Repositories   *[]Repository `json:"repositories,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListNamespaceRepositoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNamespaceRepositoriesResponse struct{}"
	}

	return strings.Join([]string{"ListNamespaceRepositoriesResponse", string(data)}, " ")
}
