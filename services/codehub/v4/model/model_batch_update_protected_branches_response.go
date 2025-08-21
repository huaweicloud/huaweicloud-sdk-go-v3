package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateProtectedBranchesResponse Response Object
type BatchUpdateProtectedBranchesResponse struct {

	// 仓库保护分支列表
	Body           *[]RepositoryProtectedBranchDto `json:"body,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o BatchUpdateProtectedBranchesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateProtectedBranchesResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateProtectedBranchesResponse", string(data)}, " ")
}
