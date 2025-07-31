package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCommonTipsRequest Request Object
type ListCommonTipsRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 每页显示个数，默认200
	Limit int32 `json:"limit"`

	// 提示信息的类型，包含如下： - host_name ：主机名称。 - host_id：主机id。 - public_ip：公网ip。 - private_ip ：私网ip。 - vpc_id ：vpc id。 - cluster_id ：集群 id。 - host_tags：主机标签。
	Type *string `json:"type,omitempty"`
}

func (o ListCommonTipsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCommonTipsRequest struct{}"
	}

	return strings.Join([]string{"ListCommonTipsRequest", string(data)}, " ")
}
