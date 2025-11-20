package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserQuotaDetail 配额信息
type ListUserQuotaDetail struct {

	// 分身数字人训练模型数量。-1表示无限制。
	ModelingCount2dModel *int32 `json:"modeling_count_2d_model,omitempty"`

	// 分身数字人flexus版训练模型数量。-1表示无限制。
	ModelingCount2dModelFlexus *int32 `json:"modeling_count_2d_model_flexus,omitempty"`

	// 分身数字人视频制作内容时间，单位分钟。-1表示无限制。
	VideoTime2dModel *float32 `json:"video_time_2d_model,omitempty"`

	// 分身数字人视频制作flexus版内容时间，单位分钟。-1表示无限制。
	VideoTimeFlexus2dModel *float32 `json:"video_time_flexus_2d_model,omitempty"`

	// 声音克隆基础版。-1表示无限制。
	VoiceCloneBasic *int32 `json:"voice_clone_basic,omitempty"`

	// 声音克隆进阶版。-1表示无限制。
	VoiceCloneMiddle *int32 `json:"voice_clone_middle,omitempty"`

	// 声音克隆高级版。-1表示无限制。
	VoiceCloneAdvance *int32 `json:"voice_clone_advance,omitempty"`

	// 声音克隆flexus版。-1表示无限制。
	VoiceCloneFlexus *int32 `json:"voice_clone_flexus,omitempty"`
}

func (o ListUserQuotaDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserQuotaDetail struct{}"
	}

	return strings.Join([]string{"ListUserQuotaDetail", string(data)}, " ")
}
