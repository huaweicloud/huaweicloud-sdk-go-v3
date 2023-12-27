package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListArtifactoryStorageStatisticRequest Request Object
type ListArtifactoryStorageStatisticRequest struct {

	// 租户id
	TenantId string `json:"tenant_id"`

	// 项目id
	ProjectId string `json:"project_id"`

	// 仓库id
	Repo *string `json:"repo,omitempty"`

	// 起始时间
	StartTime *string `json:"start_time,omitempty"`

	// 终止时间
	EndTime *string `json:"end_time,omitempty"`

	// 实例id
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o ListArtifactoryStorageStatisticRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListArtifactoryStorageStatisticRequest struct{}"
	}

	return strings.Join([]string{"ListArtifactoryStorageStatisticRequest", string(data)}, " ")
}
