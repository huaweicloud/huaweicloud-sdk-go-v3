package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomerInstanceVo struct {

	// **参数解释**：  实例ID。  **参数范围**：  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：  实例状态。  **参数范围**：  不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**：  实例名称。  **参数范围**：  不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：  更新时间。  **参数范围**：  不涉及。
	Updated *string `json:"updated,omitempty"`

	// **参数解释**：  描述。  **参数范围**：  不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：  可用区。  **参数范围**：  不涉及。
	AvailableZone *string `json:"available_zone,omitempty"`

	// **参数解释**：  虚拟私有云ID。  **参数范围**：  不涉及。
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释**：  子网ID。  **参数范围**：  不涉及。
	SubnetId *string `json:"subnet_id,omitempty"`

	// **参数解释**：  安全组ID。  **参数范围**：  不涉及。
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// **参数解释**：  节点数量。  **参数范围**：  不涉及。
	NodeCount *int32 `json:"node_count,omitempty"`

	// **参数解释**：  访问IP。  **参数范围**：  不涉及。
	AccessIp *string `json:"access_ip,omitempty"`

	// **参数解释**：  访问端口。  **参数范围**：  不涉及。
	AccessPort *string `json:"access_port,omitempty"`

	// **参数解释**：  核数。  **参数范围**：  不涉及。
	CoreCount *string `json:"core_count,omitempty"`

	// **参数解释**：  内存大小。  **参数范围**：  不涉及。
	RamCapacity *string `json:"ram_capacity,omitempty"`

	// **参数解释**：  错误信息。  **参数范围**：  不涉及。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// **参数解释**：  节点状态。  **参数范围**：  不涉及。
	NodeStatus *string `json:"node_status,omitempty"`

	// **参数解释**：  维护时间。  **参数范围**：  不涉及。
	MaintenanceTime *string `json:"maintenance_time,omitempty"`

	// **参数解释**：  企业项目ID。  **参数范围**：  不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**：  项目ID。  **参数范围**：  不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**：  引擎版本。  **参数范围**：  不涉及。
	EngineVersion *string `json:"engine_version,omitempty"`

	// **参数解释**：  订单ID。  **参数范围**：  不涉及。
	OrderId *string `json:"order_id,omitempty"`

	// **参数解释**：  管理员账号。  **参数范围**：  不涉及。
	AdminUserName *string `json:"admin_user_name,omitempty"`

	// **参数解释**：  是否支持ssl。  **参数范围**：  不涉及。
	EnableSsl *bool `json:"enable_ssl,omitempty"`

	// **参数解释**：  规格码。  **参数范围**：  不涉及。
	FlavorRef *string `json:"flavor_ref,omitempty"`

	// **参数解释**：  标签的集合。  **参数范围**：  不涉及。
	Tags *[]Tags `json:"tags,omitempty"`
}

func (o CustomerInstanceVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomerInstanceVo struct{}"
	}

	return strings.Join([]string{"CustomerInstanceVo", string(data)}, " ")
}
