package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProjectConfig 项目配置信息。
type ProjectConfig struct {

	// 项目配置名称。
	ProjectConfigId *string `json:"project_config_id,omitempty"`

	// 项目名称。
	ProjectConfigName *string `json:"project_config_name,omitempty"`

	// 存储配额。
	StorageQuota *int64 `json:"storage_quota,omitempty"`

	// 是否已经关联
	IsRelevance *bool `json:"is_relevance,omitempty"`

	// sfs集群ID。
	ClusterGroupId *string `json:"cluster_group_id,omitempty"`

	// 创建时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`
}

func (o ProjectConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectConfig struct{}"
	}

	return strings.Join([]string{"ProjectConfig", string(data)}, " ")
}
