package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowJobByIdResponse struct {

	// 实时分析作业详细配置，每个作业可选择不同的算子进行组合，各算子的使用方法详见：实时分析算子配置指南。
	JobConfig map[string]interface{} `json:"job_config,omitempty" xml:"job_config"`

	JobInfo        *StreamingJobInfoDto `json:"job_info,omitempty" xml:"job_info"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowJobByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowJobByIdResponse", string(data)}, " ")
}
