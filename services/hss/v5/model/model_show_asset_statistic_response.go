package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAssetStatisticResponse Response Object
type ShowAssetStatisticResponse struct {

	// **参数解释**： 主机账号数量 **取值范围**： 最小值0，最大值2147483647
	AccountNum *int64 `json:"account_num,omitempty"`

	// **参数解释**： 开放端口数量 **取值范围**： 最小值0，最大值2147483647
	PortNum *int64 `json:"port_num,omitempty"`

	// **参数解释**： 进程数量 **取值范围**： 最小值0，最大值2147483647
	ProcessNum *int64 `json:"process_num,omitempty"`

	// **参数解释**： 软件数量 **取值范围**： 最小值0，最大值2147483647
	AppNum *int64 `json:"app_num,omitempty"`

	// **参数解释**： 自启动进程数量 **取值范围**： 最小值0，最大值2147483647
	AutoLaunchNum *int64 `json:"auto_launch_num,omitempty"`

	// **参数解释**： web框架数量 **取值范围**： 最小值0，最大值2147483647
	WebFrameworkNum *int64 `json:"web_framework_num,omitempty"`

	// **参数解释**： Web站点数量 **取值范围**： 最小值0，最大值2147483647
	WebSiteNum *int64 `json:"web_site_num,omitempty"`

	// **参数解释**： Jar包数量 **取值范围**： 最小值0，最大值2147483647
	JarPackageNum *int64 `json:"jar_package_num,omitempty"`

	// **参数解释**： 内核模块数量 **取值范围**： 最小值0，最大值2147483647
	KernelModuleNum *int64 `json:"kernel_module_num,omitempty"`

	// **参数解释**： web服务数量 **取值范围**： 最小值0，最大值2147483647
	WebServiceNum *int64 `json:"web_service_num,omitempty"`

	// **参数解释**： web应用数量 **取值范围**： 最小值0，最大值2147483647
	WebAppNum *int64 `json:"web_app_num,omitempty"`

	// **参数解释**： 数据库数量 **取值范围**： 最小值0，最大值2147483647
	DatabaseNum    *int64 `json:"database_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowAssetStatisticResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssetStatisticResponse struct{}"
	}

	return strings.Join([]string{"ShowAssetStatisticResponse", string(data)}, " ")
}
