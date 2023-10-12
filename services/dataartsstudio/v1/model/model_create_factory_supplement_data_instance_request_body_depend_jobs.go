package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateFactorySupplementDataInstanceRequestBodyDependJobs struct {

	// 依赖的作业名称
	JobName *string `json:"job_name,omitempty"`

	// 依赖的作业名称的空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

func (o CreateFactorySupplementDataInstanceRequestBodyDependJobs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFactorySupplementDataInstanceRequestBodyDependJobs struct{}"
	}

	return strings.Join([]string{"CreateFactorySupplementDataInstanceRequestBodyDependJobs", string(data)}, " ")
}
