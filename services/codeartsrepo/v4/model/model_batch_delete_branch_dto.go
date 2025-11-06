package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteBranchDto struct {

	// 分支列表
	Branches []string `json:"branches"`
}

func (o BatchDeleteBranchDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteBranchDto struct{}"
	}

	return strings.Join([]string{"BatchDeleteBranchDto", string(data)}, " ")
}
