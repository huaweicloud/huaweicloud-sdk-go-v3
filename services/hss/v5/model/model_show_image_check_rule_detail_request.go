package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowImageCheckRuleDetailRequest Request Object
type ShowImageCheckRuleDetailRequest struct {

	// Region ID
	Region *string `json:"region,omitempty"`

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 镜像类型，包含如下:   - private_image : 私有镜像仓库   - shared_image : 共享镜像仓库   - local_image : 本地镜像   - instance_image : 企业镜像
	ImageType string `json:"image_type"`

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

	// **参数解释** : 配置检查（基线）的类型,Linux系统支持的基线一般check_type和check_name相同,例如SSH、CentOS 7。 Windows系统支持的基线一般check_type和check_name不相同，例如check_name为Windows的配置检查（基线），它的check_type包含Windows Server 2019 R2、Windows Server 2016 R2等。 **约束限制** : 不涉及 **取值范围** : check_type的值可以通过这个接口的返回数据获得：/v5/{project_id}/baseline/risk-configs **默认取值** : 不涉及
	CheckType string `json:"check_type"`

	// 检查项id
	CheckRuleId string `json:"check_rule_id"`

	// 标准类型，包含如下:   - cn_standard : 等保合规标准   - hw_standard : 云安全实践标准
	Standard string `json:"standard"`

	// 企业仓库实例ID，swr共享版无需使用该参数
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o ShowImageCheckRuleDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageCheckRuleDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowImageCheckRuleDetailRequest", string(data)}, " ")
}
