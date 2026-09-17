package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBranchRequest Request Object
type DeleteBranchRequest struct {

	// 分支URI
	BranchUri string `json:"branch_uri"`

	// 项目id
	ProjectUuid *string `json:"project_uuid,omitempty"`

	// 是否异步执行
	IsAsync *bool `json:"is_async,omitempty"`
}

func (o DeleteBranchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBranchRequest struct{}"
	}

	return strings.Join([]string{"DeleteBranchRequest", string(data)}, " ")
}
