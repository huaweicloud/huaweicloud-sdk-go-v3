package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SealRequestBody
type SealRequestBody struct {

	// 该参数与url二选一。 图片的Base64编码，要求Base64编码后大小不超过10MB。 图片最短边不小于15px，最长边不超过8192px。支持JPEG、JPG、PNG、BMP、GIF、TIFF、WEBP、PCX、ICO、PSD、PDF格式。 图片Base64编码示例如/9j/4AAQSkZJRgABAg...，带有多余前缀会产生The image format is not supported报错。
	Data *string `json:"data,omitempty"`

	// 与data二选一  要求Base64编码后大小不超过10MB。 图片最短边不小于15px，最长边不超过8192px。支持JPEG、JPG、PNG、BMP、GIF、TIFF、WEBP、PCX、ICO、PSD、PDF格式。 图片的URL路径，目前支持： - 公网http/https url - OBS提供的url，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权，详情参见[配置OBS访问权限](https://support.huaweicloud.com/api-ocr/ocr_03_0132.html)。 > 说明： - 接口响应时间依赖于图片的下载时间，如果图片下载时间过长，会返回接口调用失败。 - 请保证被检测图片所在的存储服务稳定可靠，推荐使用OBS服务存储图片数据。 - url中不能存在中文字符，若存在，中文需要进行utf8编码。
	Url *string `json:"url,omitempty"`

	// 是否返回提取的印章base64编码图片，可选值如下所示： - true：返回提取的印章图片base64编码 - false：不返回提取的印章图片base64编码  如果无该参数，系统默认不返回印章擦除后的图片。如果输入参数不是Boolean类型，则会报非法参数错误。
	ReturnSealImage *bool `json:"return_seal_image,omitempty"`

	// 是否返回印章擦除后的base64编码图片，可选值包括： - true：返回印章擦除后的图片base64编码 - false：不返回印章擦除后的图片base64编码  如果无该参数，系统默认不返回印章擦除后的图片。如果输入参数不是Boolean类型，则会报非法参数错误。
	ReturnErasedSealImage *bool `json:"return_erased_seal_image,omitempty"`

	// 指定PDF页码识别。传入该参数时，则识别指定页码的内容。如果不传该参数，则默认识别第1页。
	PdfPageNumber *int32 `json:"pdf_page_number,omitempty"`
}

func (o SealRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SealRequestBody struct{}"
	}

	return strings.Join([]string{"SealRequestBody", string(data)}, " ")
}
