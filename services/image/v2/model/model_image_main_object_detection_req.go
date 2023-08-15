package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageMainObjectDetectionReq This is a auto create Body Object
type ImageMainObjectDetectionReq struct {

	// 与url二选一  图像数据，base64编码，要求base64编码后大小不超过10M，最短边至少1px，最长边最大10000px，支持JPG/PNG/BMP格式。
	Image *string `json:"image,omitempty"`

	// 与image二选一  图片的URL路径，目前支持：  - 公网HTTP/HTTPS URL  - 华为云OBS提供的URL，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权。详请参见[配置OBS服务的访问权限](https://support.huaweicloud.com/api-image/image_03_0037.html)。  > - 接口响应时间依赖于图片的下载时间，如果图片下载时间过长，会返回接口调用失败。 > - 请保证被检测图片所在的存储服务稳定可靠，建议您使用华为云OBS存储。
	Url *string `json:"url,omitempty"`

	// 置信度的阈值（0~100），低于此置信数的检测结果，将不会返回。  默认值：30。  最小值：0。  最大值：100。
	Threshold *float32 `json:"threshold,omitempty"`
}

func (o ImageMainObjectDetectionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageMainObjectDetectionReq struct{}"
	}

	return strings.Join([]string{"ImageMainObjectDetectionReq", string(data)}, " ")
}
