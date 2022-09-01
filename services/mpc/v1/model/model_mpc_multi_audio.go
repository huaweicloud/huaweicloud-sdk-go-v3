package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MpcMultiAudio struct {
	Output *ObsObjInfo `json:"output,omitempty" xml:"output"`

	// 音频文件列表
	AudioFiles *[]AudioFile `json:"audio_files,omitempty" xml:"audio_files"`

	// 输出文件名。
	OutputFilename *string `json:"output_filename,omitempty" xml:"output_filename"`
}

func (o MpcMultiAudio) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MpcMultiAudio struct{}"
	}

	return strings.Join([]string{"MpcMultiAudio", string(data)}, " ")
}
