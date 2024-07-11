package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PeruIdCardRequestBody struct {

	// 与url二选一。 图像数据，base64编码，要求base64编码后大小不超过10M。图片最小边不小于15像素，最长边不超过8192像素，支持JPG/PNG/BMP/TIFF/PDF格式，支持识别多页PDF数据。
	Image *string `json:"image,omitempty"`

	// 与image二选一。 图片的url路径，目前支持： - 公网http/https url - OBS提供的url，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权，详情参见[配置OBS访问权限](https://support.huaweicloud.com/intl/zh-cn/api-ocr/ocr_03_0132.html)>。 > 说明： - 接口响应时间依赖于图片的下载时间，如果图片下载时间过长，会返回接口调用失败。 - 请保证被检测图片所在的存储服务稳定可靠，推荐使用OBS服务存储图片数据。 - url中不能存在中文字符，若存在，中文需要进行utf8编码。
	Url *string `json:"url,omitempty"`

	// 是否返回头像内容开关，可选值如下所示：  - true: 返回身份证头像照片的 base64 编码。 - false:  返回为空值。 未传入该参数时默认为“false”，即返回为空值。
	ReturnPortraitImage *bool `json:"return_portrait_image,omitempty"`

	// 是否返回头像坐标的开关，可选值如下所示：  - true: 返回身份证头像的位置坐标。 - false: 返回为空值 未传入该参数时默认为“false”，即返回为空。
	ReturnPortraitLocation *bool `json:"return_portrait_location,omitempty"`
}

func (o PeruIdCardRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeruIdCardRequestBody struct{}"
	}

	return strings.Join([]string{"PeruIdCardRequestBody", string(data)}, " ")
}
