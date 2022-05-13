package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateFaceReq struct {

	// [Json字符串不校验重复性，自定义字段的key值长度范围为[1,36]，string类型的value长度范围为[1,256]，具体参见[自定义字段](https://support.huaweicloud.com/api-face/face_02_0012.html)。 这里是待修改的参数，external_image_id和external_fields至少选一个。](tag:hc) [Json字符串不校验重复性，自定义字段的key值长度范围为[1,36]，string类型的value长度范围为[1,256]，具体参见[自定义字段](https://support.huaweicloud.com/intl/zh-cn/api-face/face_02_0012.html)。 这里是待修改的参数，external_image_id和external_fields至少选一个。](tag:hk)
	ExternalFields *interface{} `json:"external_fields,omitempty"`

	// 用户指定的图片外部ID，与当前图像绑定。用户没提供，系统会生成一个。该ID长度范围为1～36位，可以包含字母、数字、中划线或者下划线，不包含其他的特殊字符。 这里是待修改的参数，external_image_id和external_fields至少选一个。
	ExternalImageId *string `json:"external_image_id,omitempty"`

	// 人脸库ID，由系统内部生成的唯一ID。
	FaceId string `json:"face_id"`
}

func (o UpdateFaceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFaceReq struct{}"
	}

	return strings.Join([]string{"UpdateFaceReq", string(data)}, " ")
}
