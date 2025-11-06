package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateProtectedTagsResponse Response Object
type BatchUpdateProtectedTagsResponse struct {

	// 仓库保护Tag列表
	Body           *[]RepositoryProtectedTagDto `json:"body,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o BatchUpdateProtectedTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateProtectedTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateProtectedTagsResponse", string(data)}, " ")
}
