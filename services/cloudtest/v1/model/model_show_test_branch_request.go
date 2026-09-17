package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTestBranchRequest Request Object
type ShowTestBranchRequest struct {

	// 分支URI
	BranchUri string `json:"branch_uri"`

	// 项目ID
	ProjectUuid *string `json:"project_uuid,omitempty"`
}

func (o ShowTestBranchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTestBranchRequest struct{}"
	}

	return strings.Join([]string{"ShowTestBranchRequest", string(data)}, " ")
}
