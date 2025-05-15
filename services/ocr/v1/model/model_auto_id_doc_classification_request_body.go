package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AutoIdDocClassificationRequestBody
type AutoIdDocClassificationRequestBody struct {

	// 该参数与url二选一。  图片的Base64编码，单个图片其对应的Base64编码不超过10MB。文件在Base64编码后会大于文件原本大小，请注意做好边界判断，建议图片大小不超过7MB。 图片尺寸不小于15px，最长边不超过8192px，支持JPEG/JPG/PNG/BMP/TIFF格式。 图片Base64编码示例如/9j/4AAQSkZJRgABAg...，带有多余前缀会产生The image format is not supported报错。
	Data *string `json:"data,omitempty"`

	// 与data二选一。  单个图片其对应的Base64编码不超过10MB。文件在Base64编码后会大于文件原本大小，请注意做好边界判断，建议文件大小不超过7MB。 图片尺寸不小于15px，最长边不超过8192px，支持JPEG/JPG/PNG/BMP/TIFF格式。 图片的URL路径，目前支持： - 公网http/https url - OBS提供的url，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权，详情参见[配置OBS访问权限](https://support.huaweicloud.com/api-ocr/ocr_03_0132.html)。 > 说明： - 接口响应时间依赖于图片的下载时间，如果图片下载时间过长，会返回接口调用失败。 - 请保证被检测图片所在的存储服务稳定可靠，推荐使用OBS服务存储图片数据。 - url中不能存在中文字符，若存在，中文需要进行utf8编码。
	Url *string `json:"url,omitempty"`

	// 证件告警检测功能的开关，默认false，取值范围如下：  - true：开启证件图像告警功能。 - false：关闭证件图像告警功能。
	Alarm *bool `json:"alarm,omitempty"`
}

func (o AutoIdDocClassificationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoIdDocClassificationRequestBody struct{}"
	}

	return strings.Join([]string{"AutoIdDocClassificationRequestBody", string(data)}, " ")
}
