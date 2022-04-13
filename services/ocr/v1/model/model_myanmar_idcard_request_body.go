package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MyanmarIdcardRequestBody struct {
	// 与url二选一  图像数据，base64编码，要求base64编码后大小不超过10MB。图片最小边不小于15px，最长边不超过4096px，支持JPEG、JPG、PNG、BMP、TIFF格式。

	Image *string `json:"image,omitempty"`
	// 与image二选一 图片的URL路径，目前支持： - 公网http/https url - OBS提供的url，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权，详情参见[配置OBS访问权限](https://support.huaweicloud.com/intl/zh-cn/api-ocr/ocr_03_0132.html)。 > 说明： - 接口响应时间依赖于图片的下载时间，如果图片下载时间过长，会返回接口调用失败。 - 请保证被检测图片所在的存储服务稳定可靠，推荐使用OBS服务存储图片数据。

	Url *string `json:"url,omitempty"`
	// - true：输出为unicode格式 - false：输出为zawgyi格式 如果参数值为空或无该参数，默认输出为zawgyi格式。

	ConvertUnicode *bool `json:"convert_unicode,omitempty"`
	// 是否返回置信度的开关，可选值如下所示。 - true：返回置信度 - false：不返回置信度 如果无该参数，系统默认不返回置信度。如果输入参数不是Boolean类型，则会报非法参数错误。

	ReturnConfidence *bool `json:"return_confidence,omitempty"`
	// 是否返回头像内容开关，可选值如下所示： - true：返回身份证头像照片的 base64 编码 - false：不返回身份证头像照片的 base64 编码

	ReturnPortraitImage *bool `json:"return_portrait_image,omitempty"`
	// 是否返回头像坐标的开关，可选值如下所示： - true：返回身份证头像的位置 - false：不返回身份证头像的位置

	ReturnPortraitLocation *bool `json:"return_portrait_location,omitempty"`
	// 是否返回身份证类型的开关，可选值如下所示： - true：返回身份证的类型，类型包括身份证原件以及身份证复印件 - false：不返回身份证的类型 未传入该参数时默认为false，即不返回身份证头像照片的 base64 编码。

	ReturnIdcardType *bool `json:"return_idcard_type,omitempty"`
}

func (o MyanmarIdcardRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MyanmarIdcardRequestBody struct{}"
	}

	return strings.Join([]string{"MyanmarIdcardRequestBody", string(data)}, " ")
}
