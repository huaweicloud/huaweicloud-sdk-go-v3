package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FinancialStatementRequestBody
type FinancialStatementRequestBody struct {

	// 与url二选一。  图片的Base64编码，要求单个图片其对应的Base64编码不超过10MB。文件在Base64编码后会大于文件原本大小，请注意做好边界判断，建议文件大小不超过7MB。 图片最小边不小于15px，最长边不超过8192px，支持JPEG、JPG、PNG、BMP、TIFF格式。  图片文件Base64编码字符串，点击[这里](https://support.huaweicloud.com/ocr_faq/ocr_01_0032.html)查看详细获取方式。
	Image *string `json:"image,omitempty"`

	// 与image二选一。  单个图片其对应的Base64编码不超过10MB。文件在Base64编码后会大于文件原本大小，请注意做好边界判断，建议文件大小不超过7MB。 图片最小边不小于15px，最长边不超过8192px，支持JPEG、JPG、PNG、BMP、TIFF格式。 图片的URL路径，目前支持： - 公网http/https url - OBS提供的url，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权，详情参见[配置OBS访问权限](https://support.huaweicloud.com/api-ocr/ocr_03_0132.html)。 > 说明： - 接口响应时间依赖于图片的下载时间，如果图片下载时间过长，会返回接口调用失败。 - 请保证被检测图片所在的存储服务稳定可靠，推荐使用OBS服务存储图片数据。 - url中不能存在中文字符，若存在，中文需要进行utf8编码。
	Url *string `json:"url,omitempty"`

	// 返回文本块坐标及单元格坐标信息，可选值包括： - true：返回文本块和单元格坐标; - false：不返回。  > 说明： - 未传入该参数时默认为false，即不返回。
	ReturnTextLocation *bool `json:"return_text_location,omitempty"`

	// 返回字段识别置信度，小数点后四位。可选值包括： - true：返回字段置信度; - false：不返回。  > 说明： - 未传入该参数时默认为false，即不返回字段置信度。
	ReturnConfidence *bool `json:"return_confidence,omitempty"`

	// 是否返回表格转换Microsoft Excel的base64编码字段。可选值包括： - true：返回’excel’字段，表示xlsx格式的表格识别结果的base64编码; - false：不返回。  > 说明： - 对返回的Excel编码，可用Python函数 base64.b64decode解码后保存为xlsx文件。
	ReturnExcel *bool `json:"return_excel,omitempty"`

	// 返回表格坐标，可选值包括： - true：返回表格坐标; - false：不返回。  > 说明： - 未传入该参数时默认为false，即不返回。
	ReturnTableLocation *bool `json:"return_table_location,omitempty"`

	// 返回矫正后的图像大小，可选值包括： - true：返回矫正图像大小; - false：不返回。  > 说明： - 未传入该参数时默认为false，即不返回。
	ReturnImageSize *bool `json:"return_image_size,omitempty"`

	// 透视变换矩阵，可选值包括： - true：返回透视变换矩阵; - false：不返回。  > 说明： - 未传入该参数时默认为false，即不返回透视变换矩阵。
	ReturnRectificationMatrix *bool `json:"return_rectification_matrix,omitempty"`
}

func (o FinancialStatementRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FinancialStatementRequestBody struct{}"
	}

	return strings.Join([]string{"FinancialStatementRequestBody", string(data)}, " ")
}
