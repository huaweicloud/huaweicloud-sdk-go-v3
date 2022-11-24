package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BranchResponse struct {

	// 分支信息
	Branches *[]BranchesItem `json:"branches,omitempty"`

	// 总数
	Total *float64 `json:"total,omitempty"`
}

func (o BranchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BranchResponse struct{}"
	}

	return strings.Join([]string{"BranchResponse", string(data)}, " ")
}
