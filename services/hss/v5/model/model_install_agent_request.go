package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstallAgentRequest Request Object
type InstallAgentRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 是否允许root直连 **约束限制**: 不涉及 **取值范围**: - true：允许 - false：不允许  **默认取值**: 不涉及
	RootPermission *string `json:"root_permission,omitempty"`

	// **参数解释**: root密码 **约束限制**: 不涉及 **取值范围**: 字符长度0-256 **默认取值**: 不涉及
	RootPassword *string `json:"root_password,omitempty"`

	// **参数解释**: 登录用户 **约束限制**: 不涉及 **取值范围**: 字符长度0-256 **默认取值**: 不涉及
	LoginUser *string `json:"login_user,omitempty"`

	// **参数解释**: 登录密码 **约束限制**: 不涉及 **取值范围**: 字符长度0-256 **默认取值**: 不涉及
	LoginPassword *string `json:"login_password,omitempty"`

	// **参数解释**: 登录端口 **约束限制**: 不涉及 **取值范围**: 字符长度0-256 **默认取值**: 不涉及
	LoginPort *int32 `json:"login_port,omitempty"`

	// **参数解释**: 服务器id列表,有多个hostId的时使用逗号拼接 **约束限制**: 不涉及 **取值范围**: 字符长度0-3200 **默认取值**: 不涉及
	HostIdList string `json:"host_id_list"`

	// 密码或者密钥模板文件
	File *[]MultipartFile `json:"file,omitempty"`

	// **参数解释**: 登录类型 **约束限制**: 不涉及 **取值范围**: - single_password - single_sshKey - dew - ssh_key - password  **默认取值**: 不涉及
	LoginType string `json:"login_type"`

	// **参数解释**: 服务器组id **约束限制**: 不涉及 **取值范围**: 字符长度0-128 **默认取值**: 不涉及
	HostGroupId *string `json:"host_group_id,omitempty"`
}

func (o InstallAgentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstallAgentRequest struct{}"
	}

	return strings.Join([]string{"InstallAgentRequest", string(data)}, " ")
}
