package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAudioRecordConfigReq 创建语音录制配置请求。
type CreateAudioRecordConfigReq struct {

	// **参数解释**： 接收语音录制的obs桶名。 **约束限制**： 不涉及 **取值范围**： 字符长度1-64 **默认取值**： 不涉及
	ObsBucketName string `json:"obs_bucket_name"`

	// **参数解释**： 接收语音录制的obs终端节点。 **约束限制**： 需要为obs合法终端节点。 **取值范围**： 字符长度1-64 **默认取值**： 不涉及
	ObsEndpoint string `json:"obs_endpoint"`

	// 语音录制开关
	EnableAudioRecord bool `json:"enable_audio_record"`
}

func (o CreateAudioRecordConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAudioRecordConfigReq struct{}"
	}

	return strings.Join([]string{"CreateAudioRecordConfigReq", string(data)}, " ")
}
