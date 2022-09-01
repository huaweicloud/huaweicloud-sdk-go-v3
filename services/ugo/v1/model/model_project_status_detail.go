package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 评估任务状态详情。
type ProjectStatusDetail struct {

	// 采集的状态。
	ObjectCollectionStatus *string `json:"object_collection_status,omitempty" xml:"object_collection_status"`

	// 采集的进度。
	ObjectCollectionProgress *string `json:"object_collection_progress,omitempty" xml:"object_collection_progress"`

	// 评估的状态。
	PreMigrationStatus *string `json:"pre_migration_status,omitempty" xml:"pre_migration_status"`

	// 评估的进度。
	PreMigrationProgress *string `json:"pre_migration_progress,omitempty" xml:"pre_migration_progress"`
}

func (o ProjectStatusDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectStatusDetail struct{}"
	}

	return strings.Join([]string{"ProjectStatusDetail", string(data)}, " ")
}
