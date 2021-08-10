package model

import (
	"encoding/json"

	"strings"
)

type FaceSearchBase64Req struct {
	// 过滤条件，参考[filter语法](zh-cn_topic_0130807048.xml)。

	Filter *string `json:"filter,omitempty"`
	// 返回查询到的最相似的N张人脸，N默认为10。

	TopN *int32 `json:"top_n,omitempty"`
	// 图像数据，Base64编码，要求： • Base64编码后大小不超过8MB，建议小于MB。 • 图片为JPG/JPEG/BMP/PNG格式。

	ImageBase64 string `json:"image_base64"`
	// 指定返回的自定义字段。

	ReturnFields *[]string `json:"return_fields,omitempty"`
	// 人脸相似度阈值，低于这个阈值则不返回，取值范围0~1，一般情况下建议取值0.93，默认为0。

	Threshold *float64 `json:"threshold,omitempty"`
	// 支持字段排序，参考[sort语法](zh-cn_topic_0130807047.xml)。

	Sort *[]map[string]string `json:"sort,omitempty"`
}

func (o FaceSearchBase64Req) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "FaceSearchBase64Req struct{}"
	}

	return strings.Join([]string{"FaceSearchBase64Req", string(data)}, " ")
}
