package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProjectStatusDetail 评估任务状态详情。
type ProjectStatusDetail struct {

	// 采集的状态。
	ObjectCollectionStatus *string `json:"object_collection_status,omitempty"`

	// 采集的进度。
	ObjectCollectionProgress *string `json:"object_collection_progress,omitempty"`

	// 评估的状态。
	PreMigrationStatus *string `json:"pre_migration_status,omitempty"`

	// 评估的进度。
	PreMigrationProgress *string `json:"pre_migration_progress,omitempty"`
}

func (o ProjectStatusDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectStatusDetail struct{}"
	}

	return strings.Join([]string{"ProjectStatusDetail", string(data)}, " ")
}
