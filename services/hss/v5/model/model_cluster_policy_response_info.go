package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterPolicyResponseInfo 集群防护策略信息
type ClusterPolicyResponseInfo struct {

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 策略内容
	Content *interface{} `json:"content,omitempty"`

	// deploy内容
	DeployContent *string `json:"deploy_content,omitempty"`

	// 参数
	Parameters *string `json:"parameters,omitempty"`

	// 策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// 策略ID
	PolicyId *string `json:"policy_id,omitempty"`

	// 资源
	Resources *[]Resources `json:"resources,omitempty"`

	// 模板ID
	TemplateId *string `json:"template_id,omitempty"`

	// 模板名称
	TemplateName *string `json:"template_name,omitempty"`

	// 模板类型
	TemplateType *string `json:"template_type,omitempty"`

	// 策略描述
	Description *string `json:"description,omitempty"`

	// 更新时间
	UpdateTime *int32 `json:"update_time,omitempty"`

	// 创建时间
	CreateTime *int32 `json:"create_time,omitempty"`

	// 防护镜像数量
	ImageNum *int32 `json:"image_num,omitempty"`

	// 防护标签数量
	LabelsNum *int32 `json:"labels_num,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 白名单镜像
	WhiteImages *[]WhiteImageInfo `json:"white_images,omitempty"`
}

func (o ClusterPolicyResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterPolicyResponseInfo struct{}"
	}

	return strings.Join([]string{"ClusterPolicyResponseInfo", string(data)}, " ")
}
