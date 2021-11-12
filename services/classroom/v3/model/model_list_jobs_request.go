package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListJobsRequest struct {
	// 作业来源于课堂或课程。 取值范围： classroom:课堂作业 course:课程作业

	SourceFrom string `json:"source_from"`
	// 课堂ID或者课程ID。

	SourceId string `json:"source_id"`
	// 信息记录的起始编号

	Offset *int32 `json:"offset,omitempty"`
	// 每页包含的信息记录数

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobsRequest struct{}"
	}

	return strings.Join([]string{"ListJobsRequest", string(data)}, " ")
}
