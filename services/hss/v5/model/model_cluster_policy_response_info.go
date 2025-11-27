package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterPolicyResponseInfo 集群防护策略信息
type ClusterPolicyResponseInfo struct {

	// **参数解释**： 集群ID **取值范围**： 不涉及
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**： 集群名称 **取值范围**： 不涉及
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释**： 策略内容 **取值范围**： 不涉及
	Content *interface{} `json:"content,omitempty"`

	// **参数解释**： 部署内容 **取值范围**： 不涉及
	DeployContent *string `json:"deploy_content,omitempty"`

	// **参数解释**： 参数 **取值范围**： 不涉及
	Parameters *string `json:"parameters,omitempty"`

	// **参数解释**： 策略名称 **取值范围**： 不涉及
	PolicyName *string `json:"policy_name,omitempty"`

	// **参数解释**： 策略ID **取值范围**： 不涉及
	PolicyId *string `json:"policy_id,omitempty"`

	// **参数解释**： 资源 **取值范围**： 不涉及
	Resources *[]Resources `json:"resources,omitempty"`

	// **参数解释**： 模板ID **取值范围**： 不涉及
	TemplateId *string `json:"template_id,omitempty"`

	// **参数解释**： 模板名称 **取值范围**： 不涉及
	TemplateName *string `json:"template_name,omitempty"`

	// **参数解释**： 模板类型 **取值范围**： 不涉及
	TemplateType *string `json:"template_type,omitempty"`

	// **参数解释**： 策略描述 **取值范围**： 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**： 更新时间 **取值范围**： 不涉及
	UpdateTime *int32 `json:"update_time,omitempty"`

	// **参数解释**： 创建时间 **取值范围**： 不涉及
	CreateTime *int32 `json:"create_time,omitempty"`

	// **参数解释**： 防护镜像数量 **取值范围**： 不涉及
	ImageNum *int32 `json:"image_num,omitempty"`

	// **参数解释**： 防护标签数量 **取值范围**： 不涉及
	LabelsNum *int32 `json:"labels_num,omitempty"`

	// **参数解释**： 状态 **取值范围**： 不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释**： 白名单镜像 **取值范围**： 不涉及
	WhiteImages *[]WhiteImageInfo `json:"white_images,omitempty"`
}

func (o ClusterPolicyResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterPolicyResponseInfo struct{}"
	}

	return strings.Join([]string{"ClusterPolicyResponseInfo", string(data)}, " ")
}
