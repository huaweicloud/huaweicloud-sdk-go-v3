package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VideoSplitTaskInputData struct {

	// OBS桶名，当输入为obs类型时必填。
	Bucket *string `json:"bucket,omitempty"`

	// OBS的路径，当输入为obs类型时必填。
	Path *string `json:"path,omitempty"`

	// url输入源的地址，当输入为url类型时必填。 长度不超过1000。
	Url *string `json:"url,omitempty"`
}

func (o VideoSplitTaskInputData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoSplitTaskInputData struct{}"
	}

	return strings.Join([]string{"VideoSplitTaskInputData", string(data)}, " ")
}
