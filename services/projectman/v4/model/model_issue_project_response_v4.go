package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 项目信息
type IssueProjectResponseV4 struct {
	// 项目id

	ProjectId *string `json:"project_id,omitempty"`
	// 项目名称

	ProjectName *string `json:"project_name,omitempty"`
	// 项目数字id

	ProjectNumId *int32 `json:"project_num_id,omitempty"`
}

func (o IssueProjectResponseV4) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueProjectResponseV4 struct{}"
	}

	return strings.Join([]string{"IssueProjectResponseV4", string(data)}, " ")
}
