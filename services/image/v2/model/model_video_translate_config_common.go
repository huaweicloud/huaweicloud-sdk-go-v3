package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VideoTranslateConfigCommon struct {
	Inference *VideoTranslateInference `json:"inference"`
}

func (o VideoTranslateConfigCommon) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoTranslateConfigCommon struct{}"
	}

	return strings.Join([]string{"VideoTranslateConfigCommon", string(data)}, " ")
}
