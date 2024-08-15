package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceCompliantRequest Request Object
type ListInstanceCompliantRequest struct {

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// ECS实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// 内网ip
	Ip *string `json:"ip,omitempty"`

	// 弹性公网ip
	Eip *string `json:"eip,omitempty"`

	// 操作系统 - HuaweiCloudEulerOS - CentOS - EulerOS
	OperatingSystem *string `json:"operating_system,omitempty"`

	// 区域
	Region *string `json:"region,omitempty"`

	// 分组
	Group *string `json:"group,omitempty"`

	// 合规性状态 - non_compliant：不合规 - compliant：合规
	CompliantStatus *string `json:"compliant_status,omitempty"`

	// 工单id
	OrderId *string `json:"order_id,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 每页数量
	Limit *int32 `json:"limit,omitempty"`

	// 排序 - asc：升序 - desc：降序
	SortDir *string `json:"sort_dir,omitempty"`

	// 排序字段 - report_time：报告时间
	SortKey *string `json:"sort_key,omitempty"`

	// 报告场景 - CCE  - ECS
	ReportScene *string `json:"report_scene,omitempty"`

	// cce 集群信息id
	CceInfoId *string `json:"cce_info_id,omitempty"`
}

func (o ListInstanceCompliantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceCompliantRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceCompliantRequest", string(data)}, " ")
}
