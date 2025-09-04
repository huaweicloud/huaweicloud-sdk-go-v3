package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobStatusResultResponseBodyResult 结果
type JobStatusResultResponseBodyResult struct {

	// 是否构建中
	Building *bool `json:"building,omitempty"`

	// 构建结果
	BuildResult *string `json:"build_result,omitempty"`

	// region
	Region *string `json:"region,omitempty"`
}

func (o JobStatusResultResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobStatusResultResponseBodyResult struct{}"
	}

	return strings.Join([]string{"JobStatusResultResponseBodyResult", string(data)}, " ")
}
