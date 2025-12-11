package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWebAppAndServicesRequest Request Object
type ListWebAppAndServicesRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 资产类别 **约束限制**: 不涉及 **取值范围**: - host：主机资产 - container：容器资产  **默认取值**: host
	Category string `json:"category"`

	// **参数解释**: web应用，web服务或数据库名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-256 **默认取值**: 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**: 资产类型 **约束限制**: 不涉及 **取值范围**: - web-app：web应用 - web-service：web服务 - database：数据库  **默认取值**: 不涉及
	Catalogue string `json:"catalogue"`

	// **参数解释**: 服务器名称(可选)，可让用户根据主机名字搜索 **约束限制**: 不涉及 **取值范围**: 字符长度0-256 **默认取值**: 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器id(可选)，可让用户根据主机id搜索 **约束限制**: 不涉及 **取值范围**: 字符长度0-64 **默认取值**: 不涉及
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器ip(可选)，可让用户根据主机ip搜索 **约束限制**: 不涉及 **取值范围**: 字符长度0-256 **默认取值**: 不涉及
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释**: web应用，web服务或数据库版本，可让用户根据版本搜索 **约束限制**: 不涉及 **取值范围**: 字符长度0-256 **默认取值**: 不涉及
	Version *string `json:"version,omitempty"`

	// **参数解释**: web应用，web服务或数据库安装目录，可让用户根据安装目录搜索 **约束限制**: 不涉及 **取值范围**: 字符长度0-256 **默认取值**: 不涉及
	InstallDir *string `json:"install_dir,omitempty"`

	// **参数解释**: 是否模糊匹配，默认false表示精确匹配 **约束限制**: 不涉及 **取值范围**: - true：模糊匹配 - false：精确匹配  **默认取值**: false
	PartMatch *bool `json:"part_match,omitempty"`
}

func (o ListWebAppAndServicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWebAppAndServicesRequest struct{}"
	}

	return strings.Join([]string{"ListWebAppAndServicesRequest", string(data)}, " ")
}
