package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeploymentCreateRequest struct {

	// 应用部署到指定集群，与node_ids二选一
	ClusterId *string `json:"cluster_id,omitempty"`

	Deployment *DeploymentRequest `json:"deployment"`

	// 应用部署描述，最大长度255，不允许^ ~ # $ % & * < > ( ) [ ] { } ' \" \\
	Description *string `json:"description,omitempty"`

	// 本次部署所使用的license额度，配合订单中的计费量纲的实际计费类型，如：视频路数/实例数/QPS。
	LicenseQuota *int32 `json:"license_quota,omitempty"`

	// 应用部署名称，只允许英文小写字母、数字、中划线，最大长度32, 英文小写字母或数字开头和结尾
	Name string `json:"name"`

	// 应用部署到指定节点，与cluster_id二选一
	NodeIds *[]string `json:"node_ids,omitempty"`

	// 应用部署来源: HiLens市场(hlm) or aigallery(aig) or 自定义(userdefined)
	Source string `json:"source"`

	// 购买应用管理服务的订单ID。
	SvcOrderId *string `json:"svc_order_id,omitempty"`

	// 部署标签
	Tags *[]DeploymentTag `json:"tags,omitempty"`

	// 部署节点标签列表，当通过节点标签进行部署的时候，需要下发该字段。
	NodeTags *[]TagRequest `json:"node_tags,omitempty"`

	// 标签部署的设备数量
	NodeNum *int32 `json:"node_num,omitempty"`
}

func (o DeploymentCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeploymentCreateRequest struct{}"
	}

	return strings.Join([]string{"DeploymentCreateRequest", string(data)}, " ")
}
