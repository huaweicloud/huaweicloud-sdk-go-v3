package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VideoTaggingTaskConfigCommon struct {
	Inference *VideoTagginginference `json:"inference"`
}

func (o VideoTaggingTaskConfigCommon) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoTaggingTaskConfigCommon struct{}"
	}

	return strings.Join([]string{"VideoTaggingTaskConfigCommon", string(data)}, " ")
}
