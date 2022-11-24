package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateStreamingJobRequest struct {

	// 是否需要校验配置是否正确
	Check *bool `json:"check,omitempty"`

	// 实时分析作业详细配置，每个作业可选择不同的算子进行组合，各算子的使用方法详见：实时分析算子配置指南。
	Body map[string]interface{} `json:"body,omitempty"`
}

func (o CreateStreamingJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStreamingJobRequest struct{}"
	}

	return strings.Join([]string{"CreateStreamingJobRequest", string(data)}, " ")
}
