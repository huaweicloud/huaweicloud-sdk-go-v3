package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProtectedBranchesResponse Response Object
type ListProtectedBranchesResponse struct {

	// 仓库保护分支列表
	Body           *[]RepositoryProtectedBranchDto `json:"body,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListProtectedBranchesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectedBranchesResponse struct{}"
	}

	return strings.Join([]string{"ListProtectedBranchesResponse", string(data)}, " ")
}
