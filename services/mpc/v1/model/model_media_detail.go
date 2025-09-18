package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MediaDetail struct {

	// 任务名称
	Features *[]string `json:"features,omitempty"`

	OriginPara *OriginPara `json:"origin_para,omitempty"`

	// 多路输出片源信息
	OutputVideoParas *[]OutputVideoPara `json:"output_video_paras,omitempty"`

	// 被替换的子索引文件
	ReplaceSubIndex *[]string `json:"replace_sub_index,omitempty"`

	OutputThumbnailPara *OutputThumbnailPara `json:"output_thumbnail_para,omitempty"`

	OutputWatermarkParas *OutputWatermarkPara `json:"output_watermark_paras,omitempty"`

	// 输出文件的metadata信息
	Metadata *[]FileMetaData `json:"metadata,omitempty"`
}

func (o MediaDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MediaDetail struct{}"
	}

	return strings.Join([]string{"MediaDetail", string(data)}, " ")
}
