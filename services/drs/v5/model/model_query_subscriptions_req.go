package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuerySubscriptionsReq 查询订阅任务列表请求体
type QuerySubscriptionsReq struct {

	// 任务场景
	JobType string `json:"job_type"`

	// 引擎类型。
	EngineType *string `json:"engine_type,omitempty"`

	// 网络类型。
	NetType *string `json:"net_type,omitempty"`

	// 任务ID或名称。
	Name *string `json:"name,omitempty"`

	// 任务状态。
	Status *string `json:"status,omitempty"`

	// 企业项目ID。 缺省值：\"\"，表示查询所有企业项目任务。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 返回结果按该关键字排序，默认为“create_time”。
	SortKey *string `json:"sort_key,omitempty"`

	// 降序或升序（分别对应desc和asc，默认为“desc”）。
	SortDir *string `json:"sort_dir,omitempty"`

	// 数据库实例ID列表，最多支持10个。
	InstanceIds *[]string `json:"instance_ids,omitempty"`

	// 数据库实例IP
	InstanceIp *string `json:"instance_ip,omitempty"`

	// 标签
	Tags map[string]string `json:"tags,omitempty"`

	// 服务名称
	ServiceName *string `json:"service_name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 计费模式，包括是、否以及全部三种情况，不填默认为全部
	IsBilling *bool `json:"is_billing,omitempty"`

	// 消费时间
	BeginAt *string `json:"begin_at,omitempty"`
}

func (o QuerySubscriptionsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuerySubscriptionsReq struct{}"
	}

	return strings.Join([]string{"QuerySubscriptionsReq", string(data)}, " ")
}
