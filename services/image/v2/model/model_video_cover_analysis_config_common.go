package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VideoCoverAnalysisConfigCommon struct {
	Inference *VideoCoverAnalysisinference `json:"inference,omitempty"`
}

func (o VideoCoverAnalysisConfigCommon) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoCoverAnalysisConfigCommon struct{}"
	}

	return strings.Join([]string{"VideoCoverAnalysisConfigCommon", string(data)}, " ")
}
