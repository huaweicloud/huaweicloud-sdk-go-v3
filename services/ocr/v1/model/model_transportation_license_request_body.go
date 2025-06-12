package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TransportationLicenseRequestBody
type TransportationLicenseRequestBody struct {

	// 与url二选一。  图片的Base64编码，要求单个图片其对应的Base64编码不超过10MB。文件在Base64编码后会大于文件原本大小，请注意做好边界判断，建议文件大小不超过7MB。 图片最小边不小于15px，最长边不超过4096px，支持JPEG、JPG、PNG、BMP、TIFF格式。  图片文件Base64编码字符串，点击[这里](https://support.huaweicloud.com/ocr_faq/ocr_01_0032.html)查看详细获取方式。
	Image *string `json:"image,omitempty"`

	// 与image二选一。  单个图片其对应的Base64编码不超过10MB。文件在Base64编码后会大于文件原本大小，请注意做好边界判断，建议文件大小不超过7MB。 图片最小边不小于15px，最长边不超过4096px，支持JPEG、JPG、PNG、BMP、TIFF格式。 图片的URL路径，目前支持： - 公网http/https url - OBS提供的url，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权，详情参见[配置OBS访问权限](https://support.huaweicloud.com/api-ocr/ocr_03_0132.html)。 > 说明： - 接口响应时间依赖于图片的下载时间，如果图片下载时间过长，会返回接口调用失败。 - 请保证被检测图片所在的存储服务稳定可靠，推荐使用OBS服务存储图片数据。 - url中不能存在中文字符，若存在，中文需要进行utf8编码。
	Url *string `json:"url,omitempty"`

	// 返回道路运输证在原图中的坐标位置的开关，默认false，取值范围包括：  - true: 开启返回坐标位置的功能。 - false: 关闭返回坐标位置的功能。
	ReturnImageLocation *bool `json:"return_image_location,omitempty"`

	// 返回道路运输证（base64）的开关，默认false，取值范围包括：  - true: 开启道路运输证（base64）的功能。 - false: 关闭道路运输证（base64）的功能。
	ReturnAdjustedImage *bool `json:"return_adjusted_image,omitempty"`
}

func (o TransportationLicenseRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransportationLicenseRequestBody struct{}"
	}

	return strings.Join([]string{"TransportationLicenseRequestBody", string(data)}, " ")
}
