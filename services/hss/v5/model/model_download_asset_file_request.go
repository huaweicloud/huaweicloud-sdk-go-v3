package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadAssetFileRequest Request Object
type DownloadAssetFileRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 资产指纹类型 **约束限制**: 不涉及 **取值范围**: - users         ：账号 - auto_launch   ：自启动项 - database      ：数据库 - jar_package   ：中间件 - port          ：开放端口 - process       ：进程 - web_cms       ：web应用 - web_framework ：web框架 - web_service   ：web服务 - web_site      ：web站点 - app           ：软件 - kernel_module ：内核模块  **默认取值**: 不涉及
	AssetType string `json:"asset_type"`

	// **参数解释**: 类别 **约束限制**: 不涉及 **取值范围**: - host     ：主机 - container：容器  **默认取值**: host
	Category string `json:"category"`

	Body *DownloadAssetFileRequestBody `json:"body,omitempty"`
}

func (o DownloadAssetFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadAssetFileRequest struct{}"
	}

	return strings.Join([]string{"DownloadAssetFileRequest", string(data)}, " ")
}
