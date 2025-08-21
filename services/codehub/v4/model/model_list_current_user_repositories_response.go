package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCurrentUserRepositoriesResponse Response Object
type ListCurrentUserRepositoriesResponse struct {

	// 仓库列表
	Body           *[]RepositorySimplestDto `json:"body,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListCurrentUserRepositoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCurrentUserRepositoriesResponse struct{}"
	}

	return strings.Join([]string{"ListCurrentUserRepositoriesResponse", string(data)}, " ")
}
