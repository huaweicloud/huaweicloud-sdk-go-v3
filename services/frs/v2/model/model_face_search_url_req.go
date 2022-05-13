package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FaceSearchUrlReq struct {

	// [过滤条件，参考[filter语法](https://support.huaweicloud.com/api-face/face_02_0014.html)。](tag:hc)[过滤条件，参考[filter语法](https://support.huaweicloud.com/intl/zh-cn/api-face/face_02_0014.html)。](tag:hk)
	Filter *string `json:"filter,omitempty"`

	// 返回查询到的最相似的N张人脸，N默认为10。
	TopN *int32 `json:"top_n,omitempty"`

	// [图片的URL路径，目前仅支持华为云上OBS的URL，且人脸识别服务有权限读取该OBS桶的数据。 开通读取权限的操作请参见[服务授权](https://support.huaweicloud.com/api-face/face_02_0006.html)。(tag:hc) [图片的URL路径，目前仅支持华为云上OBS的URL，且人脸识别服务有权限读取该OBS桶的数据。 开通读取权限的操作请参见[服务授权](https://support.huaweicloud.com/intl/zh-cn/api-face/face_02_0006.html)。(tag:hk)
	ImageUrl string `json:"image_url"`

	// 指定返回的自定义字段。
	ReturnFields *[]string `json:"return_fields,omitempty"`

	// 人脸相似度阈值，低于这个阈值则不返回，取值范围0~1，一般情况下建议取值0.93，默认为0。
	Threshold *float64 `json:"threshold,omitempty"`

	// [支持字段排序，参考[sort语法](https://support.huaweicloud.com/api-face/face_02_0013.html)。](tag:hc)[支持字段排序，参考[sort语法](https://support.huaweicloud.com/intl/zh-cn/api-face/face_02_0013.html)。](tag:hk)
	Sort *[]map[string]string `json:"sort,omitempty"`
}

func (o FaceSearchUrlReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FaceSearchUrlReq struct{}"
	}

	return strings.Join([]string{"FaceSearchUrlReq", string(data)}, " ")
}
