package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeLifecycleConfig 节点自定义生命周期配置
type NodeLifecycleConfig struct {

	// **参数解释**： 安装前执行脚本。安装前/后执行脚本统一计算字符，输入的值需要经过Base64编码，方法如下： ``` echo -n \"待编码内容\" | base64 ```  **约束限制**： 长度不能超过10240字节。 **取值范围**： 不涉及 **默认取值**： 不涉及
	PreInstall *string `json:"preInstall,omitempty"`

	// **参数解释**： 安装前执行脚本。安装前/后执行脚本统一计算字符，输入的值需要经过Base64编码，方法如下： ``` echo -n \"待编码内容\" | base64 ```  **约束限制**： 长度不能超过10240字节。 **取值范围**： 不涉及 **默认取值**： 不涉及
	PostInstall *string `json:"postInstall,omitempty"`

	// **参数解释：** 该参数用于控制重置/纳管/批量重置节点时， **post-install脚本执行完成前允许节点调度** 的行为。当操作的节点属于节点池时，以节点池相关配置为准。当该参数未设置或者为false时，在kubernetes节点就绪时，容器即可被调度到可用节点。当该参数为true时，在kubernetes节点就绪时且post-install脚本执行完成时，容器才可被调度到可用节点。  **约束限制：** 不涉及  **取值范围：** - false：在kubernetes节点就绪时，容器即可被调度到可用节点。           - true：在kubernetes节点就绪时且post-install脚本执行完成时，容器才可被调度到可用节点。  **默认取值：** false
	WaitPostInstallFinish *bool `json:"waitPostInstallFinish,omitempty"`
}

func (o NodeLifecycleConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeLifecycleConfig struct{}"
	}

	return strings.Join([]string{"NodeLifecycleConfig", string(data)}, " ")
}
