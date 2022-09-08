package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IdDocumentRequestBody struct {

	// 该参数与url二选一。图像数据，base64编码，要求base64编码后大小不超过10MB。图片最小边不小于15px，最长边不超过8192px，支持JPEG、JPG、PNG、BMP、TIFF。
	Image *string `json:"image,omitempty"`

	// 该参数与image二选一。图片的url路径，目前支持：  Image URL. Currently, the following URLs are supported: - 公网http/https url - OBS提供的url，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权。 > 说明: - 接口响应时间依赖于图片的下载时间，如果图片下载时间过长，会返回接口调用失败。 - 请保证被检测图片所在的存储服务稳定可靠，推荐使用OBS服务存储图片数据。
	Url *string `json:"url,omitempty"`

	// 证件签发国家或地区代码，命名遵循ISO-3166 3位代码。可选值。支持填写1个或多个国家/地区。指定参数后，服务只识别指定国家/地区的卡证，如留空，则识别所有地区卡证。建议国家/地区固定或有限范围的情况下填写。支持国家/地区列表见表1国家/地区和证件列表。
	CountryRegion *[]string `json:"country_region,omitempty"`

	// 证件类型。可选值。支持填写1种或多种证件。指定参数后，服务只识别指定类型的卡证，如留空，默认识别所有类型卡证，建议已知证件类型的情况下填写。支持证件类型如下： - PP: passport，国际护照。 - DL: driving license，驾驶证。 - ID: identification card，各国颁发的身份证类型证件，比如身份证、选民卡、社保卡等。
	IdType *[]string `json:"id_type,omitempty"`

	// 控制是否返回portrait_image（证件中的人像图片），True代表需要返回，False代表不需要。
	ReturnPortraitImage *bool `json:"return_portrait_image,omitempty"`
}

func (o IdDocumentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdDocumentRequestBody struct{}"
	}

	return strings.Join([]string{"IdDocumentRequestBody", string(data)}, " ")
}
