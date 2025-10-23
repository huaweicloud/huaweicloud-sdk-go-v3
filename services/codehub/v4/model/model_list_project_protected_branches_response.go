package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectProtectedBranchesResponse Response Object
type ListProjectProtectedBranchesResponse struct {

	// 项目下保护分支列表
	Body           *[]ProjectProtectedBranchApiDto `json:"body,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListProjectProtectedBranchesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectProtectedBranchesResponse struct{}"
	}

	return strings.Join([]string{"ListProjectProtectedBranchesResponse", string(data)}, " ")
}
