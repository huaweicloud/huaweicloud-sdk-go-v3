package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VideoSynthesisConfigCommon struct {
	Inference *VideoSynthesisInference `json:"inference"`
}

func (o VideoSynthesisConfigCommon) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoSynthesisConfigCommon struct{}"
	}

	return strings.Join([]string{"VideoSynthesisConfigCommon", string(data)}, " ")
}
