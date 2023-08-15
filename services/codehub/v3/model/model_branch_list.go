package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BranchList struct {

	// 指定仓库的分支列表
	Branches *[]Branch `json:"branches,omitempty"`

	// 指定仓库的分支总数
	Total *int32 `json:"total,omitempty"`
}

func (o BranchList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BranchList struct{}"
	}

	return strings.Join([]string{"BranchList", string(data)}, " ")
}
