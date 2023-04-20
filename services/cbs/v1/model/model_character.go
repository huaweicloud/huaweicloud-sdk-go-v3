package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 形象
type Character struct {

	// 创建时间
	CreateTime string `json:"create_time"`

	// 更新时间
	UpdateTime string `json:"update_time"`

	// 形象的个人姓名
	CharacterName string `json:"character_name"`

	Gender *int32 `json:"gender,omitempty"`

	// 形象id
	Id *string `json:"id,omitempty"`

	// 形象名
	Name *string `json:"name,omitempty"`

	// 形象obs地址
	PhotoUrl *string `json:"photo_url,omitempty"`

	// 姿态： 0：站姿全身 1：站姿半身 2：坐姿全身 3：坐姿半身
	Posture *int32 `json:"posture,omitempty"`

	// 估算的训练结束时间
	TrainFinishTimeEstimate *string `json:"train_finish_time_estimate,omitempty"`

	// 训练开始时间
	TrainStartTime *string `json:"train_start_time,omitempty"`

	// 训练状态： 0：预处理 1：训练中 2：训练成功 3：训练失败 4：预览视频生成中
	TrainStatus *int32 `json:"train_status,omitempty"`

	// 形象类型： 0：预制形象 1：用户自定义形象
	Type *int32 `json:"type,omitempty"`
}

func (o Character) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Character struct{}"
	}

	return strings.Join([]string{"Character", string(data)}, " ")
}
