package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAssetStatisticResponse Response Object
type ShowAssetStatisticResponse struct {

	// **参数解释**： 满足查询条件的资产（主机/容器）下的主机账号总数量 **取值范围**： 最小值0，最大值2147483647，单位：个
	AccountNum *int64 `json:"account_num,omitempty"`

	// **参数解释**： 满足查询条件的资产（主机/容器）下的开放端口总数量 **取值范围**： 最小值0，最大值2147483647，单位：个
	PortNum *int64 `json:"port_num,omitempty"`

	// **参数解释**： 满足查询条件的资产（主机/容器）下的进程总数量 **取值范围**： 最小值0，最大值2147483647，单位：个
	ProcessNum *int64 `json:"process_num,omitempty"`

	// **参数解释**： 满足查询条件的资产（主机/容器）下已安装的软件总数量 **取值范围**： 最小值0，最大值2147483647，单位：个
	AppNum *int64 `json:"app_num,omitempty"`

	// **参数解释**： 满足查询条件的资产（主机/容器）下的自启动进程总数量 **取值范围**： 最小值0，最大值2147483647，单位：个
	AutoLaunchNum *int64 `json:"auto_launch_num,omitempty"`

	// **参数解释**： 满足查询条件的资产（主机/容器）下已部署的Web框架总数量 **取值范围**： 最小值0，最大值2147483647，单位：个
	WebFrameworkNum *int64 `json:"web_framework_num,omitempty"`

	// **参数解释**： 满足查询条件的资产（主机/容器）下已部署的Web站点总数量 **取值范围**： 最小值0，最大值2147483647，单位：个
	WebSiteNum *int64 `json:"web_site_num,omitempty"`

	// **参数解释**： 满足查询条件的资产（主机/容器）下的JAR包总数量 **取值范围**： 最小值0，最大值2147483647，单位：个
	JarPackageNum *int64 `json:"jar_package_num,omitempty"`

	// **参数解释**： 满足查询条件的资产（主机/容器）下已加载的内核模块总数量 **取值范围**： 最小值0，最大值2147483647，单位：个
	KernelModuleNum *int64 `json:"kernel_module_num,omitempty"`

	// **参数解释**： 满足查询条件的资产（主机/容器）下已部署的Web服务总数量 **取值范围**： 最小值0，最大值2147483647，单位：个
	WebServiceNum *int64 `json:"web_service_num,omitempty"`

	// **参数解释**： 满足查询条件的资产（主机/容器）下已部署的Web应用总数量 **取值范围**： 最小值0，最大值2147483647，单位：个
	WebAppNum *int64 `json:"web_app_num,omitempty"`

	// **参数解释**： 满足查询条件的资产（主机/容器）下已部署的数据库总数量 **取值范围**： 最小值0，最大值2147483647，单位：个
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
