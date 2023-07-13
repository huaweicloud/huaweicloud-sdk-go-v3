package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VideoTaggingTaskConfig struct {
	Common *VideoTaggingTaskConfigCommon `json:"common,omitempty"`
}

func (o VideoTaggingTaskConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoTaggingTaskConfig struct{}"
	}

	return strings.Join([]string{"VideoTaggingTaskConfig", string(data)}, " ")
}
