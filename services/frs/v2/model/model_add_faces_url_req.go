package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddFacesUrlReq struct {
	// 图片的URL路径，目前仅支持华为云上OBS的URL，且人脸识别服务有权限读取该OBS桶的数据。 开通读取权限的操作请参见服务授权。

	ImageUrl string `json:"image_url"`
	// 根据用户自定义数据类型，填入相应的数值。 创建faceset时定义该字段，Json字符串不校验重复性，参考自定义字段。

	ExternalFields *interface{} `json:"external_fields,omitempty"`
	// 用户指定的图片外部ID，与当前图像绑定。用户没提供，系统会生成一个。 该ID长度范围为1～36位，可以包含字母、数字、中划线或者下划线，不包含其他的特殊字符。

	ExternalImageId *string `json:"external_image_id,omitempty"`
	// 是否将图片中的最大人脸添加至人脸库。可选值包括: • true: 传入的单张图片中如果包含多张人脸，则只将最大人脸添加到人脸库中。 • false: 默认为false。传入的单张图片中如果包含多张人脸，则将所有人脸添加至人脸库中。

	Single *bool `json:"single,omitempty"`
}

func (o AddFacesUrlReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddFacesUrlReq struct{}"
	}

	return strings.Join([]string{"AddFacesUrlReq", string(data)}, " ")
}
