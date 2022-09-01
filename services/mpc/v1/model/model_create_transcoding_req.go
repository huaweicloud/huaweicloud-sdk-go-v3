package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTranscodingReq struct {
	Input *ObsObjInfo `json:"input,omitempty" xml:"input"`

	Output *ObsObjInfo `json:"output" xml:"output"`

	// 转码模板ID，没带av_parameter参数时，必须带该参数，数组，每一路转码输出对应一个转码配置模板ID，最多支持9个模板ID。  多个转码模板中如下参数可变，其他都必须一致：  - 视频bitrate，height，width。
	TransTemplateId *[]int32 `json:"trans_template_id,omitempty" xml:"trans_template_id"`

	// 转码参数。  若同时设置“trans_template_id”和此参数，则优先使用此参数进行转码，不带trans_template_id时，该参数必选。
	AvParameters *[]AvParameters `json:"av_parameters,omitempty" xml:"av_parameters"`

	// 输出文件名称，每一路转码输出对应一个名称，需要与转码模板ID数组的顺序对应。  - 若设置该参数，表示输出文件按该参数命名。 - 若不设置该参数，表示输出文件按默认方式命名。
	OutputFilenames *[]string `json:"output_filenames,omitempty" xml:"output_filenames"`

	// 用户自定义数据，该字段可在查询接口或消息通知中按原内容透传给用户。
	UserData *string `json:"user_data,omitempty" xml:"user_data"`

	// 图片水印参数，数组，最多支持20个成员。
	Watermarks *[]WatermarkRequest `json:"watermarks,omitempty" xml:"watermarks"`

	Thumbnail *Thumbnail `json:"thumbnail,omitempty" xml:"thumbnail"`

	// 任务优先级，取值如下： - 9代表高优先级。 - 6代表中优先级，默认为6。  暂时只支持6和9。
	Priority *int32 `json:"priority,omitempty" xml:"priority"`

	Subtitle *Subtitle `json:"subtitle,omitempty" xml:"subtitle"`

	Encryption *Encryption `json:"encryption,omitempty" xml:"encryption"`

	Crop *Crop `json:"crop,omitempty" xml:"crop"`

	AudioTrack *AudioTrack `json:"audio_track,omitempty" xml:"audio_track"`

	MultiAudio *MultiAudio `json:"multi_audio,omitempty" xml:"multi_audio"`

	VideoProcess *VideoProcess `json:"video_process,omitempty" xml:"video_process"`

	AudioProcess *AudioProcess `json:"audio_process,omitempty" xml:"audio_process"`
}

func (o CreateTranscodingReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTranscodingReq struct{}"
	}

	return strings.Join([]string{"CreateTranscodingReq", string(data)}, " ")
}
