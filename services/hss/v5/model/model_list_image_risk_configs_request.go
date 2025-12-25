package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageRiskConfigsRequest Request Object
type ListImageRiskConfigsRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释** 用于筛选指定类型的镜像安全配置检测结果，不同类型对应不同镜像存储位置 **约束限制** 取值必须在指定范围内，否则返回空结果 **取值范围** - private_image： 私有镜像仓库 - shared_image： 共享镜像仓库 - local_image： 本地镜像 - instance_image： 企业镜像 - registry： 仓库镜像 - local： 本地镜像，用于查询全局数据 **默认取值** 无
	ImageType string `json:"image_type"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释** 镜像仓库的组织（命名空间）名称，用于筛选指定组织下的镜像检测结果，仅私有/共享镜像仓库有效 **约束限制** 仅当image_type为private_image或shared_image时有效，其他类型传参无效 **取值范围** 符合镜像仓库组织命名规范的字符串 **默认取值** 无
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释** 镜像的名称，用于精准筛选指定名称的镜像安全配置检测结果 **约束限制** 支持模糊匹配（如传入'euler'可匹配所有名称含'euler'的镜像） **取值范围** 符合镜像名称命名规范的字符串 **默认取值** 无
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释** 镜像的版本标识，用于筛选指定版本的镜像安全配置检测结果，与image_name配合使用 **约束限制** 仅当指定image_name时传参有效，否则筛选条件不生效 **取值范围** 符合镜像版本命名规范的字符串、默认取值：无
	ImageVersion *string `json:"image_version,omitempty"`

	// **参数解释** 镜像的唯一标识，用于精准筛选指定镜像的安全配置检测结果，优先级高于image_name+image_version **约束限制** 传入后将忽略image_name和image_version参数，直接按ID筛选 **默认取值** 无
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释** 安全配置检测的基线名称，用于筛选指定基线的检测结果（如'CentOS 7'、'EulerOS'等） **约束限制** 仅支持功能介绍中列出的系统基线（CentOS 7、Debian 10、EulerOS、Ubuntu16） **取值范围** 支持的基线名称列表详见功能介绍 **默认取值** 无
	CheckName *string `json:"check_name,omitempty"`

	// **参数解释** 镜像安全配置检测结果的风险等级，用于筛选指定风险等级的检测记录 **约束限制** 取值必须在指定范围内，否则返回空结果 **取值范围** - Security：安全 - Low：低危 - Medium：中危 - High：高危 **默认取值** 无
	Severity *string `json:"severity,omitempty"`

	// **参数解释** 安全配置检测遵循的标准，用于筛选符合指定标准的检测结果 **约束限制** 取值必须在指定范围内，否则返回空结果 **取值范围** - cn_standard：等保合规标准 - hw_standard：云安全实践标准 **默认取值** 无
	Standard *string `json:"standard,omitempty"`

	// **参数解释** 华为云SWR（软件仓库）企业版实例的唯一标识，用于筛选指定企业仓库实例下的镜像检测结果 **约束限制** 仅当image_type为private_image且使用SWR企业版时有效，共享版/本地镜像传参无效 **取值范围** SWR企业版实例ID **默认取值** 无
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o ListImageRiskConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageRiskConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListImageRiskConfigsRequest", string(data)}, " ")
}
