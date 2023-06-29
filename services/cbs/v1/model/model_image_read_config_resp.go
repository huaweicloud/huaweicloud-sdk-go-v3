package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageReadConfigResp
type ImageReadConfigResp struct {

	// 播报内容，长度为1~2500
	ReadContent string `json:"read_content"`

	// 图片id
	ImageId string `json:"image_id"`

	Resolution *Resolution `json:"resolution"`

	// 图片地址
	ImageUrl string `json:"image_url"`

	// 图片名
	Name string `json:"name"`
}

func (o ImageReadConfigResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageReadConfigResp struct{}"
	}

	return strings.Join([]string{"ImageReadConfigResp", string(data)}, " ")
}
