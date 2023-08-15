package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReadConfig struct {

	// 插图播报配置
	ImageReadConfigs *[]ImageReadConfig `json:"image_read_configs,omitempty"`

	// 播报选项： 0：纯文本播报（使用read_content 字段） 1：插图播报（使用image_read_configs字段） 3：自定义音频播报（使用audio字段） 会根据选项进行具体的字段校验
	ReadType int32 `json:"read_type"`

	// 纯文本播报内容
	ReadContent *string `json:"read_content,omitempty"`

	// 0：左 1：中 2：右 默认：1
	CharacterPosition *int32 `json:"character_position,omitempty"`

	// 段落播报间隔 单位：ms 范围：0~5000 默认：400
	ReadContentParagraphInterval *int32 `json:"read_content_paragraph_interval,omitempty"`

	// 播报框id
	ImageFrameId *string `json:"image_frame_id,omitempty"`

	// 用户的音频文件obs地址，为https格式（如：https://cbs-digital-human-cn-north-4.obs.myhuaweicloud.com:443/audio.wav），当字段不为空时，表示将使用用户自己的音频文件。  不支持PPT和图片播报，不支持字幕。音频格式文件格式为wav 音频最长支持20分钟，支持100m大小 注意：该功能的使用需要用户启动OBS授权
	AudioUrl *string `json:"audio_url,omitempty"`
}

func (o ReadConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReadConfig struct{}"
	}

	return strings.Join([]string{"ReadConfig", string(data)}, " ")
}
