package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VideoCuttingConfigCommon struct {
	Inference *VideoCuttingInference `json:"inference"`
}

func (o VideoCuttingConfigCommon) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoCuttingConfigCommon struct{}"
	}

	return strings.Join([]string{"VideoCuttingConfigCommon", string(data)}, " ")
}
