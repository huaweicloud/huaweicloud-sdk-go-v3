package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SiteNetworkEntry 分支网络。
type SiteNetworkEntry struct {

	// 实例ID。
	Id string `json:"id"`

	// 实例名称。
	Name string `json:"name"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 实例创建时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 实例更新时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 实例所属账号ID。
	DomainId string `json:"domain_id"`

	State *SiteNetworkStateEnum `json:"state"`

	// 实例所属企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 应用策略ID。
	ApplyPolicyId string `json:"apply_policy_id"`

	// 实例标签。
	Tags *[]Tag `json:"tags,omitempty"`

	Topology *SiteNetworkTopologyEnum `json:"topology"`

	// 分支连接列表。
	Connections []SiteConnection `json:"connections"`

	// 点对点拓扑或者网状拓扑中的节点。
	Sites *[]SiteInformation `json:"sites,omitempty"`

	HubSite *SiteInformation `json:"hub_site,omitempty"`

	// 分支列表。
	SpokeSites *[]SiteInformation `json:"spoke_sites,omitempty"`
}

func (o SiteNetworkEntry) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SiteNetworkEntry struct{}"
	}

	return strings.Join([]string{"SiteNetworkEntry", string(data)}, " ")
}
