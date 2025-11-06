package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateProtectedTagsResponse Response Object
type BatchCreateProtectedTagsResponse struct {

	// 仓库保护Tag列表
	Body           *[]RepositoryProtectedTagDto `json:"body,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o BatchCreateProtectedTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateProtectedTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateProtectedTagsResponse", string(data)}, " ")
}
