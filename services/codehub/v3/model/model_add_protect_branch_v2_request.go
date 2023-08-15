package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddProtectBranchV2Request Request Object
type AddProtectBranchV2Request struct {

	// 仓库主键id
	RepositoryId int32 `json:"repository_id"`

	// 分支名称
	BranchName string `json:"branch_name"`

	Body *AddProtectRequest `json:"body,omitempty"`
}

func (o AddProtectBranchV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddProtectBranchV2Request struct{}"
	}

	return strings.Join([]string{"AddProtectBranchV2Request", string(data)}, " ")
}
