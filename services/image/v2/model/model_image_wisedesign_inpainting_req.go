package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto crerate Body Object
type ImageWisedesignInpaintingReq struct {

	// 图像数据，base64编码，要求base64编码最长边最大4000px，支持JPG/PNG/BMP/JPEG格式
	ImageBase64 *string `json:"image_base64,omitempty"`

	// 与image_base64二选一   图片的URL路径，目前支持：   - 公网HTTP/HTTPS URL   - 华为云OBS提供的URL，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权。详请参见[[配置OBS服务的访问权限](https://support.huaweicloud.com/api-moderation/moderation_03_0020.html)](tag:hc)[[配置OBS服务的访问权限](https://support.huaweicloud.com/intl/zh-cn/api-moderation/moderation_03_0020.html)](tag:hk)。   > - 接口响应时间依赖于图片的下载时间，如果图片下载时间过长，会返回接口调用失败。  > - 请保证被检测图片所在的存储服务稳定可靠，建议您使用华为云OBS存储。  > - lmage不支持跨区域OBS，OBS的区域需要和服务保持一致。
	ImageUrl *string `json:"image_url,omitempty"`

	// 指定的多个待修复区域信息，每个待修复区域的类型为list[list[list[int,int]]]，待修复区域至少由3个顺时针的坐标构成。若该参数为空，则会自动识别图像中的文字部分进行修复。默认为空
	PolygonCoord *[][][]int32 `json:"polygon_coord,omitempty"`
}

func (o ImageWisedesignInpaintingReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageWisedesignInpaintingReq struct{}"
	}

	return strings.Join([]string{"ImageWisedesignInpaintingReq", string(data)}, " ")
}
