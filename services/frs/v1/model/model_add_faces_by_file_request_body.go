package model

import (
	"encoding/json"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"strings"
)

type AddFacesByFileRequestBody struct {
	// 本地图片文件，图片不能超过8MB，建议小于1MB。上传文件时，请求格式为multipart。
	ImageFile *def.FilePart `json:"image_file"`

	// 用户指定的图片外部ID，与当前图像绑定。用户没提供，系统会生成一个。 该ID长度范围为1～36位，可以包含字母、数字、中划线或者下划线，不包含其他的特殊字符。
	ExternalImageId *def.MultiPart `json:"external_image_id,omitempty"`

	// 根据用户自定义数据类型，填入相应的数值。 创建faceset时定义该字段，Json字符串不校验重复性，参考[自定义字段](zh-cn_topic_0130807044.xml)。
	ExternalFields *def.MultiPart `json:"external_fields,omitempty"`
}

func (o AddFacesByFileRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddFacesByFileRequestBody struct{}"
	}

	return strings.Join([]string{"AddFacesByFileRequestBody", string(data)}, " ")
}
