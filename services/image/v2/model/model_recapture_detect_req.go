package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RecaptureDetectReq struct {

	// 与url二选一  图像数据，base64编码，要求base64编码后大小不超过10MB，最短边至少15px，最长边最大4096px，支持JPG、PNG、BMP格式。
	Image *string `json:"image,omitempty"`

	// 与image二选一  图片的URL路径，目前支持：  - 公网HTTP/HTTPS URL  - 华为云OBS提供的URL，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权。详请参见[配置OBS服务的访问权限](https://support.huaweicloud.com/api-image/image_03_0037.html)。  > - 接口响应时间依赖于图片的下载时间，如果图片下载时间过长，会返回接口调用失败。 > - 请保证被检测图片所在的存储服务稳定可靠，建议您使用华为云OBS存储。
	Url *string `json:"url,omitempty"`

	// 判断图片真实或虚假的阈值，默认值为0.95，取值范围（0~1），输入非该范围内值算法即取默认值。
	Threshold *float32 `json:"threshold,omitempty"`
}

func (o RecaptureDetectReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecaptureDetectReq struct{}"
	}

	return strings.Join([]string{"RecaptureDetectReq", string(data)}, " ")
}
