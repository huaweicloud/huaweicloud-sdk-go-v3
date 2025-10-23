package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectRepositoriesResponse Response Object
type ListProjectRepositoriesResponse struct {

	// 仓库信息
	Body           *[]BasicRepositoryDto `json:"body,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListProjectRepositoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectRepositoriesResponse struct{}"
	}

	return strings.Join([]string{"ListProjectRepositoriesResponse", string(data)}, " ")
}
