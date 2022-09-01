package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OutputThumbnailPara struct {

	// 抽帧图片张数
	TotalPictures *int32 `json:"total_pictures,omitempty" xml:"total_pictures"`

	// 抽帧图片宽度
	Width *int32 `json:"width,omitempty" xml:"width"`

	// 抽帧图片高度
	Height *int32 `json:"height,omitempty" xml:"height"`

	// 抽帧文件名
	FileName *string `json:"file_name,omitempty" xml:"file_name"`

	Output *ObsObjInfo `json:"output,omitempty" xml:"output"`
}

func (o OutputThumbnailPara) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OutputThumbnailPara struct{}"
	}

	return strings.Join([]string{"OutputThumbnailPara", string(data)}, " ")
}
