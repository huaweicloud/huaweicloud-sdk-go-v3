package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 单个结果文件result.json所在的OBS桶和路径
type TaskOutputHostingForDisplayObs struct {

	// 结果文件result.json所在的OBS桶
	Bucket *string `json:"bucket,omitempty"`

	// 结果文件result.json所在的路径
	Path *string `json:"path,omitempty"`
}

func (o TaskOutputHostingForDisplayObs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskOutputHostingForDisplayObs struct{}"
	}

	return strings.Join([]string{"TaskOutputHostingForDisplayObs", string(data)}, " ")
}
