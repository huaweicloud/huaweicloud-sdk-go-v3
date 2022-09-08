package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListBranchesResponse struct {

	// 项目分支列表。
	Branches       *[]string `json:"branches,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListBranchesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBranchesResponse struct{}"
	}

	return strings.Join([]string{"ListBranchesResponse", string(data)}, " ")
}
