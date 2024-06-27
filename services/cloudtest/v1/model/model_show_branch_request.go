package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBranchRequest Request Object
type ShowBranchRequest struct {

	// 分支URI
	BranchId string `json:"branch_id"`

	// 项目ID
	ProjectUuid *string `json:"project_uuid,omitempty"`
}

func (o ShowBranchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBranchRequest struct{}"
	}

	return strings.Join([]string{"ShowBranchRequest", string(data)}, " ")
}
