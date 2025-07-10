package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDesktopPoolDetailResponse Response Object
type ShowDesktopPoolDetailResponse struct {

	// 桌面池ID。
	Id *string `json:"id,omitempty"`

	// 桌面池名称。
	Name *string `json:"name,omitempty"`

	// 桌面池类型。DYNAMIC：动态池，STATIC：静态池。
	Type *string `json:"type,omitempty"`

	// 桌面池描述。
	Description *string `json:"description,omitempty"`

	// 创建时间，格式为：UTC格式，例如“2022-05-11T11:45:42.000Z”。
	CreatedTime *string `json:"created_time,omitempty"`

	// 计费模式，0：包周期，1：按需。
	ChargingMode *string `json:"charging_mode,omitempty"`

	// 桌面池总桌面数量。
	DesktopCount *int32 `json:"desktop_count,omitempty"`

	// 桌面池绑定用户的桌面个数。
	DesktopUsed *int32 `json:"desktop_used,omitempty"`

	// 可用区。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 子网ID。
	SubnetId *string `json:"subnet_id,omitempty"`

	Product *ProductInfo `json:"product,omitempty"`

	// 镜像ID。
	ImageId *string `json:"image_id,omitempty"`

	// 镜像名称。
	ImageName *string `json:"image_name,omitempty"`

	// 镜像OS类型。
	ImageOsType *string `json:"image_os_type,omitempty"`

	// 镜像OS版本。
	ImageOsVersion *string `json:"image_os_version,omitempty"`

	// 镜像OS平台。
	ImageOsPlatform *string `json:"image_os_platform,omitempty"`

	// 镜像的productCode（specCode）。
	ImageProductCode *string `json:"image_product_code,omitempty"`

	RootVolume *VolumeInfo `json:"root_volume,omitempty"`

	// 数据盘列表。
	DataVolumes *[]VolumeInfo `json:"data_volumes,omitempty"`

	// 桌面安全组。
	SecurityGroups *[]SecurityGroupInfo `json:"security_groups,omitempty"`

	// 动态池桌面断连多少分钟内，保留用户与桌面的绑定关系，超时后自动解绑。
	DisconnectedRetentionPeriod *int32 `json:"disconnected_retention_period,omitempty"`

	// 桌面池是否开启弹性伸缩类型，为false则表示不开启弹性伸缩，为true则表示开启弹性伸缩。
	EnableAutoscale *bool `json:"enable_autoscale,omitempty"`

	AutoscalePolicy *AutoscalePolicy `json:"autoscale_policy,omitempty"`

	// 桌面池状态。 - STEADY：稳态 - TEMPORARY：临时态 - EXIST_FROZEN：存在冻结桌面 - UNKNOWN：未知态
	Status *string `json:"status,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 桌面池是否处于管理员维护模式。
	InMaintenanceMode *bool `json:"in_maintenance_mode,omitempty"`

	// 策略id，用于指定生成桌面名称策略。
	DesktopNamePolicyId *string `json:"desktop_name_policy_id,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ShowDesktopPoolDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDesktopPoolDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowDesktopPoolDetailResponse", string(data)}, " ")
}
