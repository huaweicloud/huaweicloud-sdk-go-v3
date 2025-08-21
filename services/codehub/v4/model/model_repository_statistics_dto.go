package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepositoryStatisticsDto struct {

	// **参数解释：** 提交数量。 **约束限制：** 不涉及。
	CommitCount *int32 `json:"commit_count,omitempty"`

	// **参数解释：** 存储大小。 **约束限制：** 不涉及。
	StorageSize *float64 `json:"storage_size,omitempty"`

	// **参数解释：** 仓库大小。 **约束限制：** 不涉及。
	RepositorySize *float64 `json:"repository_size,omitempty"`

	// **参数解释：** LFS对象大小。 **约束限制：** 不涉及。
	LfsObjectsSize *float64 `json:"lfs_objects_size,omitempty"`

	// **参数解释：** 租户仓库大小限制。 **约束限制：** 不涉及。
	TenantRepoSizeLimit *int64 `json:"tenant_repo_size_limit,omitempty"`
}

func (o RepositoryStatisticsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryStatisticsDto struct{}"
	}

	return strings.Join([]string{"RepositoryStatisticsDto", string(data)}, " ")
}
