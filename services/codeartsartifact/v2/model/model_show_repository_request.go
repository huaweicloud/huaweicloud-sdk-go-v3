package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRepositoryRequest Request Object
type ShowRepositoryRequest struct {

	// 租户id
	TenantId string `json:"tenant_id"`

	// 项目id
	ProjectId string `json:"project_id"`

	// 仓库id
	RepoId string `json:"repo_id"`

	// 服务区
	Region *string `json:"region,omitempty"`
}

func (o ShowRepositoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryRequest struct{}"
	}

	return strings.Join([]string{"ShowRepositoryRequest", string(data)}, " ")
}
