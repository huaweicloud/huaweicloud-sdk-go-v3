package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCheckRuleDetailRequest Request Object
type ShowCheckRuleDetailRequest struct {

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 配置检查（基线）的名称，例如SSH、CentOS 7、Windows
	CheckName string `json:"check_name"`

	// 配置检查（基线）的类型,Linux系统支持的基线一般check_type和check_name相同,例如SSH、CentOS 7。 Windows系统支持的基线一般check_type和check_name不相同，例如check_name为Windows的配置检查（基线），它的check_type包含Windows Server 2019 R2、Windows Server 2016 R2等。check_type的值可以通过这个接口的返回数据获得：/v5/{project_id}/baseline/risk-configs
	CheckType string `json:"check_type"`

	// 检查项ID，值可以通过这个接口的返回数据获得：/v5/{project_id}/baseline/risk-config/{check_name}/check-rules
	CheckRuleId string `json:"check_rule_id"`

	// 标准类型，包含如下:   - cn_standard : 等保合规标准   - hw_standard : 云安全实践标准
	Standard string `json:"standard"`

	// 主机ID
	HostId *string `json:"host_id,omitempty"`
}

func (o ShowCheckRuleDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCheckRuleDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowCheckRuleDetailRequest", string(data)}, " ")
}
