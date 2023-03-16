package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VideoObjectMaskingTaskConfig struct {
	Common *VideoObjectMaskingTaskConfigCommon `json:"common"`
}

func (o VideoObjectMaskingTaskConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoObjectMaskingTaskConfig struct{}"
	}

	return strings.Join([]string{"VideoObjectMaskingTaskConfig", string(data)}, " ")
}
