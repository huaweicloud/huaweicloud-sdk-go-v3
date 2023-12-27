package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFileTreeRequest Request Object
type ShowFileTreeRequest struct {

	// 租户id
	TenantId string `json:"tenant_id"`

	// 项目id
	ProjectId string `json:"project_id"`

	// 仓库名称
	RepoName string `json:"repo_name"`

	// 仓库中路径
	Path string `json:"path"`

	// 实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// 是否是回收站文件
	IsRecycleBin *bool `json:"is_recycle_bin,omitempty"`
}

func (o ShowFileTreeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFileTreeRequest struct{}"
	}

	return strings.Join([]string{"ShowFileTreeRequest", string(data)}, " ")
}
