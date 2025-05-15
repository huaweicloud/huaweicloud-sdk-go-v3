package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GeneralTextRequestBody
type GeneralTextRequestBody struct {

	// 与url二选一。  图片的Base64编码，要求单个图片、PDF文件其对应的Base64编码不超过10MB。文件在Base64编码后会大于文件原本大小，请注意做好边界判断，建议文件大小不超过7MB。 图片最小边不小于15px，最长边不超过30000px。图片分辨率不能大于1.6亿px。 支持JPEG、JPG、PNG、BMP、GIF、TIFF、WEBP、PCX、ICO、PSD、PDF格式。  图片文件Base64编码字符串，[点击[这里](https://support.huaweicloud.com/ocr_faq/ocr_01_0032.html)](tag:hc)[点击[这里](https://support.huaweicloud.com/intl/zh-cn/ocr_faq/ocr_01_0032.html)](tag:hk)查看详细获取方式。
	Image *string `json:"image,omitempty"`

	// 与image二选一。  单个图片、PDF文件其对应的Base64编码不超过10MB。文件在Base64编码后会大于文件原本大小，请注意做好边界判断，建议文件大小不超过7MB。 图片最小边不小于15px，最长边不超过30000px。图片分辨率不能大于1.6亿px。 支持JPEG、JPG、PNG、BMP、GIF、TIFF、WEBP、PCX、ICO、PSD、PDF格式。 图片的URL路径，目前支持： - 公网http/https url - OBS提供的url，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权，[详情参见[配置OBS访问权限](https://support.huaweicloud.com/api-ocr/ocr_03_0132.html)。](tag:hc)[详情参见[配置OBS访问权限](https://support.huaweicloud.com/intl/zh-cn/api-ocr/ocr_03_0132.html)。](tag:hk)  > 说明： - 接口响应时间依赖于图片的下载时间，如果图片下载时间过长，会返回接口调用失败。 - 请保证被检测图片所在的存储服务稳定可靠，推荐使用OBS服务存储图片数据。 - url中不能存在中文字符，若存在，中文需要进行utf8编码。
	Url *string `json:"url,omitempty"`

	// 图片朝向检测开关，可选值包括： - true：打开检测图片朝向功能。 - false：关闭检测图片朝向功能。  > 说明： - 支持任意角度的图片朝向检测。未传入该参数时默认为false，即不检测图片朝向。
	DetectDirection *bool `json:"detect_direction,omitempty"`

	// 快速模式开关，针对单行文字图片（要求图片只包含一行文字，且文字区域占比超过50%），打开时可以更快返回识别。可选值包括： - true：打开快速识别模式。 - false：关闭快速识别模式。  > 说明： - 未传入该参数时默认为false，即关闭快速模式。
	QuickMode *bool `json:"quick_mode,omitempty"`

	// 单字符模式开关。可选值包括： - true：打开单字符模式。 - false：关闭单字符模式。  未传入该参数时默认为false，即不返回单个文本行的单字符信息。
	CharacterMode *bool `json:"character_mode,omitempty"`

	// 语种选择，取值范围可参考表1中英文列。未传入该参数时默认为中英文识别模式。 **表1* 语种选择说明 | 英文 |     中文     | | :--: | :----------: | | auto | 自动语种分类 | |  ms  |    马来语    | |  uk  |   乌克兰语   | |  hi  |    印地语    | |  ru  |     俄语     | |  vi  |    越南语    | |  id  |    印尼语    | |  th  |     泰语     | |  zh  |    中英文    | |  ar  |   阿拉伯语   | |  de  |     德语     | |  la  |    拉丁语    | |  fr  |     法语     | |  it  |   意大利语   | |  es  |   西班牙语   | |  pt  |   葡萄牙语   | |  ro  |  罗马尼亚语  | |  pl  |    波兰语    | |  am  |  阿姆哈拉语  | |  ja  |     日语     | |  ko  |     韩语     | |  tr  |   土耳其语   | |  no  |   挪威语     | |  da  |   丹麦语     | |  sv  |   瑞典语     | |  km  |   柬埔寨语   | |  he  |   希伯来语   |
	Language *string `json:"language,omitempty"`

	// 单朝向模式开关。可选值包括： - true：打开单朝向模式。 - false：关闭单朝向模式。 图片文字方向一致时，打开该开关可提升识别精度；图片文字方向不一致时，关闭该开关可支持多朝向文字识别。未传入该参数时默认为false，即默认图片中的文字朝向为多朝向。
	SingleOrientationMode *bool `json:"single_orientation_mode,omitempty"`

	// 指定PDF页码识别。传入该参数时，则识别指定页码的内容。如果不传该参数，则默认识别第1页。
	PdfPageNumber *int32 `json:"pdf_page_number,omitempty"`
}

func (o GeneralTextRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GeneralTextRequestBody struct{}"
	}

	return strings.Join([]string{"GeneralTextRequestBody", string(data)}, " ")
}
