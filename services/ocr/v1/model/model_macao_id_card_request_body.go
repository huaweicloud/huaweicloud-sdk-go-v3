package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MacaoIdCardRequestBody struct {

	// 与url二选一。图像数据，base64编码，要求base64编码后大小不超过10M。图片最小边不小于15像素，最长边不超过8192像素。支持JPG/PNG/BMP/TIFF格式。 图片Base64编码示例 如/9j/4AAQSkZJRgABAg...，带有多余前缀会产生The image format is not supported报错。
	Image *string `json:"image,omitempty"`

	// 与image二选一。 图片的url路径，目前支持： - 公网http/https url - OBS提供的url，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权，详情参见[配置OBS访问权限](https://support.huaweicloud.com/intl/zh-cn/api-ocr/ocr_03_0132.html)>。 > 说明： - 接口响应时间依赖于图片的下载时间，如果图片下载时间过长，会返回接口调用失败。 - 请保证被检测图片所在的存储服务稳定可靠，推荐使用OBS服务存储图片数据。 - url中不能存在中文字符，若存在，中文需要进行utf8编码。
	Url *string `json:"url,omitempty"`

	// front：身份证正面 back：身份证背面  如果参数值为空或无该参数，系统自动识别，建议填写，准确率更高。
	Side *string `json:"side,omitempty"`

	// 是否返回头像内容开关，可选值如下所示： - true: 返回身份证头像照片的 base64 编码 - false: 不返回身份证头像照片的 base64 编码 未传入该参数时默认为“false”，即不返回身份证头像照片的 base64 编码。
	ReturnPortraitImage *bool `json:"return_portrait_image,omitempty"`
}

func (o MacaoIdCardRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MacaoIdCardRequestBody struct{}"
	}

	return strings.Join([]string{"MacaoIdCardRequestBody", string(data)}, " ")
}
