package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExportContainerListRequestBody struct {

	// **参数解释**: 是否是集群内的容器。只导出集群内容器时，设置此参数值为true。只导出非集群容器时，设置此参数为false。 **约束限制**: 不涉及 **取值范围**: true或者false。 **默认取值**: 不涉及
	ClusterContainer *bool `json:"cluster_container,omitempty"`

	// **参数解释**: 集群类型。 **约束限制**: 不涉及 **取值范围**:   - cce：CCE集群   - ali：阿里云集群   - tencent：腾讯云集群   - azure：微软云集群   - aws：亚马逊集群   - self_built_hw：华为云自建集群   - self_built_idc：IDC自建集群  **默认取值**: 不涉及
	ClusterType *string `json:"cluster_type,omitempty"`

	// **参数解释**: 所属集群名称。 **约束限制**: 不涉及 **取值范围**: 字符长度1-255。 **默认取值**: 不涉及
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释**: 容器名称。 **约束限制**: 不涉及 **取值范围**: 字符长度1-255。 **默认取值**: 不涉及
	ContainerName *string `json:"container_name,omitempty"`

	// **参数解释**: 所属Pod名称。 **约束限制**: 不涉及 **取值范围**: 字符长度1-512。 **默认取值**: 不涉及
	PodName *string `json:"pod_name,omitempty"`

	// **参数解释**: 镜像名称。 **约束限制**: 不涉及 **取值范围**: 字符长度1-255。 **默认取值**: 不涉及
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**: 容器状态。 **约束限制**: 不涉及 **取值范围**:   - Running : 运行中   - Waiting : 等待   - Terminated : 终止   - Isolated : 已隔离   - Paused : 已暂停  **默认取值**: 不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释**: 是否有安全风险。只导出有安全风险容器时，设置此参数值为true。只导出无安全风险容器时，设置此参数为false。 **约束限制**: 不涉及 **取值范围**: true或者false。 **默认取值**: 不涉及
	Risky *bool `json:"risky,omitempty"`

	CreateTime *ExportContainerListRequestBodyCreateTime `json:"create_time,omitempty"`

	// **参数解释**: cpu限制。 **约束限制**: 不涉及 **取值范围**: 字符长度0-64。以m为单位，例如100m。 **默认取值**: 不涉及
	CpuLimit *string `json:"cpu_limit,omitempty"`

	// **参数解释**: 内存限制。 **约束限制**: 不涉及 **取值范围**: 字符长度0-64。以Mi、Gi为单位，例如300Mi。 **默认取值**: 不涉及
	MemoryLimit *string `json:"memory_limit,omitempty"`

	// **参数解释**: 导出容器列表的表头信息。 **约束限制**: 不涉及 **取值范围**: 合法的key值及其对应的表头名称（表头名称可自定义）   - container_id：容器id   - container_name：容器名称   - image_name：镜像名称   - pod_name：所属POD   - cluster_name：所属集群   - cluster_type：集群类型   - status：状态   - risky：是否有安全风险   - low_risk：低危风险   - medium_risk：中危风险   - high_risk：高危风险   - fatal_risk：致命风险   - create_time：创建时间   - restart_count：重启次数   - cpu_limit：CPU限制   - memory_limit：内存限制  **默认取值**: 不涉及
	ExportHeaders [][]string `json:"export_headers"`
}

func (o ExportContainerListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportContainerListRequestBody struct{}"
	}

	return strings.Join([]string{"ExportContainerListRequestBody", string(data)}, " ")
}
