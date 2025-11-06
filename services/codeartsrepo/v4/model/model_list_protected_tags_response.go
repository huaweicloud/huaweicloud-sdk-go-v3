package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProtectedTagsResponse Response Object
type ListProtectedTagsResponse struct {

	// 仓库保护Tag列表
	Body           *[]RepositoryProtectedTagDto `json:"body,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListProtectedTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectedTagsResponse struct{}"
	}

	return strings.Join([]string{"ListProtectedTagsResponse", string(data)}, " ")
}
