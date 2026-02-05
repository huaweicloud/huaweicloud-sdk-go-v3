package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloudStorage 云存储。
type CloudStorage struct {

	// 云存储id。
	Id *string `json:"id,omitempty"`

	// 云存储名称。
	Name *string `json:"name,omitempty"`

	// 云存储id。
	ProjectConfigId *string `json:"project_config_id,omitempty"`

	// sfs集群ID。
	ClusterGroupId *string `json:"cluster_group_id,omitempty"`

	// 创建时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 个人目录声明数量。
	UserClaimCount *int32 `json:"user_claim_count,omitempty"`

	// 共享目录声明数量。
	ShareClaimCount *int32 `json:"share_claim_count,omitempty"`
}

func (o CloudStorage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudStorage struct{}"
	}

	return strings.Join([]string{"CloudStorage", string(data)}, " ")
}
