package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateProtectedBranchResponse Response Object
type BatchCreateProtectedBranchResponse struct {

	// 仓库保护分支列表
	Body           *[]RepositoryProtectedBranchDto `json:"body,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o BatchCreateProtectedBranchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateProtectedBranchResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateProtectedBranchResponse", string(data)}, " ")
}
