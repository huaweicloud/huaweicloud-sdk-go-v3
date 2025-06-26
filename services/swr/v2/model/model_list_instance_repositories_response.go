package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceRepositoriesResponse Response Object
type ListInstanceRepositoriesResponse struct {

	// 仓库总数
	Total *int32 `json:"total,omitempty"`

	// 仓库列表详情
	Repositories   *[]Repository `json:"repositories,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListInstanceRepositoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceRepositoriesResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceRepositoriesResponse", string(data)}, " ")
}
