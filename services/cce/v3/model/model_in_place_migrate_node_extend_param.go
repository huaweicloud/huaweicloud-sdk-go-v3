package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InPlaceMigrateNodeExtendParam struct {

	// **参数解释**： 安装前执行脚本。 **约束限制**： 安装前/后执行脚本统一计算字符，转码后的字符总数不能超过10240。 输入的值需要经过Base64编码，方法如下：   ```   echo -n \"待编码内容\" | base64   ```  **取值范围**： 不涉及 **默认取值**： 不涉及
	AlphaCcePreInstall *string `json:"alpha.cce/preInstall,omitempty"`

	// **参数解释**： 安装后执行脚本。 **约束限制**： 安装前/后执行脚本统一计算字符，转码后的字符总数不能超过10240。 输入的值需要经过Base64编码，方法如下：   ```   echo -n \"待编码内容\" | base64   ```  **取值范围**： 不涉及 **默认取值**： 不涉及
	AlphaCcePostInstall *string `json:"alpha.cce/postInstall,omitempty"`

	// **参数解释：** 该参数用于控制腾挪节点时， **post-install脚本执行完成前允许节点调度** 的行为。当该参数未设置或者为false时，在kubernetes节点就绪时，容器即可被调度到可用节点。当该参数为true时，在kubernetes节点就绪时且post-install脚本执行完成时，容器才可被调度到可用节点。 **约束限制：** 不涉及 **取值范围：** - false：在kubernetes节点就绪时，容器即可被调度到可用节点。           - true：在kubernetes节点就绪时且post-install脚本执行完成时，容器才可被调度到可用节点。  **默认取值：** false
	WaitPostInstallFinish *bool `json:"waitPostInstallFinish,omitempty"`
}

func (o InPlaceMigrateNodeExtendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InPlaceMigrateNodeExtendParam struct{}"
	}

	return strings.Join([]string{"InPlaceMigrateNodeExtendParam", string(data)}, " ")
}
