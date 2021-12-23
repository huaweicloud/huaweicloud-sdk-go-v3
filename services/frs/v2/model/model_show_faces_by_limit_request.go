package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowFacesByLimitRequest struct {
	// 人脸库名称。

	FaceSetName string `json:"face_set_name"`
	// 从第几条数据读起，默认为0。

	Offset int32 `json:"offset"`
	// 读取多少条，默认为5。

	Limit int32 `json:"limit"`
}

func (o ShowFacesByLimitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFacesByLimitRequest struct{}"
	}

	return strings.Join([]string{"ShowFacesByLimitRequest", string(data)}, " ")
}
