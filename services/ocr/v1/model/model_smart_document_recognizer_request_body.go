package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmartDocumentRecognizerRequestBody struct {

	// 与url二选一。图片或PDF格式，base64编码，要求base64编码后大小不超过10M。 图像尺寸不小于15×15像素，最长边不超过8192像素，支持JPG/PNG/BMP/TIFF格式。 PDF以144dpi的分辨率转为图像进行文档解析，需符合上述图像尺寸规定。若PDF有多页，当前仅对第1页进行识别。
	Data *string `json:"data,omitempty"`

	// 与data二选一。 图片或PDF的URL路径，目前支持： - 公网http/https url - OBS提供的url，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权，详情参见[配置OBS访问权限](https://support.huaweicloud.com/api-ocr/ocr_03_0132.html)。 > 说明： - 接口响应时间依赖于图片的下载时间，如果图片下载时间过长，会返回接口调用失败。 - 请保证被检测图片所在的存储服务稳定可靠，推荐使用OBS服务存储图片数据。 - url中不能存在中文字符，若存在，中文需要进行utf8编码。
	Url *string `json:"url,omitempty"`

	// 单朝向模式开关。可选值包括： - true：打开单朝向模式 - false：关闭单朝向模式  未传入该参数时默认为false，即默认图片中的字段为多朝向
	SingleOrientationMode *bool `json:"single_orientation_mode,omitempty"`

	// 是否进行键值对（key-value）提取。若是，结果会以“kv_result”这一关键字返回。
	Kv *bool `json:"kv,omitempty"`

	// 是否进行表格识别。此处表格特指逻辑表格，通常具有M行N列的形式，且第一行或第一列为表头。若是，结果会以“table_result”这一关键字返回。
	Table *bool `json:"table,omitempty"`

	// 是否进行版面分析。若是，结果会以“layout_result”这一关键字返回。
	Layout *bool `json:"layout,omitempty"`

	// 仅当table为True时有效。是否返回表格转换Microsoft Excel的Base64编码字段。
	ReturnExcel *bool `json:"return_excel,omitempty"`

	// 是否进行有线表单识别。有线表单指关键信息以有线单元格形式进行呈现，例如户口本、机动车发票等。若是，结果会以\"form_result\"这一关键字返回。
	Form *bool `json:"form,omitempty"`

	// 是否进行公式识别，识别结果为latex序列。若是，结果会以“formula_result”这一关键字返回。 当前仅支持文档（例如论文）中的公式识别，不支持公式切片图像。
	Formula *bool `json:"formula,omitempty"`

	// 需要传入字典的json序列化后字符串，用于对kv_result中的特定key值进行归一化映射。例如，kv_result中包含{\"名称\"：\"小明\"}的键值对，若传入{\"名称\"：\"姓名\"}的kv_map，则返回结果为{“姓名”：“小明”}。  > 参数传入示例： - \"kv_map\":\"{\\\"名称\\\":\\\"姓名\\\"}\"
	KvMap *string `json:"kv_map,omitempty"`

	// 指定PDF页码识别。传入该参数时，则识别指定页码的内容。如果不传该参数，则默认识别第1页。
	PdfPageNumber *int32 `json:"pdf_page_number,omitempty"`
}

func (o SmartDocumentRecognizerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartDocumentRecognizerRequestBody struct{}"
	}

	return strings.Join([]string{"SmartDocumentRecognizerRequestBody", string(data)}, " ")
}
