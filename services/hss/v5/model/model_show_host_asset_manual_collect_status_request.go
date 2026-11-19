package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHostAssetManualCollectStatusRequest Request Object
type ShowHostAssetManualCollectStatusRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 资产类型 **约束限制**： 不涉及 **取值范围**： 包含如下11种： - web-app：web应用 - web-service：web服务 - web-framework：web框架 - web-site：web站点 - midware：中间件 - database：数据库 - kernel-module：内核模块 - software：软件 - account：账户 - ai-app：AI应用 - ai-tool：AI工具  **默认取值**： 不涉及
	Type string `json:"type"`

	// **参数解释**: 服务器的唯一标识ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	HostId *string `json:"host_id,omitempty"`
}

func (o ShowHostAssetManualCollectStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHostAssetManualCollectStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowHostAssetManualCollectStatusRequest", string(data)}, " ")
}
