package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MvsInvoiceRequestBody
type MvsInvoiceRequestBody struct {

	// 与url二选一。  图片的Base64编码，要求单个图片、PDF文件其对应的Base64编码不超过10MB。文件在Base64编码后会大于文件原本大小，请注意做好边界判断，建议文件大小不超过7MB。 图片最小边不小于100px，最长边不超过8000px。支持JPEG、JPG、PNG、BMP、TIFF、PDF格式。多页PDF仅识别第一页。  图片文件Base64编码字符串，点击[这里](https://support.huaweicloud.com/ocr_faq/ocr_01_0032.html)查看详细获取方式。
	Image *string `json:"image,omitempty"`

	// 与image二选一。  单个图片、PDF文件其对应的Base64编码不超过10MB。文件在Base64编码后会大于文件原本大小，请注意做好边界判断，建议文件大小不超过7MB。 图片最小边不小于100px，最长边不超过8000px。支持JPEG、JPG、PNG、BMP、TIFF、PDF格式。多页PDF仅识别第一页。 图片的URL路径，目前支持： - 公网http/https url - OBS提供的url，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权，详情参见[配置OBS访问权限](https://support.huaweicloud.com/api-ocr/ocr_03_0132.html)。 > 说明： - 接口响应时间依赖于图片的下载时间，如果图片下载时间过长，会返回接口调用失败。 - 请保证被检测图片所在的存储服务稳定可靠，推荐使用OBS服务存储图片数据。 - url中不能存在中文字符，若存在，中文需要进行utf8编码。
	Url *string `json:"url,omitempty"`

	// 如果为True，返回体中会包含text_location对象，内容是各字段的检测框四点坐标。如果是False或者没有这个key，则返回体中不包含text_location对象。
	ReturnTextLocation *bool `json:"return_text_location,omitempty"`

	// 如果为True，返回体中会包含confidence对象，内容是各字段的置信度。如果是False或者没有这个key，则返回体中不包含confidence对象。
	ReturnConfidence *bool `json:"return_confidence,omitempty"`

	// 如果没有type字段则默认返回原机动车销售发票出参； 若存在type字段但是不属于 auto、new或者used三个枚举值，API返回AIS.0101入参错误； 如果type为auto，API自动判断发票类型，并在返回参数中添加type出参以指明发票类型； 如果type为new，API在检测出的类型为机动车发票时返回原版机动车发票出参并添加type出参（“机动车销售统一发票”或“电子发票（机动车销售统一发票）”），不一致时报错AIS.0104图像质量差； 如果type为used，API在检测出的类型为二手车时返回二手车发票出参，并添加type出参（“二手车销售统一发票”或“电子发票（二手车销售统一发票）”），不一致时报错AIS.0104图像质量差。
	Type *string `json:"type,omitempty"`
}

func (o MvsInvoiceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MvsInvoiceRequestBody struct{}"
	}

	return strings.Join([]string{"MvsInvoiceRequestBody", string(data)}, " ")
}
