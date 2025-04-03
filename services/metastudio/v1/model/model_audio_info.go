package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AudioInfo struct {

	// **参数解释**： 音频id。 > * 获取方式：剧本为音频驱动时，查询剧本详情或者更新剧本会返回audio_id  **约束限制**： 不涉及 **默认取值**： 不涉及
	AudioId *int32 `json:"audio_id,omitempty"`
}

func (o AudioInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AudioInfo struct{}"
	}

	return strings.Join([]string{"AudioInfo", string(data)}, " ")
}
