package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PeruIdCardRequestBody struct {

	// 与url二选一。 图像数据，base64编码，要求base64编码后大小不超过10M。图片最小边不小于15像素，最长边不超过8192像素，支持JPG/PNG/BMP/TIFF/PDF格式，支持识别多页PDF数据。
	Image *string `json:"image,omitempty"`

	// 与image二选一。 图片的URL路径，目前仅支持华为云上OBS提供的匿名公开授权访问的URL以及公网URL。
	Url *string `json:"url,omitempty"`

	// 是否返回头像内容开关，可选值如下所示：  - true: 返回身份证头像照片的 base64 编码。 - false:  返回为空值。 未传入该参数时默认为“false”，即返回为空值。
	ReturnPortraitImage *bool `json:"return_portrait_image,omitempty"`

	// 是否返回头像坐标的开关，可选值如下所示：  - true: 返回身份证头像的位置坐标。 - false: 返回为空值 未传入该参数时默认为“false”，即返回为空。
	ReturnPortraitLocation *bool `json:"return_portrait_location,omitempty"`
}

func (o PeruIdCardRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeruIdCardRequestBody struct{}"
	}

	return strings.Join([]string{"PeruIdCardRequestBody", string(data)}, " ")
}
