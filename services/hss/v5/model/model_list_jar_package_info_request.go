package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJarPackageInfoRequest Request Object
type ListJarPackageInfoRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 服务器ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	HostId string `json:"host_id"`

	// **参数解释**: 资产类别 **约束限制**: 不涉及 **取值范围**: - host：主机资产 - container：容器资产  **默认取值**: host
	Category string `json:"category"`

	// **参数解释**: 文件名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-256 **默认取值**: 不涉及
	FileName *string `json:"file_name,omitempty"`
}

func (o ListJarPackageInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJarPackageInfoRequest struct{}"
	}

	return strings.Join([]string{"ListJarPackageInfoRequest", string(data)}, " ")
}
