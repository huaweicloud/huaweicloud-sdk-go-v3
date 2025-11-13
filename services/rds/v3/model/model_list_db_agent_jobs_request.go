package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDbAgentJobsRequest Request Object
type ListDbAgentJobsRequest struct {

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 作业类型。默认不传，查询所有作业。 若传，则取值如下: replication:发布订阅相关作业。
	JobType *string `json:"job_type,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为10，不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDbAgentJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbAgentJobsRequest struct{}"
	}

	return strings.Join([]string{"ListDbAgentJobsRequest", string(data)}, " ")
}
