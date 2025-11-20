package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDbAgentJobHistorysRequest Request Object
type ListDbAgentJobHistorysRequest struct {

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 作业ID
	JobId string `json:"job_id"`

	// 作业执行状态。默认不传，查询所有执行历史。 若传，则取值如下: failed:失败。 succeeded:成功。 retrying:重试中。 canceled:已取消。 in_progress:正在运行。
	RunStatus *string `json:"run_status,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为10，不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDbAgentJobHistorysRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbAgentJobHistorysRequest struct{}"
	}

	return strings.Join([]string{"ListDbAgentJobHistorysRequest", string(data)}, " ")
}
