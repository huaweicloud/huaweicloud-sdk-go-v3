package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type PcrTestRecordRequestBody struct {

	// 图像数据，base64编码，图片尺寸不小于15×15像素，最长边不超过8192像素，支持JPG/PNG/BMP/TIFF格式。
	Image *string `json:"image,omitempty"`

	// 与image二选一  图片的URL路径，目前仅支持华为云上OBS提供的匿名公开授权访问的URL以及公网URL。
	Url *string `json:"url,omitempty"`

	// 校正图片的倾斜角度开关，可选值如下所示：  - true：校正图片的倾斜角度  - false：不校正图片的倾斜角度  支持任意角度的校正，未传入该参数时默认为“false”。
	DetectDirection *bool `json:"detect_direction,omitempty"`
}

func (o PcrTestRecordRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PcrTestRecordRequestBody struct{}"
	}

	return strings.Join([]string{"PcrTestRecordRequestBody", string(data)}, " ")
}
