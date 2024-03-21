package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CustomTemplateRequestBody 自定义模板OCR请求体
type CustomTemplateRequestBody struct {

	// 与url二选一  图像数据，base64编码，要求base64编码后大小不超过10MB。图片最小边不小于15px，最长边不超过8192px，支持JPEG、JPG、PNG、BMP、TIFF、GIF、WEBP格式。  图片文件Base64编码字符串，点击[这里](https://support.huaweicloud.com/ocr_faq/ocr_01_0032.html)查看详细获取方式。
	Image *string `json:"image,omitempty"`

	// 与image二选一 图片的URL路径，目前支持： - 公网http/https url - OBS提供的url，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权，详情参见[配置OBS访问权限](https://support.huaweicloud.com/api-ocr/ocr_03_0132.html)。 > 说明： - 接口响应时间依赖于图片的下载时间，如果图片下载时间过长，会返回接口调用失败。 - 请保证被检测图片所在的存储服务稳定可靠，推荐使用OBS服务存储图片数据。 - url中不能存在中文字符，若存在，中文需要进行utf8编码。
	Url *string `json:"url,omitempty"`

	// 该参数与classifier_id二选一。 模板id，如果传入，启用单模板识别模式。
	TemplateId *string `json:"template_id,omitempty"`

	// 该参数与template_id二选一。 分类器id，如果传入，启用多模板识别模式。
	ClassifierId *string `json:"classifier_id,omitempty"`

	// 该参数与classifier_id参数配合使用，可选值如下所示： - true：仅返回模板分类结果 - false：正常返回多模板识别结果 > 说明： - 如果未传入该参数时默认为false，即正常返回多模板识别结果。
	ClassifierMode *bool `json:"classifier_mode,omitempty"`
}

func (o CustomTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"CustomTemplateRequestBody", string(data)}, " ")
}
