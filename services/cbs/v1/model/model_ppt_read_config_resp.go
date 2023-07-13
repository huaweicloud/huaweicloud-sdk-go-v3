package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PptReadConfigResp
type PptReadConfigResp struct {

	// 播报内容，长度为3~2500
	ReadContent string `json:"read_content"`

	// PPT转化有的图片id
	ImageId string `json:"image_id"`

	Resolution *Resolution `json:"resolution"`

	// 图片地址
	ImageUrl string `json:"image_url"`

	// 图片名
	Name string `json:"name"`
}

func (o PptReadConfigResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PptReadConfigResp struct{}"
	}

	return strings.Join([]string{"PptReadConfigResp", string(data)}, " ")
}
