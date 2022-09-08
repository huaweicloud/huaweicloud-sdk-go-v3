package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type BusinessCardRequestBody struct {

	// 与url二选一  图像数据，base64编码，要求base64编码后大小不超过10MB。图片最小边不小于15px，最长边不超过8192px，支持JPEG、JPG、PNG、BMP、TIFF格式。  图片文件Base64编码字符串，点击[这里](https://support.huaweicloud.com/ocr_faq/ocr_01_0032.html)查看详细获取方式。
	Image *string `json:"image,omitempty"`

	// 与image二选一  图片的URL路径，目前支持：  - 公网http/https url  - OBS提供的url，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权，详情参见[配置OBS访问权限](https://support.huaweicloud.com/api-ocr/ocr_03_0132.html)。  > 说明：  - 接口响应时间依赖于图片的下载时间，如果图片下载时间过长，会返回接口调用失败。  - 请保证被检测图片所在的存储服务稳定可靠，推荐使用OBS服务存储图片数据。
	Url *string `json:"url,omitempty"`

	// 图片朝向检测开关，可选值包括：  - true：检测图片朝向;  - false：不检测图片朝向。  > 说明：  - 支持任意角度的图片朝向检测。未传入该参数时默认为false，即不检测图片朝向。
	DetectDirection *bool `json:"detect_direction,omitempty"`

	// 返回矫正后的名片图像的BASE64编码的开关，可选值包括：  - true：返回BASE64编码；  - false：不返回BASE64编码。  > 说明：  - 未传入该参数时默认为false，即不返回BASE64编码。
	ReturnAdjustedImage *bool `json:"return_adjusted_image,omitempty"`
}

func (o BusinessCardRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BusinessCardRequestBody struct{}"
	}

	return strings.Join([]string{"BusinessCardRequestBody", string(data)}, " ")
}
