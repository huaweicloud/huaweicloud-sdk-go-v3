package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateData struct {

	// 规格ID
	FlavorId *string `json:"flavor_id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// AZ
	AvailabilityZoneId *string `json:"availability_zone_id,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 开启自动恢复
	AutoRecovery *bool `json:"auto_recovery,omitempty"`

	OsProfile *TemplateOsProfile `json:"os_profile,omitempty"`

	// 安全组ID列表。全网卡生效。
	SecurityGroupIds *[]string `json:"security_group_ids,omitempty"`

	NetworkInterfaces *[]TemplateNetworkInterfaceOption `json:"network_interfaces,omitempty"`

	// BDM挂载信息。按flavor限制为准。 1. 整机镜像，不修改卷属性，按原镜像配置创建。 2. 整机镜像，修改卷属性，要用户解开填写BDM。 3. 提供解镜像为BDM接口。
	BlockDeviceMappings *[]TemplateBlockDeviceMappingOption `json:"block_device_mappings,omitempty"`

	MarketOptions *TemplateMarketOptions `json:"market_options,omitempty"`

	InternetAccess *TemplateInternetAccessOption `json:"internet_access,omitempty"`

	Metadata map[string]string `json:"metadata,omitempty"`

	// 创建虚拟机标签，目前仅支持给虚拟机打标签，后续会支持同时给相关资源如卷等打标签
	TagOptions *[]TemplateTagOptions `json:"tag_options,omitempty"`
}

func (o TemplateData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateData struct{}"
	}

	return strings.Join([]string{"TemplateData", string(data)}, " ")
}
