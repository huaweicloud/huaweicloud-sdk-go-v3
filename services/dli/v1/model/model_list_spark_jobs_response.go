package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSparkJobsResponse Response Object
type ListSparkJobsResponse struct {

	// 参数解释:   起始批处理作业的索引号 示例: 0 约束限制:  无 取值范围: 无 默认取值: 无
	From *int32 `json:"from,omitempty"`

	// 参数解释:   返回批处理作业的总数 示例: 1 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	Total *int32 `json:"total,omitempty"`

	// 批处理作业信息。
	Sessions *[]SparkJob `json:"sessions,omitempty"`

	// 参数解释:   批处理作业的创建时间 示例: 1747169165821 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
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
