package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Vault struct {
	Billing *Billing `json:"billing"`
	// 存储库自定义描述信息。

	Description *string `json:"description,omitempty"`
	// 保管库ID

	Id string `json:"id"`
	// 保管库名称

	Name string `json:"name"`
	// 项目ID

	ProjectId string `json:"project_id"`
	// 保管库资源类型id

	ProviderId string `json:"provider_id"`
	// 存储库资源

	Resources []ResourceResp `json:"resources"`
	// 存储库标签

	Tags *[]Tag `json:"tags,omitempty"`
	// 企业项目id，默认为‘0’。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 是否自动绑定，默认为false，不支持。

	AutoBind *bool `json:"auto_bind,omitempty"`

	BindRules *VaultBindRules `json:"bind_rules,omitempty"`
	// 是否开启存储库自动扩容能力（只支持按需存储库）。

	AutoExpand *bool `json:"auto_expand,omitempty"`
	// 存储库smn消息通知开关

	SmnNotify *bool `json:"smn_notify,omitempty"`
	// 存储库容量阈值，已用容量占总容量达到此百分比即发送相关通知

	Threshold *int32 `json:"threshold,omitempty"`
	// 用户id

	UserId *string `json:"user_id,omitempty"`
	// 创建时间

	CreatedAt *string `json:"created_at,omitempty"`
}

func (o Vault) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Vault struct{}"
	}

	return strings.Join([]string{"Vault", string(data)}, " ")
}
