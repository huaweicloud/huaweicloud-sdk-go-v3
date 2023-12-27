package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListArtifactoryComponentRequest Request Object
type ListArtifactoryComponentRequest struct {

	// 租户id
	TenantId string `json:"tenant_id"`

	// 项目id
	ProjectId string `json:"project_id"`

	// 仓库名称
	RepoName string `json:"repo_name"`

	// 仓库中路径
	Path string `json:"path"`

	// 仓库格式
	Format string `json:"format"`

	// 实例id
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o ListArtifactoryComponentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListArtifactoryComponentRequest struct{}"
	}

	return strings.Join([]string{"ListArtifactoryComponentRequest", string(data)}, " ")
}
