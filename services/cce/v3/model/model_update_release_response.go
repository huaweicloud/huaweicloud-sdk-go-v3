package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateReleaseResponse Response Object
type UpdateReleaseResponse struct {

	// **参数解释：** 模板名称 **约束限制：** 最长64个字符 **取值范围：** 不涉及 **默认取值：** 不涉及
	ChartName *string `json:"chart_name,omitempty"`

	// **参数解释：** 是否公开模板 **约束限制：** 不涉及 **取值范围：** - true：公开模板 - false：不公开模板  **默认取值：** false
	ChartPublic *bool `json:"chart_public,omitempty"`

	// **参数解释：** 模板版本 **约束限制：** 最长64个字符 **取值范围：** 不涉及 **默认取值：** 不涉及
	ChartVersion *string `json:"chart_version,omitempty"`

	// **参数解释：** 集群ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释：** 集群名称 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释：** 创建时间 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	CreateAt *string `json:"create_at,omitempty"`

	// **参数解释：** 模板实例描述 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释：** 模板实例名称 **约束限制：** 由小写字母开头，中间由小写字母、数字和中划线(-)组成，以小写字母或数字结尾 **取值范围：** 不涉及 **默认取值：** 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释：** 模板实例所在的命名空间 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释：** 模板实例参数 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Parameters *string `json:"parameters,omitempty"`

	// **参数解释：** 模板实例需要的资源 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Resources *string `json:"resources,omitempty"`

	// **参数解释：** 模板实例状态 **约束限制：** 不涉及 **取值范围：** - DEPLOYED：已部署，表示模板实例处于正常状态 - DELETED：已删除，表示模板实例已经被删除 - FAILED：失败，表示模板实例部署失败 - DELETING：删除中，表示模板实例正处于删除过程中 - PENDING_INSTALL：待安装，表示模板正在等待安装 - PENDING_UPGRADE：待升级，表示模板正在等待升级 - PENDING_ROLLBACK：待回滚，表示模板正在等待回滚 - UNKNOWN：未知，表示模板状态异常，可尝试手动删除后重新安装  **默认取值：** 不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释：** 模板实例状态描述 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	StatusDescription *string `json:"status_description,omitempty"`

	// **参数解释：** 更新时间 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	UpdateAt *string `json:"update_at,omitempty"`

	// **参数解释：** 模板实例的值 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Values *string `json:"values,omitempty"`

	// **参数解释：** 模板实例版本 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Version        *int32 `json:"version,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateReleaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateReleaseResponse struct{}"
	}

	return strings.Join([]string{"UpdateReleaseResponse", string(data)}, " ")
}
