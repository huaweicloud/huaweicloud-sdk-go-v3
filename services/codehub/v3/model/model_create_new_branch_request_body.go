package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateNewBranchRequestBody struct {

	// 分支名称
	BranchName string `json:"branch_name"`

	// 源分支名称
	Ref string `json:"ref"`
}

func (o CreateNewBranchRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNewBranchRequestBody struct{}"
	}

	return strings.Join([]string{"CreateNewBranchRequestBody", string(data)}, " ")
}
