package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IdCardRequestBody
type IdCardRequestBody struct {

	// 与url二选一  图像数据，base64编码，要求base64编码后大小不超过10MB。图片最小边不小于15px，最长边不超过8000px。支持JPEG、JPG、PNG、BMP、TIFF格式。  图片文件Base64编码字符串，点击[这里](https://support.huaweicloud.com/ocr_faq/ocr_01_0032.html)查看详细获取方式。
	Image *string `json:"image,omitempty"`

	// 与image二选一 图片的URL路径，目前支持： - 公网http/https url - OBS提供的url，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权，详情参见[配置OBS访问权限](https://support.huaweicloud.com/api-ocr/ocr_03_0132.html)。 > 说明： - 接口响应时间依赖于图片的下载时间，如果图片下载时间过长，会返回接口调用失败。 - 请保证被检测图片所在的存储服务稳定可靠，推荐使用OBS服务存储图片数据。 - url中不能存在中文字符，若存在，中文需要进行utf8编码。
	Url *string `json:"url,omitempty"`

	//  - front：身份证人像面。 - back：身份证国徽面。 - double_side：身份证双面信息 > 说明： 如果参数值为空或无该参数，系统自动识别，建议填写，准确率更高。
	Side *string `json:"side,omitempty"`

	// 返回校验身份证号等信息的开关，默认false，可选值如下所示：  - true：返回校验信息  - false：不返回校验信息
	ReturnVerification *bool `json:"return_verification,omitempty"`

	// 识别到的文字块的区域位置信息。可选值包括：  - true：返回各个文字块区域  - false：不返回各个文字块区域  如果无该参数，系统默认不返回文字块区域。如果输入参数不是Boolean类型，则会报非法参数错误。
	ReturnTextLocation *bool `json:"return_text_location,omitempty"`

	// 返回判断身份证图像是否经过翻拍的开关，默认false，可选值如下所示：  - true ：开启判断身份证图像是否经过翻拍功能  - false：关闭判断身份证图像是否经过翻拍功能
	DetectReproduce *bool `json:"detect_reproduce,omitempty"`

	// 返回判断身份证图像是否是黑白复印件的开关，默认false，可选值如下所示：  - true ：开启判断身份证图像是否是复印件功能  - false : 关闭身份证图像是否是复印件功能
	DetectCopy *bool `json:"detect_copy,omitempty"`

	// 返回头像位置信息的开关，默认false，可选值如下所示：  - true ：开启返回头像位置信息的功能 - false : 关闭返回头像位置信息的功能
	ReturnPortraitLocation *bool `json:"return_portrait_location,omitempty"`

	// 返回头像图片信息（base64码）的开关，默认false，可选值如下所示：  - true ：开启头像图片信息（base64码）的功能 - false : 关闭头像图片信息（base64码）的功能
	ReturnPortraitImage *bool `json:"return_portrait_image,omitempty"`

	// 返回身份证卡面（base64码）的开关，默认false，可选值如下所示：  - true ：开启身份证卡面（base64码）的功能 - false : 关闭身份证卡面（base64码）的功能
	ReturnAdjustedImage *bool `json:"return_adjusted_image,omitempty"`

	// 身份证图像PS告警功能开关，默认false，可选值如下：  - true ：开启身份证图像PS告警功能 - false : 关闭身份证图像告警功能
	DetectTampering *bool `json:"detect_tampering,omitempty"`

	// 身份证图像边框完整性告警功能开关，默认false，可选值如下：  - true ：打开身份证图像边框完整性告警功能 - false : 关闭身份证图像边框完整性告警功能
	DetectBorderIntegrity *bool `json:"detect_border_integrity,omitempty"`

	// 身份证图像边框内部是否有异物遮挡的告警功能开关，默认false，可选值如下：  - true ：开启身份证边框内部异物遮挡告警功能 - false : 关闭身份证边框内部异物遮挡告警功能
	DetectBlockingWithinBorder *bool `json:"detect_blocking_within_border,omitempty"`

	// 身份证图像模糊告警功能的开关，默认false，可选值如下：  - true ：开启身份证图像模糊告警功能 - false : 关闭身份证图像模糊告警功能
	DetectBlur *bool `json:"detect_blur,omitempty"`

	// 临时身份证告警功能开关，默认false，可选值如下：  - true ：开启临时身份证告警功能 - false : 关闭临时身份证告警功能
	DetectInterim *bool `json:"detect_interim,omitempty"`

	// 身份证反光告警功能开关，默认false，可选值如下：  - true ：开启身份证反光告警功能  - false : 关闭身份证反光告警功能
	DetectGlare *bool `json:"detect_glare,omitempty"`
}

func (o IdCardRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdCardRequestBody struct{}"
	}

	return strings.Join([]string{"IdCardRequestBody", string(data)}, " ")
}
