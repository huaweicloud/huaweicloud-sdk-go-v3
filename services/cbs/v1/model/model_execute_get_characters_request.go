package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteGetCharactersRequest Request Object
type ExecuteGetCharactersRequest struct {
	Offset *string `json:"offset,omitempty"`

	Limit *string `json:"limit,omitempty"`

	// 形象类型： 0：预制形象 1：用户自定义形象
	Type *int32 `json:"type,omitempty"`

	// 训练状态： 0：预处理 1：训练中 2：训练成功 3：训练失败 4：预览视频生成中
	TrainStatus *int32 `json:"train_status,omitempty"`

	CharacterName *string `json:"character_name,omitempty"`

	// 是否只获取支持交互式的数字人；默认：false 获取全部
	SupportInteract *bool `json:"support_interact,omitempty"`

	// 性别
	Gender *string `json:"gender,omitempty"`
}

func (o ExecuteGetCharactersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteGetCharactersRequest struct{}"
	}

	return strings.Join([]string{"ExecuteGetCharactersRequest", string(data)}, " ")
}
