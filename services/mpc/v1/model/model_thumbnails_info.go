package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ThumbnailsInfo struct {

	// 截图文件信息。
	PicInfo *[]PicInfo `json:"pic_info,omitempty"`

	Output *ObsObjInfo `json:"output,omitempty"`

	// 截图压缩包名。
	OutputName *string `json:"output_name,omitempty"`
}

func (o ThumbnailsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThumbnailsInfo struct{}"
	}

	return strings.Join([]string{"ThumbnailsInfo", string(data)}, " ")
}
