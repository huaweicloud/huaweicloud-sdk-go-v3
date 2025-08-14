package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageRiskConfigRulesRequest Request Object
type ListImageRiskConfigRulesRequest struct {

	// Region ID
	Region *string `json:"region,omitempty"`

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 镜像类型，包含如下:   - private_image : 私有镜像仓库   - shared_image : 共享镜像仓库   - local_image : 本地镜像   - instance_image : 企业镜像
	ImageType string `json:"image_type"`

	// 偏移量：指定返回记录的开始位置
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示数量
	Limit *int32 `json:"limit,omitempty"`

	// 组织名称（没有镜像相关信息时，表示查询所有镜像）
	Namespace *string `json:"namespace,omitempty"`

	// 镜像名称
	ImageName *string `json:"image_name,omitempty"`

	// 镜像版本名称
	ImageVersion *string `json:"image_version,omitempty"`

	// 镜像id
	ImageId *string `json:"image_id,omitempty"`

	// 基线名称
	CheckName string `json:"check_name"`

	// 标准类型，包含如下:   - cn_standard : 等保合规标准   - hw_standard : 云安全实践标准
	Standard string `json:"standard"`

	// 结果类型，包含如下： - pass：已通过 - failed：未通过
	ResultType *string `json:"result_type,omitempty"`

	// 检查项名称，支持模糊匹配
	CheckRuleName *string `json:"check_rule_name,omitempty"`

	// 风险等级，包含如下:   - Security : 安全   - Low : 低危   - Medium : 中危   - High : 高危   - Critical : 危急
	Severity *string `json:"severity,omitempty"`

	// 企业仓库实例ID，swr共享版无需使用该参数
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o ListImageRiskConfigRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageRiskConfigRulesRequest struct{}"
	}

	return strings.Join([]string{"ListImageRiskConfigRulesRequest", string(data)}, " ")
}
