package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto crerate Body Object
type ImageWisedesignCombineBody struct {

	// 图层标识，范围5个图层以内, 0 代表背景层
	Id int32 `json:"id"`

	// 图层左上角的横坐标位置，单位为px，默认为0
	X *int32 `json:"x,omitempty"`

	// 图层左上角的纵坐标位置，单位为px，默认为0
	Y *int32 `json:"y,omitempty"`

	// 图层宽度，单位为px，默认为上传图片的宽度
	W *int32 `json:"w,omitempty"`

	// 图层高度，单位为px，默认为上传图片的高度
	H *int32 `json:"h,omitempty"`

	// 是否水平翻转，默认值为false
	Flipx *bool `json:"flipx,omitempty"`

	// 是否垂直翻转，默认值为false
	Flipy *bool `json:"flipy,omitempty"`

	// 图层旋转角度，范围[-180, +180]，默认为0
	Rotate *int32 `json:"rotate,omitempty"`

	// 图层透明度，范围[0, 1]，默认为1
	Opacity *float32 `json:"opacity,omitempty"`

	// 非背景图的图像数据，base64编码，要求base64编码最长边最大3000px，支持JPG/PNG/BMP/JPEG格式
	ImageBase64 *string `json:"image_base64,omitempty"`

	// 与image_base64二选一  图片的URL路径，目前支持：  - 公网HTTP/HTTPS URL  - 华为云OBS提供的URL，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权。详请参见[[配置OBS服务的访问权限](https://support.huaweicloud.com/api-moderation/moderation_03_0020.html)](tag:hc)[[配置OBS服务的访问权限](https://support.huaweicloud.com/intl/zh-cn/api-moderation/moderation_03_0020.html)](tag:hk)。  > - 接口响应时间依赖于图片的下载时间，如果图片下载时间过长，会返回接口调用失败。 > - 请保证被检测图片所在的存储服务稳定可靠，建议您使用华为云OBS存储。 > - lmage不支持跨区域OBS，OBS的区域需要和服务保持一致。
	ImageUrl *string `json:"image_url,omitempty"`

	Backgroundattrs *ImageWisedesignCombineBodyBackgroundattrs `json:"backgroundattrs,omitempty"`
}

func (o ImageWisedesignCombineBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageWisedesignCombineBody struct{}"
	}

	return strings.Join([]string{"ImageWisedesignCombineBody", string(data)}, " ")
}
