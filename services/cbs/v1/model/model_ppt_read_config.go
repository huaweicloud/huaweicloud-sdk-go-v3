package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PptReadConfig
type PptReadConfig struct {

	// 播报内容，长度为3~2500
	ReadContent string `json:"read_content"`

	// PPT转化有的图片id
	ImageId string `json:"image_id"`

	Resolution *Resolution `json:"resolution"`
}

func (o PptReadConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PptReadConfig struct{}"
	}

	return strings.Join([]string{"PptReadConfig", string(data)}, " ")
}
