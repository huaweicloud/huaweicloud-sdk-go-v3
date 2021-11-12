package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type InsurancePolicyRequestBody struct {
	// 图像数据，base64编码，要求base64编码后大小不超过10MB。图片最小边不小于15px，最长边不超过8192px，支持JPEG、JPG、PNG、BMP、TIFF格式。  图片文件Base64编码字符串，点击[这里](https://support.huaweicloud.com/ocr_faq/ocr_01_0032.html)查看详细获取方式。

	Image string `json:"image"`
	// 为Boolean类型，若不传该字段，默认不检测图像倾斜角度文字方向，为True时，会检测倾斜角度并矫正识别

	DetectDirection *bool `json:"detect_direction,omitempty"`
}

func (o InsurancePolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InsurancePolicyRequestBody struct{}"
	}

	return strings.Join([]string{"InsurancePolicyRequestBody", string(data)}, " ")
}
