package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ArtifactInfo struct {

	// 产物类型，可选值：final(最终产物)、middle(中间产物)。
	ArtifactType *string `json:"artifact_type,omitempty"`

	// 是否最优。
	IsBest *bool `json:"is_best,omitempty"`

	// 产物id。最终产物为模型ID，中间产物为断点ID。
	ArtifactId *string `json:"artifact_id,omitempty"`

	// 续训任务数量。
	ContinueTrainNums *int32 `json:"continue_train_nums,omitempty"`

	// 产物发布成功后的资产id。
	AssetId *string `json:"asset_id,omitempty"`

	// 产物发布成功后的资产名称。
	AssetName *string `json:"asset_name,omitempty"`

	// 发布状态。
	Status *string `json:"status,omitempty"`

	// 轮数。
	Epoch *int32 `json:"epoch,omitempty"`

	// 步数。
	Steps *int32 `json:"steps,omitempty"`

	// loss值
	Loss *float64 `json:"loss,omitempty"`

	// 创建时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 产物发布失败的错误信息。
	PublishErrorMsg *string `json:"publish_error_msg,omitempty"`

	// 相关任务信息
	TaskInfos *[]ContinueTrainTask `json:"task_infos,omitempty"`
}

func (o ArtifactInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ArtifactInfo struct{}"
	}

	return strings.Join([]string{"ArtifactInfo", string(data)}, " ")
}
