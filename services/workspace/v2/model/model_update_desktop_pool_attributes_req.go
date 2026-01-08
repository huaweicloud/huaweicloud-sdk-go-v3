package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDesktopPoolAttributesReq 修改桌面池属性请求。
type UpdateDesktopPoolAttributesReq struct {

	// 桌面池名称，桌面池名称必须保证唯一。桌面名称只允许输入中文、大写字母、小写字母、数字、中划线，长度范围为1~255。
	Name *string `json:"name,omitempty"`

	// 桌面池描述。
	Description *string `json:"description,omitempty"`

	// OU名称，在对接AD时使用，需提前在AD中创建OU。
	OuName *string `json:"ou_name,omitempty"`

	// 标签列表。
	Tags *[]Tag `json:"tags,omitempty"`

	// 桌面断连多少分钟内，保留用户与桌面的绑定关系，超时后自动解绑。
	DisconnectedRetentionPeriod *int32 `json:"disconnected_retention_period,omitempty"`

	// 桌面池是否开启弹性伸缩类型，为false则表示不开启弹性伸缩，为true则表示开启弹性伸缩。
	EnableAutoscale *bool `json:"enable_autoscale,omitempty"`

	AutoscalePolicy *AutoscalePolicy `json:"autoscale_policy,omitempty"`

	// 是否处于管理员维护模式。
	InMaintenanceMode *bool `json:"in_maintenance_mode,omitempty"`

	// 策略id，用于指定生成桌面名称策略。
	DesktopNamePolicyId *string `json:"desktop_name_policy_id,omitempty"`

	// 桌面池的可用区。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	RootVolume *DesktopPoolVolumeInfo `json:"root_volume,omitempty"`

	// 数据盘列表。
	DataVolumes *[]DesktopPoolVolumeInfo `json:"data_volumes,omitempty"`

	// VPC ID
	VpcId *string `json:"vpc_id,omitempty"`

	// 桌面池子网ID列表
	SubnetIds *[]string `json:"subnet_ids,omitempty"`

	// 桌面使用的安全组，如果不指定则默认使用桌面代理中指定的安全组
	SecurityGroups *[]SecurityGroup `json:"security_groups,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o UpdateDesktopPoolAttributesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDesktopPoolAttributesReq struct{}"
	}

	return strings.Join([]string{"UpdateDesktopPoolAttributesReq", string(data)}, " ")
}
