package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageToVideoConfigCommon struct {
	Inference *ImageToVideoInference `json:"inference"`
}

func (o ImageToVideoConfigCommon) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageToVideoConfigCommon struct{}"
	}

	return strings.Join([]string{"ImageToVideoConfigCommon", string(data)}, " ")
}
