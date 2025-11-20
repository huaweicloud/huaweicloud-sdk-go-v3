package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MigrateNodeExtendParam struct {

	// **参数解释**： 节点最大允许创建的实例数(Pod)，该数量包含系统默认实例。 该设置的目的为防止节点因管理过多实例而负载过重，请根据您的业务需要进行设置。 节点可以创建多少个Pod，受多个参数影响，具体请参见[节点可创建的最大Pod数量说明](maxPods.xml)。 **约束限制**： 不涉及 **取值范围**： 取值范围为16~256。 **默认取值**： 不涉及
	MaxPods *int32 `json:"maxPods,omitempty"`

	// **参数解释**： Docker数据盘配置项。  待迁移节点的磁盘类型须和创建时一致（即“DockerLVMConfigOverride”参数中“diskType”字段的值须和创建时一致），请确保单次接口调用时批量选择的节点磁盘类型一致。  默认配置示例如下：  ``` \"DockerLVMConfigOverride\":\"dockerThinpool=vgpaas/90%VG;kubernetesLV=vgpaas/10%VG;diskType=evs;lvType=linear\" ```  **约束限制**： 不涉及 **取值范围**： 包含如下字段：   - userLV（可选）：用户空间的大小，示例格式：vgpaas/20%VG   - userPath（可选）：用户空间挂载路径，示例格式：/home/wqt-test   - diskType：磁盘类型，目前只有evs、hdd和ssd三种格式   - lvType：逻辑卷的类型，目前支持linear和striped两种，示例格式：striped   - dockerThinpool：Docker盘的空间大小，示例格式：vgpaas/60%VG   - kubernetesLV：Kubelet空间大小，示例格式：vgpaas/20%VG  **默认取值**： 不涉及
	DockerLVMConfigOverride *string `json:"DockerLVMConfigOverride,omitempty"`

	// **参数解释**： 安装前执行脚本。 **约束限制**： 安装前/后执行脚本统一计算字符，转码后的字符总数不能超过10240。 输入的值需要经过Base64编码，方法如下：   ```   echo -n \"待编码内容\" | base64   ``` **取值范围**： 不涉及 **默认取值**： 不涉及
	AlphaCcePreInstall *string `json:"alpha.cce/preInstall,omitempty"`

	// **参数解释**： 安装后执行脚本。 **约束限制**： 安装前/后执行脚本统一计算字符，转码后的字符总数不能超过10240。 输入的值需要经过Base64编码，方法如下：   ```   echo -n \"待编码内容\" | base64   ``` **取值范围**： 不涉及 **默认取值**： 不涉及
	AlphaCcePostInstall *string `json:"alpha.cce/postInstall,omitempty"`

	// **参数解释**： 指定待切换目标操作系统所使用的用户镜像ID。 > 当指定“alpha.cce/NodeImageID”参数时，“os”参数必须和用户自定义镜像的操作系统一致。  **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	AlphaCceNodeImageID *string `json:"alpha.cce/NodeImageID,omitempty"`
}

func (o MigrateNodeExtendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateNodeExtendParam struct{}"
	}

	return strings.Join([]string{"MigrateNodeExtendParam", string(data)}, " ")
}
