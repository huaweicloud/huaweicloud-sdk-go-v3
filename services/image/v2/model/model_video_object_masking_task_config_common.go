package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VideoObjectMaskingTaskConfigCommon struct {
	Inference *VideoObjectMaskingInference `json:"inference"`
}

func (o VideoObjectMaskingTaskConfigCommon) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoObjectMaskingTaskConfigCommon struct{}"
	}

	return strings.Join([]string{"VideoObjectMaskingTaskConfigCommon", string(data)}, " ")
}
