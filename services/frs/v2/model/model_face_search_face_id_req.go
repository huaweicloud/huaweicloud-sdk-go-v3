package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FaceSearchFaceIdReq struct {
	// 过滤条件，参考[filter语法](https://support.huaweicloud.com/api-face/face_02_0014.html)。

	Filter *string `json:"filter,omitempty"`
	// 返回查询到的最相似的N张人脸，N默认为10。

	TopN *int32 `json:"top_n,omitempty"`
	// 导入人脸时，系统返回的人脸编号，即人脸ID。

	FaceId string `json:"face_id"`
	// 指定返回的自定义字段。

	ReturnFields *[]string `json:"return_fields,omitempty"`
	// 人脸相似度阈值，低于这个阈值则不返回，取值范围0~1，一般情况下建议取值0.93，默认为0。

	Threshold *float64 `json:"threshold,omitempty"`
	// 支持字段排序，参考[sort语法](https://support.huaweicloud.com/api-face/face_02_0013.html)。

	Sort *[]map[string]string `json:"sort,omitempty"`
}

func (o FaceSearchFaceIdReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FaceSearchFaceIdReq struct{}"
	}

	return strings.Join([]string{"FaceSearchFaceIdReq", string(data)}, " ")
}
