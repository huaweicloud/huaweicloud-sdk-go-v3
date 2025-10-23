package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VehicleLicenseRequestBody
type VehicleLicenseRequestBody struct {

	// 与url二选一。  图片的Base64编码，要求单个图片其对应的Base64编码不超过10MB。文件在Base64编码后会大于文件原本大小，请注意做好边界判断，建议文件大小不超过7MB。 图片最小边不小于100px，最长边不超过8192px。支持JPEG、JPG、PNG、BMP、TIFF格式。  图片文件Base64编码字符串，点击[这里](https://support.huaweicloud.com/ocr_faq/ocr_01_0032.html)查看详细获取方式。
	Image *string `json:"image,omitempty"`

	// 与image二选一。  单个图片其对应的Base64编码不超过10MB。文件在Base64编码后会大于文件原本大小，请注意做好边界判断，建议文件大小不超过7MB。 图片最小边不小于100px，最长边不超过8192px。支持JPEG、JPG、PNG、BMP、TIFF格式。 图片的URL路径，目前支持： - 公网http/https url - OBS提供的url，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权，详情参见[配置OBS访问权限](https://support.huaweicloud.com/api-ocr/ocr_03_0132.html)。 > 说明： - 接口响应时间依赖于图片的下载时间，如果图片下载时间过长，会返回接口调用失败。 - 请保证被检测图片所在的存储服务稳定可靠，推荐使用OBS服务存储图片数据。 - url中不能存在中文字符，若存在，中文需要进行utf8编码。
	Url *string `json:"url,omitempty"`

	//  - front：行驶证主页  - back：行驶证副页  - double_side：行驶证双页信息  > 说明： 如果参数值为空或无该参数，系统默认识别主页，建议填写，准确率更高。
	Side *string `json:"side,omitempty"`

	// 是否返回发证机关的开关，可选值包括： - true：返回发证机关 - false：不返回发证机关  > 说明： - 如果无该参数，系统默认不返回发证机关。如果输入参数不是Boolean类型，则会报非法参数错误。
	ReturnIssuingAuthority *bool `json:"return_issuing_authority,omitempty"`

	// 识别到的文字块的区域位置信息。取值范围：  - true：返回各个文字块区域  - false：不返回各个文字块区域  如果无该参数，系统默认不返回文字块区域。如果输入参数不是Boolean类型，则会报非法参数错误。
	ReturnTextLocation *bool `json:"return_text_location,omitempty"`

	// 是否支持识别电子行驶证，取值范围：  - true：支持识别电子行驶证  - false：不支持识别电子行驶证  默认不支持识别电子行驶证。如果输入参数不是Boolean类型，则会报非法参数错误。
	RecognizeElectronicLicense *bool `json:"recognize_electronic_license,omitempty"`

	// 是否返回纸质行驶证图像的告警信息，可选值包括： - true：返回纸质行驶证图像的告警信息 - false：不返回纸质行驶证图像的告警信息 如果无该参数，系统默认不返回告警信息。如果输入参数不是Boolean类型，则会报非法参数错误。
	Alarm *bool `json:"alarm,omitempty"`
}

func (o VehicleLicenseRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VehicleLicenseRequestBody struct{}"
	}

	return strings.Join([]string{"VehicleLicenseRequestBody", string(data)}, " ")
}
