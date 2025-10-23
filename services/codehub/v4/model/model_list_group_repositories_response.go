package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupRepositoriesResponse Response Object
type ListGroupRepositoriesResponse struct {

	// 仓库信息
	Body           *[]BasicRepositoryDto `json:"body,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListGroupRepositoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupRepositoriesResponse struct{}"
	}

	return strings.Join([]string{"ListGroupRepositoriesResponse", string(data)}, " ")
}
