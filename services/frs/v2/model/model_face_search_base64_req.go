package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FaceSearchBase64Req struct {

	// [过滤条件，参考[filter语法](https://support.huaweicloud.com/api-face/face_02_0014.html)。](tag:hc) [过滤条件，参考[filter语法](https://support.huaweicloud.com/intl/zh-cn/api-face/face_02_0014.html)。](tag:hk)
	Filter *string `json:"filter,omitempty" xml:"filter"`

	// 返回查询到的最相似的N张人脸，N默认为10。
	TopN *int32 `json:"top_n,omitempty" xml:"top_n"`

	// 图像数据，Base64编码，要求： • Base64编码后大小不超过8MB，建议小于MB。 • 图片为JPG/JPEG/BMP/PNG格式。
	ImageBase64 string `json:"image_base64" xml:"image_base64"`

	// 指定返回的自定义字段。
	ReturnFields *[]string `json:"return_fields,omitempty" xml:"return_fields"`

	// 人脸相似度阈值，低于这个阈值则不返回，取值范围0~1，一般情况下建议取值0.93，默认为0。
	Threshold *float64 `json:"threshold,omitempty" xml:"threshold"`

	// [支持字段排序，参考[sort语法](https://support.huaweicloud.com/api-face/face_02_0013.html)。](tag:hc) [支持字段排序，参考[sort语法](https://support.huaweicloud.com/intl/zh-cn/api-face/face_02_0013.html)。](tag:hk)
	Sort *[]map[string]string `json:"sort,omitempty" xml:"sort"`
}

func (o FaceSearchBase64Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FaceSearchBase64Req struct{}"
	}

	return strings.Join([]string{"FaceSearchBase64Req", string(data)}, " ")
}
