package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListContainersRequest Request Object
type ListContainersRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 容器名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	ContainerName *string `json:"container_name,omitempty"`

	// **参数解释**: 所属Pod名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	PodName *string `json:"pod_name,omitempty"`

	// **参数解释**: 镜像名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**: 是否是集群纳管的容器 **约束限制**: 不涉及 **取值范围**:   - ture：是集群纳管的容器。   - false：不是集群纳管的容器。 **默认取值**: false
	ClusterContainer *bool `json:"cluster_container,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 取值0-2000000 **默认取值**: 0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListContainersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListContainersRequest struct{}"
	}

	return strings.Join([]string{"ListContainersRequest", string(data)}, " ")
}
