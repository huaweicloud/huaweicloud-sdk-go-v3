package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TtsJobDetail 文本转语音任务合成记录。
type TtsJobDetail struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// websocket任务ID。
	WebsocketJobId *string `json:"websocket_job_id,omitempty"`

	// 音色ID
	AssetId *string `json:"asset_id,omitempty"`

	// 声音版本
	TtsServiceEnum *string `json:"tts_service_enum,omitempty"`

	// 文本长度
	TextLength *int32 `json:"text_length,omitempty"`

	// 任务创建时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 输出音频文件格式。默认WAV。 * WAV：wav格式。 * MP3：mp3格式。 * PCM：pcm格式。
	AudioFormat *string `json:"audio_format,omitempty"`

	// 是否需要时间戳。false为不需要，true为需要返回时间戳信息。默认值为false。
	NeedTimestamp *bool `json:"need_timestamp,omitempty"`

	// 是否开启字幕
	GenSrt *bool `json:"gen_srt,omitempty"`

	// 任务类型。 * PRELOAD：预加载任务 * WEBSOCKET：websocket接口任务 * ASYNC_JOB：异步任务 * AUDITION：试听任务
	JobType *string `json:"job_type,omitempty"`

	// 任务状态。 * WAITING：等待中 * PROCESSING：合成中 * FAILED：合成失败 * SUCCEED：合成成功
	JobState *string `json:"job_state,omitempty"`

	// 任务合成错误信息
	FailMsg *string `json:"fail_msg,omitempty"`

	// 任务合成文件列表。
	Files *[]TtsJobFile `json:"files,omitempty"`
}

func (o TtsJobDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TtsJobDetail struct{}"
	}

	return strings.Join([]string{"TtsJobDetail", string(data)}, " ")
}
