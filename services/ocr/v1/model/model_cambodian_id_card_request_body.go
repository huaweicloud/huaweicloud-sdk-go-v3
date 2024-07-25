package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CambodianIdCardRequestBody struct {

	// 与url二选一。图像数据，base64编码。图片尺寸不小于15像素，最长边不超过8192像素，支持JPG/PNG/BMP/TIF格式。
	Image *string `json:"image,omitempty"`

	// 与image二选一。 图片的url路径，目前支持： - 公网http/https url - OBS提供的url，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权，详情参见[配置OBS访问权限](https://support.huaweicloud.com/intl/zh-cn/api-ocr/ocr_03_0132.html)。 > 说明： - 接口响应时间依赖于图片的下载时间，如果图片下载时间过长，会返回接口调用失败。 - 请保证被检测图片所在的存储服务稳定可靠，推荐使用OBS服务存储图片数据。 - url中不能存在中文字符，若存在，中文需要进行utf8编码。
	Url *string `json:"url,omitempty"`

	// 是否返回头像内容开关，可选值如下所示： - true: 返回身份证头像照片的 base64 编码 - false: 不返回身份证头像照片的 base64 编码 未传入该参数时默认为“false”，即不返回身份证头像照片的 base64 编码。
	ReturnPortraitImage *bool `json:"return_portrait_image,omitempty"`

	// 是否返回头像坐标的开关，可选值如下所示： - true: 返回身份证头像的位置坐标 - false: 不返回身份证头像的位置坐标 未传入该参数时默认为“false”，即不返回身份证的头像坐标。
	ReturnPortraitLocation *bool `json:"return_portrait_location,omitempty"`

	// 是否返回身份证类型的开关，可选值如下所示： - true:返回身份证的类型，类型包括身份证原件以及身份证复印件 - false：不返回身份证的类型
	ReturnIdcardType *bool `json:"return_idcard_type,omitempty"`

	// 返回身份证边框完整性的告警结果的开关，可选值如下所示 - true：打开身份证图像边框完整性告警功能  - false：关闭身份证图像边框完整性告警功能
	DetectBorderIntegrity *bool `json:"detect_border_integrity,omitempty"`

	// 返回身份证内部是否有被遮挡的告警结果的开关，可选值如下所示 - true：打开身份证内部是否有被遮挡的告警功能  - false：关闭身份证内部是否有被遮挡的告警功能
	DetectBlockingWithinBorder *bool `json:"detect_blocking_within_border,omitempty"`

	// 返回身份证模糊告警结果的开关，可选值如下所示 - true:打开身份证是否模糊的告警功能 - false：关闭身份证是否模糊的告警功能
	DetectBlur *bool `json:"detect_blur,omitempty"`

	// 返回身份证是否反光的告警结果的开关，可选值如下所示 - true：打开身份证是否反光的告警功能  - false：关闭身份证是否反光的告警功能
	DetectGlare *bool `json:"detect_glare,omitempty"`

	// 返回身份证四点原图的base64编码 - true: 返回身份证原图的base64编码  - false：不返回身份证原图的base64编码
	ReturnAdjustedImage *bool `json:"return_adjusted_image,omitempty"`

	// 返回身份证人像是否被篡改的告警结果的开关，可选值如下所示 - true:  打开身份证人像是否被篡改的告警功能  - false：关闭身份证人像被篡改的告警功能 不支持精细化的P图
	DetectTampering *bool `json:"detect_tampering,omitempty"`

	// 返回判断身份证图像是否经过翻拍告警的开关，可选值如下所示 - true:打开判断身份证图像是否经过翻拍告警的功能  - false:关闭判断身份证图像是否经过翻拍告警的功能
	DetectReproduce *bool `json:"detect_reproduce,omitempty"`
}

func (o CambodianIdCardRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CambodianIdCardRequestBody struct{}"
	}

	return strings.Join([]string{"CambodianIdCardRequestBody", string(data)}, " ")
}
