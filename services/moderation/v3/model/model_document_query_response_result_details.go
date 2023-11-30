package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DocumentQueryResponseResultDetails struct {

	// 当前内容片段的处置建议： block：包含敏感信息，不通过 review：需要人工复检
	Suggestion *string `json:"suggestion,omitempty"`

	// 当前内容片段的类型，可取值有： text: 文本 image: 图像 video: 视频
	Type *string `json:"type,omitempty"`

	// 当前内容片段的风险类型： politics：涉政 terrorism：暴恐 porn：色情 sexy：性感 abuse：辱骂 ad：广告 qr_code：二维码 watermark：水印 meaningless：无意义 ban：违禁 bad_scene：不良场景 moan：娇喘
	Label *string `json:"label,omitempty"`

	// 当前处理的片段索引
	Index *int32 `json:"index,omitempty"`

	// 当前内容片段中的文本内容，仅当type为text时存在
	Text *string `json:"text,omitempty"`

	// 当前文本内容片段在输入中的起始位置，仅当type为text时存在
	StartPosition *int32 `json:"start_position,omitempty"`

	// 当前文本内容片段在输入中的结束位置，仅当type为text时存在
	EndPosition *int32 `json:"end_position,omitempty"`

	// 网页图片url，仅当type为image且文档格式为webpage时存在
	ImageUrl *string `json:"image_url,omitempty"`

	// 命中的风险片段信息列表，仅在有命中敏感词时才返回
	Segments *[]DocumentQueryResponseResultSegments `json:"segments,omitempty"`

	// 网页视频中截帧部分审核详情
	VideoImageDetails *[]DocumentVideoImageDetail `json:"video_image_details,omitempty"`

	// 网页视频中音频部分审核详情
	AudioDetails *[]DocumentAudioDetail `json:"audio_details,omitempty"`
}

func (o DocumentQueryResponseResultDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DocumentQueryResponseResultDetails struct{}"
	}

	return strings.Join([]string{"DocumentQueryResponseResultDetails", string(data)}, " ")
}
