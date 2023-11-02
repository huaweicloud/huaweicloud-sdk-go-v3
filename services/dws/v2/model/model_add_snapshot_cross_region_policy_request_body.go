package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddSnapshotCrossRegionPolicyRequestBody 设置快照跨区域策略请求信息
type AddSnapshotCrossRegionPolicyRequestBody struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 目的项目ID
	DestinationProjectId string `json:"destination_project_id"`

	// 目的区域
	DestinationRegion string `json:"destination_region"`

	// 状态
	Status bool `json:"status"`

	// 保留天数
	BackKeepDay int32 `json:"back_keep_day"`
}

func (o AddSnapshotCrossRegionPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSnapshotCrossRegionPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"AddSnapshotCrossRegionPolicyRequestBody", string(data)}, " ")
}
