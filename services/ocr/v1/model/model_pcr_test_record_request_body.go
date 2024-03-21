package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PcrTestRecordRequestBody
type PcrTestRecordRequestBody struct {

	// 图像数据，base64编码，图片尺寸不小于15×15像素，最长边不超过8192像素，支持JPG/PNG/BMP/TIFF格式。
	Image *string `json:"image,omitempty"`

	// 与image二选一 图片的URL路径，目前支持： - 公网http/https url - OBS提供的url，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权，详情参见[配置OBS访问权限](https://support.huaweicloud.com/api-ocr/ocr_03_0132.html)。 > 说明： - 接口响应时间依赖于图片的下载时间，如果图片下载时间过长，会返回接口调用失败。 - 请保证被检测图片所在的存储服务稳定可靠，推荐使用OBS服务存储图片数据。 - url中不能存在中文字符，若存在，中文需要进行utf8编码。
	Url *string `json:"url,omitempty"`

	// 校正图片的倾斜角度开关，可选值如下所示： - true：校正图片的倾斜角度 - false：不校正图片的倾斜角度  支持任意角度的校正，未传入该参数时默认为“false”。
	DetectDirection *bool `json:"detect_direction,omitempty"`
}

func (o PcrTestRecordRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PcrTestRecordRequestBody struct{}"
	}

	return strings.Join([]string{"PcrTestRecordRequestBody", string(data)}, " ")
}
