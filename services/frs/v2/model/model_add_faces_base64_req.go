package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddFacesBase64Req struct {

	// 图像数据，Base64编码，要求： • Base64编码后大小不超过8MB，建议小于1MB。 • 图片为JPG/JPEG/BMP/PNG格式。
	ImageBase64 string `json:"image_base64"`

	// 根据用户自定义数据类型，填入相应的数值。 创建faceset时定义该字段，Json字符串不校验重复性，参考[自定义字段](https://support.huaweicloud.com/api-face/face_02_0012.html)。
	ExternalFields *interface{} `json:"external_fields,omitempty"`

	// 用户指定的图片外部ID，与当前图像绑定。用户没提供，系统会生成一个。 该ID长度范围为1～36位，可以包含字母、数字、中划线或者下划线，不包含其他的特殊字符。
	ExternalImageId *string `json:"external_image_id,omitempty"`

	// 是否将图片中的最大人脸添加至人脸库。可选值包括: • true: 传入的单张图片中如果包含多张人脸，则只将最大人脸添加到人脸库中。 • false: 默认为false。传入的单张图片中如果包含多张人脸，则将所有人脸添加至人脸库中。
	Single *bool `json:"single,omitempty"`
}

func (o AddFacesBase64Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddFacesBase64Req struct{}"
	}

	return strings.Join([]string{"AddFacesBase64Req", string(data)}, " ")
}
