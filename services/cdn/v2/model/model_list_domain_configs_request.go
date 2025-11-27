package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainConfigsRequest Request Object
type ListDomainConfigsRequest struct {

	// **参数解释：** 加速域名  - 多个域名使用英文半角逗号分隔  - 当查询cname状态时，该参数必传  **约束限制：** 仅支持查询已经在CDN添加成功的域名 **取值范围：** 不涉及 **默认取值：** 不涉及
	DomainNames *string `json:"domain_names,omitempty"`

	// **参数解释：** 查询数据类型 **约束限制：** 不涉及 **取值范围：** - cname_status: 域名cname状态 - copy: 查询账号下哪些加速域名支持复制配置  **默认取值：** 不涉及
	Item string `json:"item"`

	// **参数解释：** 企业项目id。您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id **约束限制：** 当用户开启企业项目功能时，该参数生效，表示查询资源所属项目，\"all\"表示所有项目 **取值范围：** 不涉及 **默认取值：** 不涉及
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListDomainConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListDomainConfigsRequest", string(data)}, " ")
}
