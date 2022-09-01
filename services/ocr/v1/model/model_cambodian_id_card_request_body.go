package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CambodianIdCardRequestBody struct {

	// 与url二选一。 图像数据，base64编码。图片尺寸不小于15像素，最长边不超过8192像素，支持JPG/PNG/BMP/TIF格式。
	Image *string `json:"image,omitempty" xml:"image"`

	// 与image二选一。  图片的URL路径，目前仅支持华为云上OBS提供的匿名公开授权访问的URL以及公网URL。
	Url *string `json:"url,omitempty" xml:"url"`

	// 是否返回头像内容开关，可选值如下所示： - true: 返回身份证头像照片的 base64 编码 - false: 不返回身份证头像照片的 base64 编码 未传入该参数时默认为“false”，即不返回身份证头像照片的 base64 编码。
	ReturnPortraitImage *bool `json:"return_portrait_image,omitempty" xml:"return_portrait_image"`

	// 是否返回头像坐标的开关，可选值如下所示： - true: 返回身份证头像的位置坐标 - false: 不返回身份证头像的位置坐标 未传入该参数时默认为“false”，即不返回身份证的头像坐标。
	ReturnPortraitLocation *bool `json:"return_portrait_location,omitempty" xml:"return_portrait_location"`

	// 是否返回身份证类型的开关，可选值如下所示： - true:返回身份证的类型，类型包括身份证原件以及身份证复印件 - false：不返回身份证的类型
	ReturnIdcardType *bool `json:"return_idcard_type,omitempty" xml:"return_idcard_type"`
}

func (o CambodianIdCardRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CambodianIdCardRequestBody struct{}"
	}

	return strings.Join([]string{"CambodianIdCardRequestBody", string(data)}, " ")
}
