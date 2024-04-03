package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSparkJobsResponse Response Object
type ListSparkJobsResponse struct {

	// 起始批处理作业的索引号。
	From *int32 `json:"from,omitempty"`

	// 返回批处理作业的总数。
	Total *int32 `json:"total,omitempty"`

	// 批处理作业信息。
	Sessions *[]SparkJob `json:"sessions,omitempty"`

	// 批处理作业的创建时间。
	CreateTime     *int64 `json:"create_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSparkJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSparkJobsResponse struct{}"
	}

	return strings.Join([]string{"ListSparkJobsResponse", string(data)}, " ")
}
