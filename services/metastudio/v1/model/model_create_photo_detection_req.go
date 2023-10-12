package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePhotoDetectionReq struct {

	// 人物照片，需要Base64编码。照片分辨率不超过1080P。
	HumanImage string `json:"human_image"`
}

func (o CreatePhotoDetectionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePhotoDetectionReq struct{}"
	}

	return strings.Join([]string{"CreatePhotoDetectionReq", string(data)}, " ")
}
