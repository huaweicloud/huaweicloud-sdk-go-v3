package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 背景图层元素的详情，即id为0时填写
type ImageWisedesignCombineBodyBackgroundattrs struct {

	// 背景类型：color-颜色，image-背景图，transparent-透明，默认为transparent-透明
	Backgroundtype *string `json:"backgroundtype,omitempty"`

	// RGB值 如#FFFFFF为白色
	Color *string `json:"color,omitempty"`

	// 背景内容，base64编码，要求base64编码最长边最大3000px，支持JPG/PNG/BMP/JPEG格式
	Image *string `json:"image,omitempty"`

	// 与image二选一  图片的URL路径，目前支持：  - 公网HTTP/HTTPS URL  - 华为云OBS提供的URL，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权。详请参见[[配置OBS服务的访问权限](https://support.huaweicloud.com/api-moderation/moderation_03_0020.html)](tag:hc)[[配置OBS服务的访问权限](https://support.huaweicloud.com/intl/zh-cn/api-moderation/moderation_03_0020.html)](tag:hk)。  > - 接口响应时间依赖于图片的下载时间，如果图片下载时间过长，会返回接口调用失败。 > - 请保证被检测图片所在的存储服务稳定可靠，建议您使用华为云OBS存储。 > - lmage不支持跨区域OBS，OBS的区域需要和服务保持一致。
	Url *string `json:"url,omitempty"`
}

func (o ImageWisedesignCombineBodyBackgroundattrs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageWisedesignCombineBodyBackgroundattrs struct{}"
	}

	return strings.Join([]string{"ImageWisedesignCombineBodyBackgroundattrs", string(data)}, " ")
}
