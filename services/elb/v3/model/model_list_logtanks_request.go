package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListLogtanksRequest struct {
	// 每页返回的个数。

	Limit *int32 `json:"limit,omitempty"`
	// 上一页最后一条记录的ID。  使用说明：  - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。

	Marker *string `json:"marker,omitempty"`
	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。  使用说明： - 必须与limit一起使用。

	PageReverse *bool `json:"page_reverse,omitempty"`
	// 企业项目ID。 支持多值查询，查询条件格式：enterprise_project_id=xxx&enterprise_project_id=xxx。

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
	// 云日志记录ID。 支持多值查询，查询条件格式：id=xxx&id=xxx。

	Id *[]string `json:"id,omitempty"`
	// 负载均衡器ID。 支持多值查询，查询条件格式：loadbalancer_id=xxx&loadbalancer_id=xxx。

	LoadbalancerId *[]string `json:"loadbalancer_id,omitempty"`
	// 云日志分组ID。 支持多值查询，查询条件格式：log_group_id=xxx&log_group_id=xxx。

	LogGroupId *[]string `json:"log_group_id,omitempty"`
	// 云日志主题ID 支持多值查询，查询条件格式：log_topic_id=xxx&log_topic_id=xxx。

	LogTopicId *[]string `json:"log_topic_id,omitempty"`
}

func (o ListLogtanksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogtanksRequest struct{}"
	}

	return strings.Join([]string{"ListLogtanksRequest", string(data)}, " ")
}
