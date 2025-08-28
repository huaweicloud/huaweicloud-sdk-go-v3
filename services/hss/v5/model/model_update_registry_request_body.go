package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRegistryRequestBody 镜像仓更新信息
type UpdateRegistryRequestBody struct {

	// **参数解释**： 镜像仓接口版本 **约束限制**： 不涉及 **取值范围**：   - V1：V1版本   - V2：V2版本  **默认取值**： 不涉及
	ApiVersion *string `json:"api_version,omitempty"`

	// **参数解释**： 协议类型 **约束限制**： 不涉及 **取值范围**：   - http：http协议   - https：https协议  **默认取值**： 不涉及
	Protocol *string `json:"protocol,omitempty"`

	// **参数解释**： 镜像仓地址 **约束限制**： 网址/IP:端口。如：myharbor.com。 **取值范围**： 字符长度1-256位  **默认取值**： 不涉及
	RegistryAddr *string `json:"registry_addr,omitempty"`

	// **参数解释**： 用于登录镜像仓的用户名。 **约束限制**： 不涉及 **取值范围**： 字符长度1-128位  **默认取值**： 不涉及
	RegistryUsername *string `json:"registry_username,omitempty"`

	// **参数解释**： 用于登录镜像仓的密码。仅用于访问镜像仓，HSS服务不会存储您的镜像仓密码。 **约束限制**： 不涉及 **取值范围**： 字符长度1-128位  **默认取值**： 不涉及
	RegistryPassword *string `json:"registry_password,omitempty"`

	// **参数解释**： 镜像仓项目,用来指定扫描组件要上传到的镜像仓目录。get_scan_image_channel为Other时需要填写。 **约束限制**： 不涉及 **取值范围**： 字符长度1-128位  **默认取值**： 不涉及
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**： 承载集群id。请选择一个已接入HSS的集群作为承载集群，镜像同步组件和扫描组件会以任务的方式运行在您所选的集群上，来帮助主机安全获取您的镜像数据和扫描镜像安全风险。 **约束限制**： 不涉及 **取值范围**： 字符长度1-64位  **默认取值**： 不涉及
	ConnectClusterId *string `json:"connect_cluster_id,omitempty"`

	// **参数解释**： 获取扫描组件的方式 **约束限制**： 不涉及 **取值范围**： - Swr：访问swr获取扫描组件 - Other：手动上传扫描组件到承载集群。选择此方式，需要调用接口/v5/{project_id}/image/registries/image-upload-command 获取扫描组件镜像上传指令。  **默认取值**： 不涉及
	GetScanImageChannel *string `json:"get_scan_image_channel,omitempty"`
}

func (o UpdateRegistryRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRegistryRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateRegistryRequestBody", string(data)}, " ")
}
