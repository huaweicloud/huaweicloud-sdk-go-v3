package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWebTamperHostPolicyRequest Request Object
type UpdateWebTamperHostPolicyRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 服务器ID，要求服务器已开启网页防篡改防护 **约束限制**: 需要使用 ListWtpProtectHost 接口查询网页防篡改主机防护状态列表信息，在 ListWtpProtectHost 接口的响应体中，protect_status 等于 opened，open_failed，protection_pause或partial_protection 的 host_id 是符合修改条件的服务器ID **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	HostId string `json:"host_id"`

	Body *UpdateWebTamperHostPolicyRequestInfo `json:"body,omitempty"`
}

func (o UpdateWebTamperHostPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWebTamperHostPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateWebTamperHostPolicyRequest", string(data)}, " ")
}
